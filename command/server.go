package command

import (
	"net/http"
	"sync"

	steehttp "github.com/milanrodriguez/stee/http"
	"github.com/milanrodriguez/stee/stee"
	"github.com/spf13/cobra"
)

func init() {
	rootCommand.AddCommand(serverCommand)
  }
  
  var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "Start the Stee server",
	Long:  `Start the Stee server`,
	Run: ServerRun,
  }
  

// ServerConfig is the configuration for the Stee server
type ServerConfig struct {
	// Address on which the http server will listen. Defaults to ":8008".
	ListenAddress string

	// Prefix reserved for API paths. This string must end by a trailing slash. Defaults to "_api/". No leading slash required.
	APIReservedURLPrefix string

	// Flag to enable the web UI
	EnableUI bool

	// Prefix reserved for UI paths. This string must end by a trailing slash. Defaults to "_ui/". No leading slash required.
	UIReservedURLPrefix string
}

// ServerRun runs a Stee server. Blocking.
func ServerRun(cmd *cobra.Command, args []string) {

	// Create configuration

	cfg := DefaultConfig()

	// Start HTTP Listener

	core := stee.NewCore()
	handler := steehttp.RootHandler(core, cfg.APIReservedURLPrefix, cfg.UIReservedURLPrefix)

	var wg sync.WaitGroup
	wg.Add(1)
	go func () {
		defer wg.Done()
		defer println("Stopped listing at http://" + cfg.ListenAddress)
		http.ListenAndServe(cfg.ListenAddress, handler)
	}()
	println("Started listing at http://" + cfg.ListenAddress)
	
	wg.Wait()
}

// DefaultConfig returns the default config for Stee.
func DefaultConfig() ServerConfig {
	config := ServerConfig{
		ListenAddress: ":8008",
		APIReservedURLPrefix: "/_api/",
		EnableUI: true,
		UIReservedURLPrefix: "/_ui/",
	}

	return config
}