/*
 *  Copyright (c) 2019, Salesforce.com, Inc.
 *  All rights reserved.
 *  Licensed under the BSD 3-Clause license.
 *  For full license text, see LICENSE.txt file in the repo root  or https://opensource.org/licenses/BSD-3-Clause
 */

package bazenv

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/mitchellh/go-homedir"
)

// ResolveBazelDirectory converts a bazel version name into the path to a bazel install directory, or retuns an error
// if the install doesn't exist.
func ResolveBazelDirectory(version string) (string, error) {
	homedir, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	bazelDir := filepath.Join(homedir, BazenvDir, BazenvVersionsDir, version)
	if _, err := os.Stat(bazelDir); os.IsNotExist(err) {
		// bazelDir doesn't exist
		return "", errors.New("Bazel version '" + version + "' does not exist. Use 'bazenv global' to configure a" +
			" global bazel version.")
	}

	// Determine if bazel is symlink
	fi, err := os.Lstat(bazelDir)
	if err != nil {
		return "", err
	}
	if fi.Mode()&os.ModeSymlink == os.ModeSymlink {
		// This is a symlink
		bazelLink, err := os.Readlink(bazelDir)
		if err != nil {
			return "", err
		}
		if bazelLink != "" {
			return bazelLink, nil
		}
	}

	// Not a symlink
	return bazelDir, nil
}

// SniffBazelVersion uses the 'bazel version' command to return the version name of a bazel directory, or returns
// an error.
func SniffBazelVersion(path string) (string, error) {
	rawVersion, err := exec.Command(filepath.Join(path, "bin", "bazel"), "version").Output()
	if err != nil {
		// Error executing bazel
		return "", err
	}

	versionRegex := regexp.MustCompile(`Build label: (.*)\n`)
	match := versionRegex.FindStringSubmatch(string(rawVersion))

	if len(match) != 2 {
		// Version string not found
		return "", errors.New("Not a bazel install directory: " + path)
	}

	return match[1], nil
}

// SniffIsBazenvStub returns true of the installed bazel command is a bazenv stub
func SniffIsBazenvStub() (bool, error) {
	output, err := exec.Command("bazel", "bazenv").Output()
	if err != nil {
		return false, err
	}

	return strings.TrimSpace(string(output)) == "yes", nil
}
