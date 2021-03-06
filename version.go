// WARNING: THIS FILE IS AUTOMATICALLY GENERATED BY THE MAKEFILE
//
// Example:
//
// LDFLAGS=-ldflags "-X rithium.uk/stor/version.Version=${VERSION} -X rithium.uk/stor/version.Build=${BUILD}"
package version

import "fmt"

var (
	Version string
	Build	string
)

const versionFormat = "v%s-%s"

func GetVersion() (string) {
	return fmt.Sprintf(versionFormat, Version, Build)
}
