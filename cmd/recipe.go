// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/fatih/color"
	"github.com/miclip/mybrewgo/recipe"
	"github.com/spf13/cobra"
)

var (
	name    string
	version int
)

// recipeCmd represents the recipe command
var recipeCmd = &cobra.Command{
	Use:   "recipe",
	Short: "View and Manage a recipe store in the local repository ",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		recipes := recipe.NewRecipes()
		if name != "" {
			r := recipes.FindByKey(name, version)
			if r == nil {
				color.Red("Recipe '%s' version %v not found.", name, version)
				return
			}
			r.Print()
		}
	},
}

func init() {
	RootCmd.AddCommand(recipeCmd)
	recipeCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the recipe")
	recipeCmd.Flags().IntVarP(&version, "version", "v", 0, "Version of the recipe")
}
