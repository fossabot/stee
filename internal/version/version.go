package version

import (
	"time"
)

var version string = "undefined-version"

var gitCommit string = "undefined-gitcommit"

var buildTime string = "undefined-buildtime"

var buildType string = "undifined-buildtype"

var buildPlatform string = "undifined-buildarch"

// Version returns Stee version with "v" appended (vx.y.z).
func Version() string {
	return version
}

// GitCommit returns Stee last commit hash.
func GitCommit() string {
	return gitCommit
}

// BuildTime returns Stee build time.
func BuildTime() time.Time {
	parsedBuildTime, _ := time.Parse(time.RFC3339, buildTime)
	return parsedBuildTime
}

// BuildType returns Stee build type.
// release, release-candidate, devel
func BuildType() string {
	return buildType
}

// BuildPlatform returns the build os/arch.
func BuildPlatform() string {
	return buildPlatform
}
