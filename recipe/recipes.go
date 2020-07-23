package recipe

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/miclip/mybrew/ui"

	yaml "gopkg.in/yaml.v2"
)

// Recipes ...
type Recipes struct {
	Recipes map[string]*Recipe
	ui      ui.UI
}

// NewRecipes ...
func NewRecipes(ui ui.UI) Recipes {
	r := Recipes{ui: ui}
	r.GetRecipes()
	return r
}

// FindByKey ..
func (r *Recipes) FindByKey(name string, version int) *Recipe {
	return r.Recipes[r.recipeKeyByName(name, version)]
}

// SearchByName ...
func (r *Recipes) SearchByName(name string) []string {
	var matches []string
	for _, v := range r.Recipes {
		if strings.Contains(strings.ToLower(v.Name), strings.TrimSpace(strings.ToLower(name))) {
			matches = append(matches, v.Name)
		}
	}
	return matches
}

// SaveRecipes internal recipes to disk
func (r *Recipes) SaveRecipes() error {
	data, err := yaml.Marshal(r.Recipes)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(r.recipeFilepath(), data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// DeleteRecipes deletes the local repository
func (r *Recipes) DeleteRecipes() error {
	var err = os.Remove(r.recipeFilepath())
	if err != nil {
		return err
	}
	return nil
}

// AddRecipe adds a recipe to the internal repository
func (r *Recipes) AddRecipe(recipe *Recipe) {
	r.Recipes[r.RecipeKey(recipe)] = recipe
}

func (r *Recipes) recipeFilepath() string {
	return filepath.Join("./", ".recipes.yml")
}

// GetRecipes ...
func (r *Recipes) GetRecipes() error {
	if r.Recipes == nil {
		r.Recipes = make(map[string]*Recipe)
	}
	f := r.recipeFilepath()
	data, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &r.Recipes)
	if err != nil {
		return err
	}
	return nil
}

// GetRecipeNames returns a slice containing the key (name/version)
func (r *Recipes) GetRecipeNames() []string {
	var names []string
	for k := range r.Recipes {
		names = append(names, k)
	}
	return names
}

// RecipeKey returns the unique key for a recipe
func (r *Recipes) RecipeKey(recipe *Recipe) string {
	return r.recipeKeyByName(recipe.Name, recipe.Version)
}

func (r *Recipes) recipeKeyByName(name string, version int) string {
	return fmt.Sprintf("%s\\%v", name, version)
}
