package bazenv

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/mitchellh/go-homedir"
)

// ResolveBazelDirectory converts a bazel version name into the path to a bazel install directory, or retuns an error
// if the install doesn't exist.
func ResolveBazelDirectory(version string) (string, error) {
	homedir, err := homedir.Dir()
	check(err)

	bazelDir := filepath.Join(homedir, BazenvDir, BazenvVersionsDir, version)
	if _, err := os.Stat(bazelDir); os.IsNotExist(err) {
		// bazelDir doesn't exist
		return "", errors.New("Bazel version '" + version + "' does not exist. Use 'bazenv global' to configure a" +
			" global bazel version.")
	}

	// Try to dereference bazelDir, if it is a symlink
	bazelDir, err = os.Readlink(bazelDir)

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
