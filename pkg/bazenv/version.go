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
	// BazenvFile is the name of a bazenv version file, prefixed with dot when not in BazenvDir
	BazenvFile = "bazenv_version"
)

// ReadBazenvFile reads the content of .bazenv_version, looking locally up the directiory tree first, then in the
// global bazenv file. This returns the name of the active bazel version.
func ReadBazenvFile() (string, error) {
	// Local bazenv file takes priority
	bazenv, err := findAndReadLocalBazenvFile()
	if err != nil {
		return "", err
	}
	if bazenv != nil {
		return *bazenv, nil
	}

	// Fall back to global bazenv
	bazenv, err = readGlobalBazenvFile()
	if err != nil {
		return "", err
	}
	if bazenv != nil {
		return *bazenv, nil
	}

	return "", errors.New("Could not find global or local " + BazenvFile)
}

// ResolveBazelDirectory converts a bazel version name into the path to a bazel install directory, or retuns an error
// if the install doesn't exist.
func ResolveBazelDirectory(version string) (string, error) {
	homedir, err := homedir.Dir()
	check(err)

	bazelDir := filepath.Join(homedir, BazenvDir, "bazel", version)
	if _, err := os.Stat(bazelDir); os.IsNotExist(err) {
		// bazelDir doesn't exist
		return "", errors.New("Bazel version " + version + " does not exist")
	}

	return bazelDir, nil
}

// SetGlobalBazelVersion sets a global bazenv_version file in ~/.bazenv/bazenv_version
func SetGlobalBazelVersion(version string) {
	homedir, err := homedir.Dir()
	check(err)

	ioutil.WriteFile(filepath.Join(homedir, BazenvDir, BazenvFile), []byte(version), 0644)
}

// SetLocalBazelVersion sets a local .bazenv_version file in CWD
func SetLocalBazelVersion(version string) {
	currentDir, err := os.Getwd()
	check(err)

	ioutil.WriteFile(filepath.Join(currentDir, "."+BazenvFile), []byte(version), 0644)
}

func findAndReadLocalBazenvFile() (*string, error) {
	currentDir, err := os.Getwd()
	check(err)

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

func readGlobalBazenvFile() (*string, error) {
	homedir, err := homedir.Dir()
	check(err)

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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
