// Copyright 2020 Saferwall. All rights reserved.
// Use of this source code is governed by Apache v2 license
// license that can be found in the LICENSE file.

package pe

import (
	"errors"
	"fmt"
	"io"
	"os"

	mmap "github.com/edsrzf/mmap-go"
)

// A File represents an open PE file.
type File struct {
	DosHeader    ImageDosHeader `json:",omitempty"`
	RichHeader   RichHeader `json:",omitempty"`
	NtHeader     ImageNtHeader `json:",omitempty"`
	Sections     []ImageSectionHeader `json:",omitempty"`
	Imports      []Import `json:",omitempty"`
	Export      Export `json:",omitempty"`
	Debugs       []DebugEntry `json:",omitempty"`
	Relocations  []Relocation `json:",omitempty"`
	Resources    ResourceDirectory `json:",omitempty"`
	TLS          TLSDirectory `json:",omitempty"`
	LoadConfig   interface{} `json:",omitempty"`
	Exceptions   []Exception `json:",omitempty"`
	Certificates Certificate `json:",omitempty"`
	DelayImports []DelayImport `json:",omitempty"`
	BoundImports []BoundImportDescriptorData `json:",omitempty"`
	GlobalPtr    uint32 `json:",omitempty"`
	CLRHeader  *ImageCOR20Header `json:",omitempty"`
	Header    []byte
	data      mmap.MMap
	closer    io.Closer
	Is64      bool
	Is32      bool
	Anomalies []string `json:",omitempty"`
	size      uint32
	f         *os.File
}

// Open opens the named file using os.Open and prepares it for use as a PE binary.
func Open(name string) (File, error) {

	// Init an File instance
	file := File{}

	f, err := os.Open(name)
	if err != nil {
		return file, err
	}

	// Memory map the file insead of using read/write.
	data, err := mmap.Map(f, mmap.RDONLY, 0)
	if err != nil {
		f.Close()
		return file, err
	}

	file.data = data
	file.size = uint32(len(file.data))
	file.f = f
	return file, nil
}

// Close closes the File.
func (pe *File) Close() error {
	var err error
	if pe.f != nil {
		err = pe.f.Close()
	}
	return err
}

// Parse performs the file parsing for a PE binary.
func (pe *File) Parse() error {

	// check for the smallest PE size.
	if len(pe.data) < TinyPESize {
		return ErrInvalidPESize
	}

	// Parse the DOS header.
	err := pe.ParseDOSHeader()
	if err != nil {
		return err
	}

	// Parse the Rich header.
	pe.ParseRichHeader()
	if err != nil {
		return err
	}

	// Parse the NT header.
	err = pe.ParseNTHeader()
	if err != nil {
		return err
	}

	// Parse the Section Header.
	err = pe.ParseSectionHeader()
	if err != nil {
		return err
	}

	// Parse the Data Directory entries.
	err = pe.ParseDataDirectories()
	return err
}

// ParseDataDirectories parses the data directores. The DataDirectory is an 
// array of 16 structures. Each array entry has a predefined meaning for what
// it refers to. 
func (pe *File) ParseDataDirectories() error {
	oh32 := ImageOptionalHeader32{}
	oh64 := ImageOptionalHeader64{}
	switch pe.Is64 {
	case true:
		oh64 = pe.NtHeader.OptionalHeader.(ImageOptionalHeader64)
	case false:
		oh32 = pe.NtHeader.OptionalHeader.(ImageOptionalHeader32)
	}

	// Maps data directory index to function which parses that directory.
	funcMaps := map[int](func(uint32, uint32) error){
		ImageDirectoryEntryExport:      pe.parseExportDirectory,
		ImageDirectoryEntryImport:      pe.parseImportDirectory,
		ImageDirectoryEntryResource:    pe.parseResourceDirectory,
		ImageDirectoryEntryException:   pe.parseExceptionDirectory,
		ImageDirectoryEntryCertificate: pe.parseSecurityDirectory,
		ImageDirectoryEntryBaseReloc:   pe.parseRelocDirectory,
		ImageDirectoryEntryDebug:       pe.parseDebugDirectory,
		ImageDirectoryEntryArchitecture:pe.parseArchitectureDirectory,
		ImageDirectoryEntryGlobalPtr:   pe.parseGlobalPtrDirectory,
		ImageDirectoryEntryTLS:         pe.parseTLSDirectory,
		ImageDirectoryEntryLoadConfig:  pe.parseLoadConfigDirectory,
		ImageDirectoryEntryBoundImport: pe.parseBoundImportDirectory,
		ImageDirectoryEntryIAT:         pe.parseIATDirectory,
		ImageDirectoryEntryDelayImport: pe.parseDelayImportDirectory,
		ImageDirectoryEntryCLR:         pe.parseCLRHeaderDirectory,
	}

	// Iterate over data directories and call the appropriate function.
	var errorMsg string
	for entryIndex := 0; entryIndex < ImageNumberOfDirectoryEntries; entryIndex++ {

		var va, size uint32
		switch pe.Is64 {
		case true:
			dirEntry := oh64.DataDirectory[entryIndex]
			va = dirEntry.VirtualAddress
			size = dirEntry.Size
		case false:
			dirEntry := oh32.DataDirectory[entryIndex]
			va = dirEntry.VirtualAddress
			size = dirEntry.Size
		}

		if va != 0 {
			err := funcMaps[entryIndex](va, size)
			if err != nil {
				// append error but keep parsing other directories.
				errorMsg += fmt.Sprintf("%s,", err.Error()) 
			}
		}
	}

	if errorMsg != "" {
		return errors.New(errorMsg)
	} 
	return nil
}
