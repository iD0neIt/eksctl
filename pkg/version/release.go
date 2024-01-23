package version

// This file was generated by release_generate.go; DO NOT EDIT.

// Version is the version number in semver format X.Y.Z
var Version = "0.170.0"

// PreReleaseID can be empty for releases, "rc.X" for release candidates and "dev" for snapshots
var PreReleaseID = "dev"

// gitCommit is the short commit hash. It will be set by the linker.
var gitCommit = ""

// buildDate is the time of the build with format yyyy-mm-ddThh:mm:ssZ. It will be set by the linker.
var buildDate = ""
