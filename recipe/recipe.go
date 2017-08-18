package recipe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
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

// Create adds a recipe via the cli
func Create(ui ui.UI) *Recipe {
	recipe := &Recipe{}
	recipe.Name = ui.AskForText("Recipe Name:")
	recipe.Version, _ = ui.AskForInt("Version Number:")
	recipe.Batch = ui.AskForFloat("Batch Size:")
	recipe.BoilTime = ui.AskForFloat("Boil Time:")
	recipe.Efficiency = ui.AskForFloat("Efficiency:")
	recipe.Method = ui.AskForText("Method:")
	recipe.Style = ui.AskForText("Style:")
	color.White("Ingredients...")
	addIngedients := true
	for addIngedients {
		i := ui.AskForText("Add Ingredient, Fermentable (f), Hop (h), Yeast (y), Save/Exit (s):")
		if i == "s" {
			break
		}
		if strings.ToLower(i) != "f" && strings.ToLower(i) != "h" && strings.ToLower(i) != "y" {
			color.Red("Invalid Ingredient, must be either h,f,y or s")
		}
	}
	return recipe
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
