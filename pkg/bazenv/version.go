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
	bazenvDir  = ".bazenv"
	bazenvFile = "bazenv_version"
)

// ReadBazenvFile reads the content of .bazenv_version, looking locally up the directiory tree first, then in the
// global bazenv file. This returns the name of the active bazel profile.
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

	return "", errors.New("Could not find global or local " + bazenvFile)
}

// ResolveBazelDirectory converts a bazel profile name into the path to a bazel install directory, or retuns an error
// if the install doesn't exist.
func ResolveBazelDirectory(profile string) (string, error) {
	homedir, err := homedir.Dir()
	check(err)

	bazelDir := filepath.Join(homedir, bazenvDir, "bazel", profile)
	if _, err := os.Stat(bazelDir); os.IsNotExist(err) {
		// bazelDir doesn't exist
		return "", errors.New("Bazel profile " + profile + " does not exist")
	}

	return bazelDir, nil
}

func findAndReadLocalBazenvFile() (*string, error) {
	currentDir, err := os.Getwd()
	check(err)

	for {
		// Try reading .bazenv_version
		versionName, err := ioutil.ReadFile(filepath.Join(currentDir, "."+bazenvFile))
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

	versionName, err := ioutil.ReadFile(filepath.Join(homedir, bazenvDir, bazenvFile))
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
