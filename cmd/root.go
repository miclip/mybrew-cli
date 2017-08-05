package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/miclip/mybrewgo/recipe"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

var cfgFile, projectBase, userLicense string

// Recipes ...
var Recipes map[string]*recipe.Recipe

// RootCmd ...
var RootCmd = &cobra.Command{
	Use:   "recipe",
	Short: "Recipe is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

// SaveRecipes ...
func SaveRecipes() error {
	var r []*recipe.Recipe
	for _, v := range Recipes {
		r = append(r, v)
	}
	data, err := yaml.Marshal(r)
	if err != nil {
		return err
	}
	ioutil.WriteFile(recipeFilepath(), data, 777)
	return nil
}

func recipeFilepath() string {
	return filepath.Join("./", "mybrewgo_recipes.yml")
}

// GetRecipes ...
func GetRecipes() error {
	if Recipes == nil {
		Recipes = make(map[string]*recipe.Recipe)
	}
	f := recipeFilepath()
	data, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}
	var r []*recipe.Recipe
	err = yaml.Unmarshal(data, &r)
	if err != nil {
		return err
	}
	for _, v := range r {
		Recipes[recipeKey(v)] = v
	}
	return nil
}

func recipeKey(recipe *recipe.Recipe) string {
	return fmt.Sprintf("%s\\%v", recipe.Name, recipe.Version)
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	RootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
	RootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
	RootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	RootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	viper.BindPFlag("author", RootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("projectbase", RootCmd.PersistentFlags().Lookup("projectbase"))
	viper.BindPFlag("useViper", RootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "apache")
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

		// Search config in home directory with name ".mybrewgo" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".mybrewgo")
	}

	// if err := viper.ReadInConfig(); err != nil {
	// 	fmt.Println("Can't read config:", err)
	// 	os.Exit(1)
	// }
}
