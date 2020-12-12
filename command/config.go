package command

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	configCommand.AddCommand(configShowCommand)
	rootCommand.AddCommand(configCommand)
}

var configCommand = &cobra.Command{
	Use:   "config",
	Short: "Configuration of Stee",
	Long:  `Configuration of Stee`,
}

var configShowCommand = &cobra.Command{
	Use:   "showjson",
	Short: "Shows the current configuration of Stee",
	Long:  `Shows the current configuration of Stee`,
	Run: func(cmd *cobra.Command, args []string) {
		json, _ := json.MarshalIndent(loadConfig(), "", "    ")
		fmt.Printf("%s\n", string(json))
	},
}

// Config is the global configuration struct
type Config struct {
	Server ServerConfig
}

// loadConfig loads the config from a file
func loadConfig() Config {
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
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %s", err))
	}

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
	viper.SetDefault("server.api.ReservedURLPrefix", "_ui")
}
