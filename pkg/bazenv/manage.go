// Functions for managing the repository of bazel versions

package bazenv

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

// ListBazelVersions lists all bazel versions known to bazenv
func ListBazelVersions() ([]string, error) {
	homedir, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	bazelFiles, err := ioutil.ReadDir(filepath.Join(homedir, BazenvDir, BazenvVersionsDir))
	if err != nil {
		return nil, err
	}

	var bazelFileNames []string
	for _, bazelFile := range bazelFiles {
		bazelFileNames = append(bazelFileNames, bazelFile.Name())
	}
	return bazelFileNames, nil
}

// AddBazelVersion adds an existing bazel version (specified by path) to the set fo bazel versions known to bazenv
func AddBazelVersion(version, path string) error {
	homedir, err := homedir.Dir()
	if err != nil {
		return err
	}

	newpath := filepath.Join(homedir, BazenvDir, BazenvVersionsDir, version)
	return os.Symlink(path, newpath)
}

// RemoveBazelVersion removes a bazel version from the set of versions known to bazenv. If the version is a symlink
// the symlink is deleted. If the version is a directory, the entire tree is deleted.
func RemoveBazelVersion(version string) error {
	homedir, err := homedir.Dir()
	if err != nil {
		return err
	}

	path := filepath.Join(homedir, BazenvDir, BazenvVersionsDir, version)
	// Check for symlink
	fi, err := os.Lstat(path)
	if err != nil {
		return err
	}
	isSymlink := fi.Mode()&os.ModeSymlink != 0

	if isSymlink {
		return os.Remove(path)
	} else {
		return os.RemoveAll(path)
	}
}
