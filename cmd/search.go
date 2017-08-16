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

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for a item",
	Long: `Search for a Recipe or Ingredient in the local repository
	by Name.`,
	Run: func(cmd *cobra.Command, args []string) {
		ui := ui.NewConsoleUI()
		handleSearch(args, ui)
	},
}

func handleSearch(args []string, ui ui.UI) {
	recipes := recipe.NewRecipes(ui)
	if len(args) == 0 {
		ui.ErrorLinef("No search arguments provided.")
		return
	}
	findName := args[0]
	matches := recipes.SearchByName(findName)
	if len(matches) == 0 {
		ui.ErrorLinef("No recipes found for '%s'.", findName)
		return
	}
	if len(matches) == 1 {
		ui.SystemLinef("One recipe found, displaying recipe:")
		r := recipes.FindByKey(matches[0], 0)
		r.Print(ui)
		return
	}
	ui.SystemLinef("Search results for '%s', %d recipes found:", findName, len(matches))
	ui.DisplayColumns(matches, 3)
	i := ui.AskForInt("Please select a result:")
	r := recipes.FindByKey(matches[i], 0)
	r.Print(ui)
}

func init() {
	recipesCmd.AddCommand(searchCmd)
}
