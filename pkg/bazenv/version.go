/*
 *  Copyright (c) 2019, Salesforce.com, Inc.
 *  All rights reserved.
 *  Licensed under the BSD 3-Clause license.
 *  For full license text, see LICENSE.txt file in the repo root  or https://opensource.org/licenses/BSD-3-Clause
 */

// Functions for reading .bazenv_version files

package bazenv

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
)

const (
	// BazenvDir is the name of the bazenv config directory in the user's home directory
	BazenvDir = ".bazenv"
	// BazenvVersionsDir is the directory under BazenvDir where bazel versions are stored
	BazenvVersionsDir = "bazel"
	// BazenvFile is the name of a bazenv version file, prefixed with dot when not in BazenvDir
	BazenvFile = "bazenv_version"
)

// EnsureBazenvDir creates the bazenv working directoy if not found
func EnsureBazenvDir() {
	homedir, _ := homedir.Dir()
	os.MkdirAll(filepath.Join(homedir, BazenvDir, BazenvVersionsDir), os.ModePerm)
}

// ReadBazenvFile reads the content of .bazenv_version, looking locally up the directiory tree first, then in the
// global bazenv file. This returns the name of the active bazel version.
func ReadBazenvFile() (string, error) {
	// Local bazenv file takes priority
	bazenv, err := FindAndReadLocalBazenvFile()
	if err != nil {
		return "", err
	}
	if bazenv != nil {
		return *bazenv, nil
	}

	// Fall back to global bazenv
	bazenv, err = ReadGlobalBazenvFile()
	if err != nil {
		return "", err
	}
	if bazenv != nil {
		return *bazenv, nil
	}

	return "", errors.New("Could not find global or local " + BazenvFile)
}

// SetGlobalBazelVersion sets a global bazenv_version file in ~/.bazenv/bazenv_version
func SetGlobalBazelVersion(version string) error {
	homedir, err := homedir.Dir()
	if err != nil {
		return err
	}

	ioutil.WriteFile(filepath.Join(homedir, BazenvDir, BazenvFile), []byte(version), 0644)
	return nil
}

// SetLocalBazelVersion sets a local .bazenv_version file in CWD
func SetLocalBazelVersion(version string) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	ioutil.WriteFile(filepath.Join(currentDir, "."+BazenvFile), []byte(version), 0644)
	return nil
}

// FindAndReadLocalBazenvFile returns the contents of the local bazenv file, walking up the directory tree if needed
// to find one
func FindAndReadLocalBazenvFile() (*string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	for {
		// Try reading .bazenv_version
		versionName, err := ioutil.ReadFile(filepath.Join(currentDir, "."+BazenvFile))
		if err != nil && !os.IsNotExist(err) {
			return nil, err
		}

		if err == nil {
			versionNameString := strings.TrimSpace(string(versionName))
			return &versionNameString, nil
		}

		// File not found - go up a level, if we can
		if currentDir == "/" {
			// got to root directory, but nothing was found
			return nil, nil
		}
		currentDir = filepath.Dir(currentDir)
	}
}

// ReadGlobalBazenvFile returns the contents of global bazenv file
func ReadGlobalBazenvFile() (*string, error) {
	homedir, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	versionName, err := ioutil.ReadFile(filepath.Join(homedir, BazenvDir, BazenvFile))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	versionNameString := strings.TrimSpace(string(versionName))
	return &versionNameString, nil
}
