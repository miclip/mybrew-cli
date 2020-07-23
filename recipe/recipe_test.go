package recipe_test

import (
	"time"

	"github.com/miclip/mybrew/fakes"
	"github.com/miclip/mybrew/ingredients"
	. "github.com/miclip/mybrew/recipe"
	"github.com/miclip/mybrew/ui"
	"github.com/miclip/mybrew/utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Recipe", func() {
	Context("Recipes can serialize and deserialize", func() {
		It("Unmarshals a recipe in yml", func() {

			recipeData := `---
recipe:
name: Test Recipe
batchsize: 5
efficiency: 70
method: All Grain
boiltime: 60
hops:
- name: Cascade
  amount: 1
  alpha: 6.7
  form: Pellet
  method: boil
fermentables:
- name: 2 Row
  amount: 5
  yield: 77.9
  potential: 1.036
  lovibond: 2
  type: Grain
- name: Crystal 10
  amount: 5
  yield: 77.9
  potential: 1.036
  lovibond: 2
  type: Grain
yeasts:
- name: California Ale
  attenutation: 85`

			var recipe *Recipe
			recipe, err := UnmarshalRecipeYML([]byte(recipeData))
			Ω(err).Should(Succeed())
			Ω(recipe).ShouldNot(BeNil())
			Ω(recipe.Name).Should(Equal("Test Recipe"))
			Ω(len(recipe.Fermentables)).Should(Equal(2))
			Ω(recipe.Fermentables[0].Amount).Should(Equal(5.0))
			Ω(recipe.Fermentables[0].Yield).Should(Equal(77.9))
			Ω(recipe.Fermentables[0].Potential).Should(Equal(1.036))
			Ω(len(recipe.Hops)).Should(Equal(1))
			Ω(recipe.Hops[0].Amount).Should(Equal(1.0))
			Ω(recipe.Hops[0].Alpha).Should(Equal(6.7))
			Ω(len(recipe.Yeasts)).Should(Equal(1))
			Ω(recipe.Yeasts[0].Name).Should(Equal("California Ale"))
		})
		It("Unmarshals a recipe in json", func() {

			recipeData := `{"Name":"Accidental IPA","Version":0,"Batch":11,"Style":"","Efficiency":83,"Method":"","BoilTime":90,"Hops":[{"Name":"Galaxy","Alpha":13,"Amount":1.25,"Form":"Pellet","Method":"Boil","AdditionTime":60},{"Name":"Centennial","Alpha":9.9,"Amount":1,"Form":"Pellet","Method":"Boil","AdditionTime":10},{"Name":"Cascade","Alpha":6.7,"Amount":1,"Form":"Pellet","Method":"Boil","AdditionTime":10},{"Name":"Centennial","Alpha":9.9,"Amount":1,"Form":"Pellet","Method":"Boil","AdditionTime":0},{"Name":"Cascade","Alpha":6.7,"Amount":1,"Form":"Pellet","Method":"Boil","AdditionTime":0},{"Name":"Citra","Alpha":12,"Amount":1,"Form":"Pellet","Method":"Dry Hop","AdditionTime":0},{"Name":"Galaxy","Alpha":13,"Amount":1,"Form":"Pellet","Method":"Dry Hop","AdditionTime":0}],"Fermentables":[{"Name":"2 Row","Amount":23.35,"Yield":77.9,"Potential":1.036,"Lovibond":2,"Type":""},{"Name":"Vienna Malt","Amount":1.6,"Yield":77.9,"Potential":1.036,"Lovibond":4,"Type":""},{"Name":"White Wheat","Amount":1,"Yield":86.7,"Potential":1.04,"Lovibond":2,"Type":""}],"Yeasts":[{"Name":"Safale American","Attenuation":77}]}`

			var recipe *Recipe
			recipe, err := UnmarshalRecipeJSON([]byte(recipeData))
			Ω(err).Should(Succeed())
			Ω(recipe).ShouldNot(BeNil())
			Ω(recipe.Name).Should(Equal("Accidental IPA"))
			Ω(len(recipe.Fermentables)).Should(Equal(3))
			Ω(recipe.Fermentables[0].Amount).Should(Equal(23.35))
			Ω(recipe.Fermentables[0].Yield).Should(Equal(77.9))
			Ω(recipe.Fermentables[0].Potential).Should(Equal(1.036))
			Ω(len(recipe.Hops)).Should(Equal(7))
			Ω(recipe.Hops[0].Amount).Should(Equal(1.25))
			Ω(recipe.Hops[0].Alpha).Should(Equal(13.0))
			Ω(len(recipe.Yeasts)).Should(Equal(1))
			Ω(recipe.Yeasts[0].Name).Should(Equal("Safale American"))
		})
		It("Unmarshal fails to read a recipe in yml", func() {

			recipeData := `---
recipe:
INVALID`

			var recipe *Recipe
			recipe, err := UnmarshalRecipeYML([]byte(recipeData))
			Ω(err).ShouldNot(Succeed())
			Ω(recipe).Should(BeNil())
		})
		It("Can open a yml file and parse a recipe", func() {
			fileName := "../test_data/accidental-ipa.yml"
			recipe, err := OpenRecipe(fileName)
			Ω(err).Should(Succeed())
			Ω(recipe).ShouldNot(BeNil())
		})
		It("Can open a json file and parse a recipe", func() {
			fileName := "../test_data/dry-irish-stout.json"
			recipe, err := OpenRecipe(fileName)
			Ω(err).Should(Succeed())
			Ω(recipe).ShouldNot(BeNil())
		})
		It("Cannot open a file", func() {
			fileName := "../test_data/does-not-exist.yml"
			recipe, err := OpenRecipe(fileName)
			Ω(err).ShouldNot(Succeed())
			Ω(recipe).Should(BeNil())
		})
	})
	Context("Recipe Calculations", func() {
		var (
			recipe Recipe
		)
		BeforeEach(func() {
			recipe = Recipe{
				Name:       "Accidental IPA",
				Batch:      11,
				Efficiency: 83,
				BoilTime:   90,
				Hops: []ingredients.Hop{
					ingredients.Hop{
						Name:         "Galaxy",
						Amount:       1.25,
						Alpha:        13,
						Form:         "Pellet",
						Method:       "Boil",
						AdditionTime: 60,
					},
					ingredients.Hop{
						Name:         "Centennial",
						Amount:       1.0,
						Alpha:        9.9,
						Form:         "Pellet",
						Method:       "Boil",
						AdditionTime: 10,
					},
					ingredients.Hop{
						Name:         "Cascade",
						Amount:       1.0,
						Alpha:        6.7,
						Form:         "Pellet",
						Method:       "Boil",
						AdditionTime: 10,
					},
					ingredients.Hop{
						Name:         "Centennial",
						Amount:       1.0,
						Alpha:        9.9,
						Form:         "Pellet",
						Method:       "Boil",
						AdditionTime: 0,
					},
					ingredients.Hop{
						Name:         "Cascade",
						Amount:       1.0,
						Alpha:        6.7,
						Form:         "Pellet",
						Method:       "Boil",
						AdditionTime: 0,
					},
					ingredients.Hop{
						Name:   "Citra",
						Amount: 1.0,
						Alpha:  12.0,
						Form:   "Pellet",
						Method: "Dry Hop",
					},
					ingredients.Hop{
						Name:   "Galaxy",
						Amount: 1.0,
						Alpha:  13.0,
						Form:   "Pellet",
						Method: "Dry Hop",
					},
				},
				Fermentables: []ingredients.Fermentable{
					ingredients.Fermentable{
						Name:      "2 Row",
						Amount:    23.35,
						Potential: 1.036,
						Yield:     77.9,
						Lovibond:  2,
					},
					ingredients.Fermentable{
						Name:      "Vienna Malt",
						Amount:    1.6,
						Potential: 1.036,
						Yield:     77.9,
						Lovibond:  4,
					},
					ingredients.Fermentable{
						Name:      "White Wheat",
						Amount:    1.0,
						Potential: 1.04,
						Yield:     86.7,
						Lovibond:  2,
					},
				},
				Yeasts: []ingredients.Yeast{
					ingredients.Yeast{
						Name:        "Safale American",
						Attenuation: 77,
					},
				},
			}
		})
		It("Calculates BoilSpecificGravity", func() {
			Ω(utils.Round(recipe.BoilSpecificGravity(), .5, 3)).Should(Equal(1.085))
		})
		It("Calculates OriginalGravity", func() {
			Ω(utils.Round(recipe.OriginalGravity(), .5, 3)).Should(Equal(1.07))
		})
		It("Calculates OriginalGravityPoints", func() {
			Ω(utils.Round(recipe.OriginalGravityPoints(), .5, 3)).Should(Equal(70.47))
		})
		It("Calculates 0.0 OriginalGravity when no efficiency provided", func() {
			recipe.Efficiency = 0
			Ω(utils.Round(recipe.OriginalGravity(), .5, 3)).Should(Equal(0.0))
		})
		It("Calculates 0.0 OriginalGravity when no Batch size", func() {
			recipe.Batch = 0
			Ω(utils.Round(recipe.OriginalGravity(), .5, 3)).Should(Equal(0.0))
		})
		It("Calculates EstimatedPreBoilVolume", func() {
			Ω(utils.Round(recipe.EstimatedPreBoilVolume(), .5, 1)).Should(Equal(13.2))
		})
		It("Calculates EstimatedFinalGravity", func() {
			Ω(utils.Round(recipe.EstimatedFinalGravity(), .5, 3)).Should(Equal(1.016))
		})
		It("Calculates AlcoholByVolume", func() {
			Ω(utils.Round(recipe.AlcoholByVolume(), .5, 2)).Should(Equal(7.10))
		})
		It("Calculates AlcoholByWeight", func() {
			Ω(utils.Round(recipe.AlcoholByWeight(), .5, 2)).Should(Equal(5.52))
		})
		It("Calculates InternationalBitteringUnits", func() {
			Ω(utils.Round(recipe.InternationalBitteringUnits(), .5, 1)).Should(Equal(37.8))
		})
		It("Calculates Color SRM", func() {
			Ω(utils.Round(recipe.Color(), .5, 1)).Should(Equal(9.4))
		})
		It("Calculates Color SRM of 0 when no fermentables", func() {
			recipe.Fermentables = []ingredients.Fermentable{}
			Ω(utils.Round(recipe.Color(), .5, 1)).Should(Equal(0.0))
		})
	})

	Context("UI", func() {
		var (
			bOut *gbytes.Buffer
			bErr *gbytes.Buffer
			bIn  *gbytes.Buffer
			ui   ui.UI
			r    Recipes
		)
		BeforeEach(func() {
			bOut, bErr, bIn = gbytes.NewBuffer(), gbytes.NewBuffer(), gbytes.NewBuffer()
			_, _ = gbytes.TimeoutWriter(bOut, time.Second*2), gbytes.TimeoutWriter(bErr, time.Second)
			_ = gbytes.TimeoutReader(bIn, time.Second)
			ui = fakes.NewFakeUI(bOut, bErr, bIn)
			r = NewRecipes(ui)
		})
		AfterEach(func() {
			r.DeleteRecipes()
		})
		It("Prints a recipe", func() {
			r := Recipe{
				Name:       "Accidental IPA",
				Style:      "India Pale Ale",
				Batch:      11,
				Efficiency: 83,
				BoilTime:   90,
				Hops: []ingredients.Hop{
					ingredients.Hop{
						Name:         "Galaxy",
						Amount:       1.25,
						Alpha:        13,
						Form:         "Pellet",
						Method:       "Boil",
						AdditionTime: 60,
					},
					ingredients.Hop{
						Name:         "Centennial",
						Amount:       1.0,
						Alpha:        9.9,
						Form:         "Pellet",
						Method:       "Boil",
						AdditionTime: 10,
					},
					ingredients.Hop{
						Name:         "Cascade",
						Amount:       1.0,
						Alpha:        6.7,
						Form:         "Pellet",
						Method:       "Boil",
						AdditionTime: 10,
					},
					ingredients.Hop{
						Name:         "Centennial",
						Amount:       1.0,
						Alpha:        9.9,
						Form:         "Pellet",
						Method:       "Boil",
						AdditionTime: 0,
					},
					ingredients.Hop{
						Name:         "Cascade",
						Amount:       1.0,
						Alpha:        6.7,
						Form:         "Pellet",
						Method:       "Boil",
						AdditionTime: 0,
					},
					ingredients.Hop{
						Name:   "Citra",
						Amount: 1.0,
						Alpha:  12.0,
						Form:   "Pellet",
						Method: "Dry Hop",
					},
					ingredients.Hop{
						Name:   "Galaxy",
						Amount: 1.0,
						Alpha:  13.0,
						Form:   "Pellet",
						Method: "Dry Hop",
					},
				},
				Fermentables: []ingredients.Fermentable{
					ingredients.Fermentable{
						Name:      "2 Row",
						Amount:    23.35,
						Potential: 1.036,
						Yield:     77.9,
						Lovibond:  2,
						Type:      "Grain",
					},
					ingredients.Fermentable{
						Name:      "Vienna Malt",
						Amount:    1.6,
						Potential: 1.036,
						Yield:     77.9,
						Lovibond:  4,
						Type:      "Grain",
					},
					ingredients.Fermentable{
						Name:      "White Wheat",
						Amount:    1.0,
						Potential: 1.04,
						Yield:     86.7,
						Lovibond:  2,
						Type:      "Grain",
					},
				},
				Yeasts: []ingredients.Yeast{
					ingredients.Yeast{
						Name:        "Safale American",
						Attenuation: 77,
					},
				},
			}

			r.Print(ui)
			Ω(bOut).To(gbytes.Say("Recipe: Accidental IPA Version: 0\nStyle: India Pale Ale\nBatch Size: 11 Boil Time: 90\nOG: 1.07 FG: 1.016 IBU: 37.8 ABV: 7.1 SRM: 9.4\nFermentables: \n2 Row Amount: 23.4 Yield: 77.9 Potential: 1.036 Lovibond: 2 Type: Grain\nVienna Malt Amount: 1.6 Yield: 77.9 Potential: 1.036 Lovibond: 4 Type: Grain\nWhite Wheat Amount: 1 Yield: 86.7 Potential: 1.04 Lovibond: 2 Type: Grain\nHops: \nGalaxy Amount: 1.25 Time: 60 Alpha: 13 Form: Pellet Method: Boil\nCentennial Amount: 1 Time: 10 Alpha: 9.9 Form: Pellet Method: Boil\nCascade Amount: 1 Time: 10 Alpha: 6.7 Form: Pellet Method: Boil\nCentennial Amount: 1 Time: 0 Alpha: 9.9 Form: Pellet Method: Boil\nCascade Amount: 1 Time: 0 Alpha: 6.7 Form: Pellet Method: Boil\nCitra Amount: 1 Time: 0 Alpha: 12 Form: Pellet Method: Dry Hop\nGalaxy Amount: 1 Time: 0 Alpha: 13 Form: Pellet Method: Dry Hop\nYeasts: \nSafale American Attenuation: 77"))
		})
		It("creates a new recipe interactively", func() {
			ui.SetMaxInvalidInput(1)
			bIn.Write([]byte("Test Name\n1\n11.0\n61.5\n71.0\nAll Grain\nTest Style\nf\nTest Fermentable\n12.0\n1.036\n77.9\n2\nGrain\nh\nTest Hop\n1.25\n12\n35\nPellet\nBoil\ny\nTest Yeast\n77.9\ns\n"))
			recipe, err := CreateInteractively(ui)
			Ω(bErr).ToNot(gbytes.Say("Error reading input."))
			Ω(recipe).ShouldNot(BeNil())
			Ω(err).Should(Succeed())
			Ω(recipe.Name).Should(Equal("Test Name"))
			Ω(recipe.Version).Should(Equal(1))
			Ω(recipe.Batch).Should(Equal(11.0))
			Ω(recipe.BoilTime).Should(Equal(61.5))
			Ω(recipe.Efficiency).Should(Equal(71.0))
			Ω(recipe.Method).Should(Equal("All Grain"))
			Ω(recipe.Style).Should(Equal("Test Style"))
			Ω(len(recipe.Fermentables)).Should(Equal(1))
			Ω(recipe.Fermentables[0].Name).Should(Equal("Test Fermentable"))
			Ω(recipe.Fermentables[0].Amount).Should(Equal(12.0))
			Ω(recipe.Fermentables[0].Potential).Should(Equal(1.036))
			Ω(recipe.Fermentables[0].Yield).Should(Equal(77.9))
			Ω(recipe.Fermentables[0].Lovibond).Should(Equal(2.0))
			Ω(recipe.Fermentables[0].Type).Should(Equal("Grain"))
			Ω(len(recipe.Hops)).Should(Equal(1))
			Ω(recipe.Hops[0].Name).Should(Equal("Test Hop"))
			Ω(recipe.Hops[0].Amount).Should(Equal(1.25))
			Ω(recipe.Hops[0].Alpha).Should(Equal(12.0))
			Ω(recipe.Hops[0].AdditionTime).Should(Equal(35))
			Ω(recipe.Hops[0].Form).Should(Equal("Pellet"))
			Ω(recipe.Hops[0].Method).Should(Equal("Boil"))
			Ω(recipe.Yeasts[0].Name).Should(Equal("Test Yeast"))
			Ω(recipe.Yeasts[0].Attenuation).Should(Equal(77.9))
		})
	})
})
