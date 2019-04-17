package main

import (
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
	// Capture the location of the Java 8 JDK and set JAVA_HOME
	jdk, err := exec.Command("/usr/libexec/java_home", "--version", "1.8").Output()
	if err != nil {
		panic(err)
	}
	os.Setenv("JAVA_HOME", strings.TrimSpace(string(jdk)))

	// Execute the selected real bazel entry point
	binary := "/usr/local/Cellar/bazel/0.23.2/libexec/bin/bazel"
	args := os.Args
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
