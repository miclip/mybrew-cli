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
	"github.com/miclip/mybrewgo/recipe"
	"github.com/miclip/mybrewgo/ui"
	"github.com/spf13/cobra"
)

var (
	path string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a recipe to the internal repo.",
	Long: `Adds a recipe from an external yaml file to the internal
	repository. Calculations are displayed.`,
	Run: func(cmd *cobra.Command, args []string) {
		ui := ui.NewConsoleUI()
		handleAdd(args, ui)
	},
}

func handleAdd(args []string, ui ui.UI) {
	recipes := recipe.NewRecipes(ui)
	ui.SystemLinef("Adding Recipe...")
	if path != "" {
		r, err := recipe.OpenRecipe(path)
		if err != nil {
			ui.ErrorLinef("Error adding recipe: %v", err)
			return
		}
		if recipes.Recipes[recipes.RecipeKey(r)] != nil {
			ui.ErrorLinef("A Recipe %v already exists, increment the version number.", recipes.RecipeKey(r))
			return
		}
		recipes.AddRecipe(r)
		err = recipes.SaveRecipes()
		if err != nil {
			ui.ErrorLinef("Error saving recipe store with %v", err)
		}
		r.Print(ui)
		return
	}

	r, _ := recipe.CreateInteractively(ui)
	recipes.AddRecipe(r)
	// err := recipes.SaveRecipes()
	// if err != nil {
	// 	color.Red("Error saving recipe store with %v", err)
	// }
	r.Print(ui)
}

func init() {
	recipesCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&path, "path", "p", "", "Help message for path")

}
