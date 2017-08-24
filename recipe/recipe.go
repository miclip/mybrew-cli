package recipe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/miclip/mybrewgo/hoputils"
	"github.com/miclip/mybrewgo/ingredients"
	"github.com/miclip/mybrewgo/ui"
	"github.com/miclip/mybrewgo/utils"

	yaml "gopkg.in/yaml.v2"
)

const (
	yieldToPoints           = 0.46
	gravityBase             = 0.00001
	preBoilGallonsFactorEst = 1.2
)

// Recipe ...
type Recipe struct {
	Name         string
	Version      int
	Batch        float64
	Style        string
	Efficiency   float64
	Method       string
	BoilTime     float64
	Hops         []ingredients.Hop
	Fermentables []ingredients.Fermentable
	Yeasts       []ingredients.Yeast
}

// CreateInteractively adds a recipe via the cli
func CreateInteractively(ui ui.UI) (*Recipe, error) {
	recipe := &Recipe{}
	name, err := ui.AskForText("Recipe Name:")
	if err != nil {
		return nil, fmt.Errorf("Invalid recipe name with: %v", err)
	}
	recipe.Name = name
	version, err := ui.AskForInt("Version Number:")
	if err != nil {
		return nil, fmt.Errorf("Invalid recipe version with: %v", err)
	}
	recipe.Version = version
	batch, err := ui.AskForFloat("Batch Size:")
	if err != nil {
		return nil, fmt.Errorf("Invalid recipe batch with: %v", err)
	}
	recipe.Batch = batch
	boilTime, err := ui.AskForFloat("Boil Time:")
	if err != nil {
		return nil, fmt.Errorf("Invalid recipe boil time with: %v", err)
	}
	recipe.BoilTime = boilTime
	efficiency, err := ui.AskForFloat("Efficiency:")
	if err != nil {
		return nil, fmt.Errorf("Invalid recipe efficiency with: %v", err)
	}
	recipe.Efficiency = efficiency
	method, err := ui.AskForText("Method:")
	if err != nil {
		return nil, fmt.Errorf("Invalid recipe brewing method with: %v", err)
	}
	recipe.Method = method
	style, err := ui.AskForText("Style:")
	if err != nil {
		return nil, fmt.Errorf("Invalid recipe style with: %v", err)
	}
	recipe.Style = style
	ui.SystemLinef("Ingredients...")
	addIngedients := true
	for addIngedients {
		i, _ := ui.AskForText("Add Ingredient, Fermentable (f), Hop (h), Yeast (y), Save & Exit (s):")
		if i == "s" {
			break
		}
		if strings.ToLower(i) != "f" && strings.ToLower(i) != "h" && strings.ToLower(i) != "y" {
			ui.ErrorLinef("Invalid Ingredient, must be either h, f, y or s")
		}
		if strings.ToLower(i) == "f" {
			if recipe.Fermentables == nil {
				recipe.Fermentables = []ingredients.Fermentable{}
			}
			f, _ := ingredients.CreateFermentableInteractively(ui)
			recipe.Fermentables = append(recipe.Fermentables, *f)
		}
		if strings.ToLower(i) == "h" {
			if recipe.Hops == nil {
				recipe.Hops = []ingredients.Hop{}
			}
			h, _ := ingredients.CreateHopInteractively(ui)
			recipe.Hops = append(recipe.Hops, *h)
		}
		if strings.ToLower(i) == "y" {
			if recipe.Yeasts == nil {
				recipe.Yeasts = []ingredients.Yeast{}
			}
			y, _ := ingredients.CreateYeastInteractively(ui)
			recipe.Yeasts = append(recipe.Yeasts, *y)
		}
	}
	return recipe, nil
}

// OpenRecipe ...
func OpenRecipe(fileName string) (recipe *Recipe, err error) {
	filePath, _ := filepath.Abs(fileName)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file with: %v", err)
	}
	if strings.Contains(strings.ToLower(fileName), ".yml") {
		return UnmarshalRecipeYML(data)
	}
	return UnmarshalRecipeJSON(data)
}

// UnmarshalRecipeYML ...
func UnmarshalRecipeYML(recipeYamlData []byte) (recipe *Recipe, err error) {
	err = yaml.Unmarshal(recipeYamlData, &recipe)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal yml recipe data with, %v", err)
	}
	return recipe, nil
}

// UnmarshalRecipeJSON ...
func UnmarshalRecipeJSON(recipeJSONData []byte) (recipe *Recipe, err error) {
	err = json.Unmarshal(recipeJSONData, &recipe)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal json recipe data with, %v", err)
	}
	return recipe, nil
}

