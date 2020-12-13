package command

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	steehttp "github.com/milanrodriguez/stee/http"
	"github.com/milanrodriguez/stee/stee"
	"github.com/spf13/cobra"
)

func init() {
	rootCommand.AddCommand(serverCommand)
}

var serverCommand = &cobra.Command{
	Use:     "server",
	Short:   "Start the Stee server",
	Long:    `Start the Stee server`,
	Run:     ServerRun,
	Aliases: []string{"serve", "srv"},
}

// ServerConfig is the configuration for the Stee server
type ServerConfig struct {
	Address string
	Port    string
	TLS     struct {
		Enable   bool
		CertPath string
		KeyPath  string
	}
	API struct {
		Enable        bool
		URLPathPrefix string
		SimpleAPI     struct {
			Enable bool
		}
	}
	UI struct {
		Enable        bool
		URLPathPrefix string
	}
}

// ServerRun runs a Stee server. Blocking.
func ServerRun(cmd *cobra.Command, args []string) {
	// Create configuration
	config := loadConfig().Server

	// Create core
	core := stee.NewCore()

	srv := steehttp.NewServer(steehttp.ServerConfig{
		ListenAddress: config.Address + ":" + config.Port,
		Handler: steehttp.HandleRoot(
			steehttp.Core(core),
			steehttp.EnableAPI(config.API.Enable, config.API.URLPathPrefix),
			steehttp.EnableSimpleAPI(config.API.SimpleAPI.Enable),
			steehttp.EnableUI(config.UI.Enable, config.UI.URLPathPrefix),
		),
	})

	// Starting to listen
	if config.TLS.Enable {
		go func() {
			defer fmt.Printf("Stopped listening at %s\n", srv.Addr)
			fmt.Printf("Listening at https://%s/\n", srv.Addr)
			err := srv.ListenAndServeTLS(config.TLS.CertPath, config.TLS.KeyPath)
			if err != http.ErrServerClosed {
				panic(fmt.Errorf("Server closed unexpectedly: %e", err))
			}
		}()
	} else {
		go func() {
			defer fmt.Printf("Stopped listening at %s\n", srv.Addr)
			fmt.Printf("\n⚠️ You are running the server without TLS encryption!\n⚠️ You should consider setting up HTTPS for production use.\n\n")
			fmt.Printf("Listening at http://%s/\n", srv.Addr)
			err := srv.ListenAndServe()
			if err != http.ErrServerClosed {
				panic(fmt.Errorf("Server closed unexpectedly: %e", err))
			}
		}()
	}

	//////////////////////////////////////////////////////////////////
	// At this point initialization is done, the server is running. //
	// We're now listening for OS signals                           //
	//////////////////////////////////////////////////////////////////

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for range c {
		fmt.Printf("\nInterruption requested. We're gonna perform a clean shutdown...\n")
		srv.Shutdown(context.Background())
		core.Close()
		fmt.Printf("Bye!\n")
		break
	}

}
