package bazenv

import (
	"errors"
	"os"
	"path/filepath"

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

	return bazelDir, nil
}
