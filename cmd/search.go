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
	"strconv"

	"github.com/fatih/color"
	"github.com/miclip/mybrewgo/recipe"
	"github.com/miclip/mybrewgo/utils"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for a item",
	Long: `Search for a Recipe or Ingredient in the local repository
	by Name.`,
	Run: func(cmd *cobra.Command, args []string) {
		recipes := recipe.NewRecipes()

		if len(args) == 0 {
			color.Red("No search arguments provided.")
			return
		}
		findName := args[0]
		matches := recipes.SearchByName(findName)
		if len(matches) == 0 {
			color.Red("No recipes found for '%s'.", findName)
			return
		}
		if len(matches) == 1 {
			color.White("One recipe found, displaying recipe:")
			r := recipes.FindByKey(matches[0], 0)
			r.Print()
			return
		}
		color.White("Search results for '%s', %d recipes found:", findName, len(matches))
		for i, v := range matches {
			color.Green("%d. %s", i, v)
		}
		v := utils.RequestUserInput("Please select a result:")
		i, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			color.Red("Invalid value. %v", v)
			return
		}
		r := recipes.FindByKey(matches[i], 0)
		r.Print()
	},
}

func init() {
	recipesCmd.AddCommand(searchCmd)
}
