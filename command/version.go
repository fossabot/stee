package command

import (
	"os"
	"text/template"
	"time"

	steeversion "github.com/milanrodriguez/stee/internal/version"
	"github.com/spf13/cobra"
)

type versionInfo struct {
	Version        string
	GitCommit      string
	BuildTime      string
	BuildType      string
	BuildPlatform  string
	AdditionalInfo string
}

var version versionInfo

func init() {
	version = versionInfo{
		Version:       steeversion.Version(),
		GitCommit:     steeversion.GitCommit(),
		BuildTime:     steeversion.BuildTime().UTC().Format(time.RFC3339),
		BuildType:     steeversion.BuildType(),
		BuildPlatform: steeversion.BuildPlatform(),
	}

	rootCommand.AddCommand(versionCommand)
}

const versionTemplate string = `stee version
version:   {{.Version}} ({{.GitCommit}})
platform:  {{.BuildPlatform}}
build:     {{.BuildType}} ({{.BuildTime}})
`

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Print the version of Stee",
	Long: `Print the version of Stee in the following format:
	
  version:   version (gitCommit)
  platform:  platform
  build:     buildType (buildDate)

  Stee versioning follows Semantic Versioning 2.0.0 (https://semver.org/)

  The buildType can be one of those:
    release:
      This build corresponds to a release.
      It does not mean that this is an official build. Check checksums for that.
    pre-release:
      This build corresponds to a pre-release.
      It does not mean that this is an official build. Check checksums for that.
    development (local HEAD):
      This build was made out of the last commit (local).
    development:
      This build was made out of untracked changes.
`,
	Run: func(cmd *cobra.Command, args []string) {
		t, err := template.New("version").Parse(versionTemplate)
		if err != nil {
			panic(err)
		}
		err = t.Execute(os.Stdout, version)
		if err != nil {
			panic(err)
		}
	},
	Aliases: []string{"v"},
}
