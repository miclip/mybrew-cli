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
	"fmt"
	"log"

	"github.com/miclip/mybrewgo/recipe"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a recipe to the internal repo.",
	Long: `Adds a recipe from an external yaml file to the internal
	repository. Calculations are performed and displayed.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Recipe Add...")
		r, err := recipe.OpenRecipe(args[0])
		if err != nil {
			log.Printf("Error opening recipe file with: %v", err)
		}
		r.Print()
	},
}

func init() {
	recipeCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("path", "p", "", "Help message for path")
}
