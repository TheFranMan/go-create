package main

import "regexp"

var versionRegex = regexp.MustCompile(`^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)
var nameRegex = regexp.MustCompile(`^[a-zA-Z0-9\-_]+$`)

func validateVersion(version string) bool {
	return versionRegex.Match([]byte(version))
}

func validateName(name string) bool {
	return nameRegex.Match([]byte(name))
}
