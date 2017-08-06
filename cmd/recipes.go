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

var (
	findName string
)

// recipeCmd represents the recipe command
var recipesCmd = &cobra.Command{
	Use:   "recipes",
	Short: "Manage the recipe repository",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		recipes := recipe.NewRecipes()
		if findName != "" {
			matches := recipes.SearchByName(findName)
			if len(matches) == 0 {
				color.Red("No recipes found for '%s'.", findName)
				return
			}
			if len(matches) == 1 {
				r := recipes.FindByKey(matches[0], 0)
				r.Print()
				return
			}
			color.White("Search results for '%s':", findName)
			for i, v := range matches {
				color.Green("%d. %s", i, v)
			}
			v := utils.AskForUserInput("Please select a result:")
			i, err := strconv.ParseInt(v, 10, 0)
			if err != nil {
				color.Red("Invalid value. %v", v)
				return
			}
			r := recipes.FindByKey(matches[i], 0)
			r.Print()
		}
	},
}

func init() {
	RootCmd.AddCommand(recipesCmd)
	recipesCmd.Flags().StringVarP(&findName, "search", "s", "", "Search a recipe by name")
}
