package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"
)

const versionPkgImportPath string = "github.com/milanrodriguez/stee/internal/version"

// https://semver.org/#is-there-a-suggested-regular-expression-regex-to-check-a-semver-string
const semVerRegex string = "^(?P<major>0|[1-9]\\d*)\\.(?P<minor>0|[1-9]\\d*)\\.(?P<patch>0|[1-9]\\d*)(?:-(?P<prerelease>(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\.(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\\.[0-9a-zA-Z-]+)*))?$"

func main() {
	associations := []string{
		"version=" + getVersion(),
		"gitCommit=" + getGitCommit(),
		"buildTime=" + getBuildTime(),
		"buildType=" + getBuildType(),
		"buildPlatform=" + getBuildPlatform(),
	}

	varInjects := ""
	for _, association := range associations {
		fmt.Println(association)
		if strings.HasSuffix(association, "=") {
			// If we don't have a value for the variable, we keep it as is.
			continue
		}
		varInjects += "-X '" + versionPkgImportPath + "." + association + "' "
	}

	command := "go build -mod vendor -o bin/stee -ldflags \"-w -s " + varInjects + "\" ."
	_, err := run(command)
	if err != nil {
		panic(err)
	}
	out, err := run("./bin/stee version")
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n✔ Succesful build:\n\n")
	fmt.Printf("%v", out)

}

func getVersion() string {
	out, err := run("git for-each-ref --sort=-creatordate --format '%(refname:short)' refs/tags")
	if err != nil {
		println(out)
		panic(err)
	}
	lines := strings.Split(out, "\n")
	semVerRegexMatcher := regexp.MustCompile(semVerRegex)
	for _, line := range lines {
		if !strings.HasPrefix(line, "v") {
			continue
		}
		line = strings.TrimSuffix(line, "\n")
		if semVerRegexMatcher.MatchString(strings.TrimPrefix(line, "v")) {
			return line
		}
	}
	return ""
}

func getGitCommit() string {
	out, err := run("git rev-parse --short HEAD")
	if err != nil {
		panic(err)
	}
	return strings.TrimSuffix(out, "\n")
}

func getBuildTime() string {
	return time.Now().UTC().Format(time.RFC3339)
}

// Release if tag is semver without prerelease and 'git diff tag..HEAD' == ""
// Prerelease if tag is semver with prerelease and 'git diff tag..HEAD' == ""
// Development-HEAD if 'git diff HEAD' == ""
// Development if 'git diff HEAD' != ""
func getBuildType() string {
	version := getVersion()
	prerelease := getPrerelease(version)
	gitDiffFromVersion, err := run("git diff " + getVersion())
	if err != nil {
		panic(err)
	}
	gitDiffFromHEAD, err := run("git diff " + getVersion())
	if err != nil {
		panic(err)
	}

	if gitDiffFromVersion == "" {
		if prerelease == "" {
			return "release"
		}
		return "pre-release"
	}
	if gitDiffFromHEAD == "" {
		return "development (local HEAD)"
	}
	return "development"
}

func getPrerelease(version string) string {
	prereleaseStart := strings.Index(version, "-")
	buildMetadataStart := strings.Index(version, "+")
	s := prereleaseStart
	e := buildMetadataStart
	if prereleaseStart == -1 {
		s = 0
	}
	if buildMetadataStart == -1 {
		e = len(version)
	}
	return version[s:e]
}

func getBuildPlatform() string {
	// Todo change that: innacurate in case of crosscompilation
	return runtime.GOOS + "/" + runtime.GOARCH
}

func run(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Run()

	if stderr.String() != "" {
		return stdout.String(), fmt.Errorf("%v", stderr.String())
	}
	return stdout.String(), nil
}
