# Bazenv design doc

## Commands

* `bazenv` - Configures the bazel environment
  * `global name` - Sets the global bazel version
  * `local name` - Sets the local (this directory) bazel version by creating a `.bazenv_version` file
  * `add path [name]` - Adds an existing bazel install directory to the set of known bazel versions, optionally setting
    a name
  * `list` - Lists the set of known bazel version names
  * `available` - Lists the available versions of bazel that can be installed from bazel's github releases
  * `install name` - Downloads and installs a bazel version from bazel's github releases
  * `remove name` - Removes a bazel version from the set of known versions
* `bazel` - A shim that delegates to the chosen real `bazel` command, setting the correct `JAVA_HOME` environment
  variable and passing in the provided command line parameters

## File Structure

* `~/.bazenv` - User-specific bazenv config
  * `bazenv_version` - Contains the name of the global bazel version
  * `/bazel/[name]*` - One directory for each named bazel install (existing installs added with `add` are symlinks)
* `.bazenv_version` - Contains the name of the local bazel version

## TODO

* Implement basic commands
* Bash completions?
* ZSH completions?
