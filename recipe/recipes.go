package recipe

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

// Recipes ...
type Recipes struct {
	Recipes map[string]*Recipe
}

// NewRecipes ...
func NewRecipes() Recipes {
	r := Recipes{}
	r.GetRecipes()
	return r
}

// FindByKey ..
func (r *Recipes) FindByKey(name string, version int) *Recipe {
	return r.Recipes[r.recipeKeyByName(name, version)]
}

// SaveRecipes ...
func (r *Recipes) SaveRecipes() error {
	var ra []*Recipe
	for _, v := range r.Recipes {
		ra = append(ra, v)
	}
	data, err := yaml.Marshal(r)
	if err != nil {
		return err
	}
	ioutil.WriteFile(r.recipeFilepath(), data, 777)
	return nil
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
	var ra []*Recipe
	err = yaml.Unmarshal(data, &ra)
	if err != nil {
		return err
	}
	for _, v := range ra {
		r.Recipes[r.RecipeKey(v)] = v
	}
	return nil
}

func (r *Recipes) RecipeKey(recipe *Recipe) string {
	return r.recipeKeyByName(recipe.Name, recipe.Version)
}

func (r *Recipes) recipeKeyByName(name string, version int) string {
	return fmt.Sprintf("%s\\%v", name, version)
}
