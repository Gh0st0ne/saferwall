// Copyright 2018 Saferwall. All rights reserved.
// Use of this source code is governed by Apache v2 license
// license that can be found in the LICENSE file.

package sophos

import (
	"strings"

	"github.com/saferwall/saferwall/pkg/utils"
)

// Our consts
const (
	savscan = "/opt/sophos/bin/savscan"
)

// Result represent detection results
type Result struct {
	Infected bool   `json:"infected"`
	Output   string `json:"output"`
	Version  string `json:"version"`
}

// Version represents all sophos components' versions
type Version struct {
	ProductVersion       string `json:"product_version"`
	EngineVersion        string `json:"engine_version"`
	VirusDataVersion     string `json:"virus_data_version"`
	UserInterfaceVersion string `json:"user_interface_version"`
}

// ScanFile a file with Sophos scanner
func ScanFile(filePath string) (Result, error) {

	//  Scan parameters
	// -f  		: Full Scan
	// -c  		: Ask for confirmation before disinfection/deletion
	// -b  		: Sound bell on virus detection
	// -ss 		: Don't display anything except on error or virus
	// archive  : All of the above
	// loopback : Scan inside loopback-type files
	// mime 	: Scan files encoded in MIME format
	// oe   	: Scan Microsoft Outlook Express mailbox files (requires -mime)
	// tnef 	: Scan inside TNEF files
	// pua : Scan for adware/potentially unwanted applications (PUAs).

	savscanOut, err := utils.ExecCommand(savscan, "-f", "-nc", "-nb", "-ss",
		"-archive", "-loopback", "-mime", "-oe", "-tnef", "-pua", filePath)
	if err != nil && err.Error() != "exit status 3" {
		return Result{}, err
	}

	lines := strings.Split(savscanOut, "\n")
	r := Result{}
	for _, line := range lines {
		if strings.HasPrefix(line, ">>> Virus ") {
			detection := strings.Split(line, " ")[2]
			r.Output = detection[1 : len(detection)-1]
			r.Infected = true
			break
		}
	}

	return r, nil
}

// GetVersion returns Sophos components' version
func GetVersion() (Version, error) {

	versionOut, err := utils.ExecCommand(savscan, "--version")

	// SAVScan virus detection utility
	// Copyright (c) 1989-2018 Sophos Limited. All rights reserved.

	// System time 19:28:51, System date 24 December 2018

	// Product version           : 5.53.0
	// Engine version            : 3.74.2
	// Virus data version        : 5.55
	// User interface version    : 2.03.074
	// Platform                  : Linux/AMD64
	// Released                  : 18 September 2018
	// Total viruses (with IDEs) : 25676226

	if err != nil {
		return Version{}, err
	}

	v := Version{}
	lines := strings.Split(versionOut, "\n")
	for _, line := range lines {
		if strings.Contains(line, "Product version") {
			v.ProductVersion = strings.TrimSpace(strings.Split(line, ":")[1])
		} else if strings.Contains(line, "Engine version") {
			v.EngineVersion = strings.TrimSpace(strings.Split(line, ":")[1])
		} else if strings.Contains(line, "Virus data version") {
			v.VirusDataVersion = strings.TrimSpace(strings.Split(line, ":")[1])
		} else if strings.Contains(line, "User interface version") {
			v.UserInterfaceVersion = strings.TrimSpace(strings.Split(line, ":")[1])
		}
	}

	return v, nil
}