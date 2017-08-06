package recipe

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"

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

// SearchByName
func (r *Recipes) SearchByName(name string) *Recipe {
	var names []string
	for _, v := range r.Recipes {
		names = append(names, v.Name)
	}
	sort.Strings(names)
	index := sort.SearchStrings(names, name)
	log.Printf("Search Index %v", index)
	return r.FindByKey(names[index-1], 0)
}

// SaveRecipes ...
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

func (r *Recipes) RecipeKey(recipe *Recipe) string {
	return r.recipeKeyByName(recipe.Name, recipe.Version)
}

func (r *Recipes) recipeKeyByName(name string, version int) string {
	return fmt.Sprintf("%s\\%v", name, version)
}
