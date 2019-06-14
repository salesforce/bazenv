# Master your Bazel environment with Bazenv

Bazel is fiddley. Projects that compile with one version of Bazel don't work with another, and keeping multiple
versions of Bazel around is a pain. Bazenv solves this problem by managing your Bazel version on a project by project
basis. Bazenv is inspired by `jenv` and `rbenv`, which solve the same problem for Java and Ruby.

## Installation

1. Bazenv is built with go. To download: `go get -d github.com/salesforce/bazenv`
1. Install Bazenv with `make deps install`. This will build `bazenv` and its `bazel` stub.
1. Make sure `$GOHOME/bin` is at the beginning of your path.
1. Install a version of bazel with `bazenv install <version>`.
1. Set the global Bazel version with `bazenv global <version>`.

## Installing additional Bazel versions

* Use `bazenv install` to download and install any Bazel version on
  [Bazel's Github releases page](https://github.com/bazelbuild/bazel/releases).
* Use `bazenv add` to add an existing Bazel install directory to Bazenv.

## Choosing which Bazel version to run

* To see all avaliable Bazel versions, use `bazenv list`.
* Use `bazenv global` to set the global Bazel version. This version will be used unless specifically overridden with
  a local version.
* Use `bazenv local` to set a local Bazel version for a given directory. This version will be used for all child
  directories.

## How Bazenv works

* Bazenv stores configuration and all installed bazel versions in the `~/.bazenv` directory.
* `.bazenv_version` files are used to configure the local Bazel version.
* The Bazenv `bazel` shim reads from Bazenv's configuration to locate the correct Bazel binary and execute it.

## Troubleshooting

* Make sure Bazenv's `bazel` shim is before any other `bazel` binary on your path.
* `bazenv doctor` will evaluate your Bazenv environment.
