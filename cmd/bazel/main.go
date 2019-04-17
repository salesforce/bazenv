package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/ryanmichela/bazenv/pkg/bazenv"
)

func main() {
	// Capture the location of the Java 8 JDK and set JAVA_HOME, if possible
	jdk, err := exec.Command("/usr/libexec/java_home", "--version", "1.8").Output()
	if err == nil {
		os.Setenv("JAVA_HOME", strings.TrimSpace(string(jdk)))
	}

	bazelProfile, err := bazenv.ReadBazenvFile()
	check(err)
	bazelDir, err := bazenv.ResolveBazelDirectory(bazelProfile)
	check(err)

	// Execute the selected real bazel entry point
	// binary := "/usr/local/Cellar/bazel/0.23.2/libexec/bin/bazel"
	binary := filepath.Join(bazelDir, "libexec", "bin", "bazel")
	args := os.Args
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	check(execErr)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