// EstimatedPreBoilVolume estimates the preboil volume
func (r *Recipe) EstimatedPreBoilVolume() float64 {
	return r.Batch * preBoilGallonsFactorEst
}

// OriginalGravity calculates the original gravity
func (r *Recipe) OriginalGravity() float64 {
	og := 0.0
	if r.Efficiency == 0 || r.Batch == 0 || len(r.Fermentables) == 0 {
		return og
	}
	for i := range r.Fermentables {
		og = og +
			r.Fermentables[i].PointsByAmount()*(r.Efficiency*gravityBase)
	}
	return og/r.Batch + 1
}

// OriginalGravityPoints converts OG in Specific Gravity to Points
func (r *Recipe) OriginalGravityPoints() float64 {
	return (r.OriginalGravity() - 1) * 1000
}

// BoilSpecificGravity calculates the specific gravity post boil
func (r *Recipe) BoilSpecificGravity() float64 {
	return r.EstimatedPreBoilVolume()/r.Batch*(r.OriginalGravity()-1) + 1
}

// Color for recipe in SRM
func (r *Recipe) Color() float64 {
	color := 0.0
	if len(r.Fermentables) == 0 {
		return color
	}
	for i := range r.Fermentables {
		color = color + r.Fermentables[i].ColorMCU()
	}
	return ((color / r.Batch) * 0.2) + 8.4
}

// InternationalBitteringUnits calculates IBU for the recipe using a variation on Tinseth’s formula that
// incorporates a gravity/time adjustment instead of the bigness factor as documented by Randy Mosher
// in the "Brewer's Companion". No IBU formula is perfect so expect variations.
func (r *Recipe) InternationalBitteringUnits() float64 {
	hopUtils := hoputils.NewHopUtilizations()
	ibu := 0.0
	if len(r.Hops) == 0 {
		return 0.0
	}
	for i := range r.Hops {
		ibu = ibu + r.Hops[i].InternationalBitteringUnits(hopUtils, r.Batch, r.OriginalGravity())
	}
	return ibu
}

// EstimatedFinalGravity ...
func (r *Recipe) EstimatedFinalGravity() float64 {
	averageAttenuation := 0.0
	if len(r.Yeasts) == 0 {
		return 0.0
	}
	for i := range r.Yeasts {
		averageAttenuation = averageAttenuation + r.Yeasts[i].Attenuation/100
	}
	averageAttenuation = averageAttenuation / float64(len(r.Yeasts))

	return r.OriginalGravityPoints()*(1-averageAttenuation)/1000 + 1
}

// AlcoholByVolume ...
func (r *Recipe) AlcoholByVolume() float64 {
	if r.EstimatedFinalGravity() == 0 {
		return 0.0
	}
	return ((1.05 * (r.OriginalGravity() - r.EstimatedFinalGravity())) / r.EstimatedFinalGravity()) / 0.79 * 100
}

// AlcoholByWeight ...
func (r *Recipe) AlcoholByWeight() float64 {
	if r.EstimatedFinalGravity() == 0 {
		return 0.0
	}
	return (0.79 * r.AlcoholByVolume()) / r.EstimatedFinalGravity()
}

// Print writes recipe and ingredient details to stdout
func (r *Recipe) Print(ui ui.UI) {
	ui.PrintLinef("Recipe: %s Version: %d", r.Name, r.Version)
	ui.PrintLinef("Style: %s", r.Style)
	ui.PrintLinef("Batch Size: %v Boil Time: %v", r.Batch, r.BoilTime)
	ui.PrintLinef("OG: %v FG: %v IBU: %v ABV: %v SRM: %v", utils.Round(r.OriginalGravity(), .5, 3), utils.Round(r.EstimatedFinalGravity(), .5, 3),
		utils.Round(r.InternationalBitteringUnits(), .5, 1), utils.Round(r.AlcoholByVolume(), .5, 1), utils.Round(r.Color(), .5, 1))
	if len(r.Fermentables) > 0 {
		ui.PrintLinef("Fermentables: ")
	}
	for _, v := range r.Fermentables {
		v.Print(ui)
	}
	if len(r.Hops) > 0 {
		ui.PrintLinef("Hops: ")
	}
	for _, v := range r.Hops {
		v.Print(ui)
	}
	if len(r.Yeasts) > 0 {
		ui.PrintLinef("Yeasts: ")
	}
	for _, v := range r.Yeasts {
		v.Print(ui)
	}
}
