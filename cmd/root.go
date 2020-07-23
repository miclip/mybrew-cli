package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile, projectBase, userLicense string

// RootCmd ...
var RootCmd = &cobra.Command{
	Use:   "mybrew",
	Short: "mybrew-cli is a fast command line interface for managing homebrew recipes.",
	Long: `mybrew-cli is a fast command line interface for managing homebrew recipes. mybrew-cli
	supports recipes in either YAML, JSON, XML and can be added directly via the cli.

	Recipes are stored local to the executable in the human readable YAML format. This enables
	the user to choose a source code repository like github.com to store and backup your recipes.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	RootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "./")
	RootCmd.PersistentFlags().StringP("author", "a", "Michael Lipscombe", "Author name for copyright attribution")
	RootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "MIT License")
	RootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	viper.BindPFlag("author", RootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("projectbase", RootCmd.PersistentFlags().Lookup("projectbase"))
	viper.BindPFlag("useViper", RootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "Michael Lipscombe michael@mybrewco.com")
	viper.SetDefault("license", "MIT")
}

// Execute ...
func Execute() {
	RootCmd.Execute()
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".mybrew" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".mybrew")
	}

	// if err := viper.ReadInConfig(); err != nil {
	// 	fmt.Println("Can't read config:", err)
	// 	os.Exit(1)
	// }
}
