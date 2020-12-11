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
	"github.com/spf13/viper"
)

func init() {
	rootCommand.AddCommand(serverCommand)
}
  
var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "Start the Stee server",
	Long:  `Start the Stee server`,
	Run: ServerRun,
	Aliases: []string{"serve", "srv"},
}

// ServerConfig is the configuration for the Stee server
type ServerConfig struct {
	Server struct { 
	Address string
	Port string
	TLS struct {
		Enabled bool
		TLSCertPath string
		TLSKeyPath string
	}
	API struct {
		Enable bool
		URLPathPrefix string
		SimpleAPI struct {
			Enable bool
		}
	}
	UI struct {
		Enable bool
		URLPathPrefix string
	}}
}

// ServerRun runs a Stee server. Blocking.
func ServerRun(cmd *cobra.Command, args []string) {
	// Create configuration
	config := loadConfig().Server
	// Create core
	core := stee.NewCore()

	srv := steehttp.NewServer(steehttp.ServerConfig{
		ListenAddress: config.Address + ":" + config.Port,
		Handler: steehttp.RootHandler(core, config.API.URLPathPrefix, config.UI.URLPathPrefix),
	})

	// Starting to listen
	if config.TLS.Enabled {
		go func() {
			defer println("Stopped listening at https://" + srv.Addr)
			println("Listening at https://" + srv.Addr)
			err := srv.ListenAndServeTLS(config.TLS.TLSCertPath, config.TLS.TLSKeyPath)
			if err != http.ErrServerClosed {
				panic(fmt.Errorf("Server closed unexpectedly: %e", err))
			}
		}()
	} else {
		go func() {
			defer println("Stopped listening at http://" + srv.Addr)
			println("Listening at http://" + srv.Addr)
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

// loadConfig loads the config from a file
func loadConfig() ServerConfig {
	// We're looking for a file named "stee.yaml"
	viper.SetConfigName("stee")
	viper.SetConfigType("yaml")

	// In those directories
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath("/etc/stee/")

	// Environement variables take precedence over file config. See https://github.com/spf13/viper#why-viper
	viper.AutomaticEnv()

	// Defaults are the less important values. https://github.com/spf13/viper#why-viper
	setDefaults()

	err := viper.ReadInConfig()
	if err != nil {panic(fmt.Errorf("Fatal error config file: %s", err))}

	var cfg ServerConfig
	err = viper.Unmarshal(&cfg)
	if err != nil {panic(fmt.Errorf("unable to decode into struct, %s", err))}

	return cfg
}

// setDefaults sets the default config for Stee.
func setDefaults() {
	viper.SetDefault("server.address", "localhost")
	viper.SetDefault("server.port", "8008")

	viper.SetDefault("server.api.enable", true)
	viper.SetDefault("server.api.ReservedURLPrefix", "_api")
	viper.SetDefault("server.api.SimpleAPI.enable", true)

	viper.SetDefault("server.ui.enable", true)
}