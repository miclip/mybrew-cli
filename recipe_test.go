package mybrewgo_test

import (
	. "github.com/miclip/mybrewgo"
	yaml "gopkg.in/yaml.v2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Recipe", func() {
	Context("Recipes can serialize and deserialize", func() {
		It("Unmarshals a recipe in yml", func() {

			recipeData := `---
recipe:
name: Test Recipe
batchsize: 5
batchunitofmeasure: gl
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

			var recipe Recipe
			err := yaml.Unmarshal([]byte(recipeData), &recipe)
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
				Hops: []Hop{
					Hop{
						Name:         "Galaxy",
						Amount:       1.25,
						Alpha:        13,
						Form:         "Pellet",
						Method:       "Boil",
						AdditionTime: 60,
					},
					Hop{
						Name:         "Centennial",
						Amount:       1.0,
						Alpha:        9.9,
						Form:         "Pellet",
						Method:       "Boil",
						AdditionTime: 10,
					},
					Hop{
						Name:         "Cascade",
						Amount:       1.0,
						Alpha:        6.7,
						Form:         "Pellet",
						Method:       "Boil",
						AdditionTime: 10,
					},
					Hop{
						Name:         "Centennial",
						Amount:       1.0,
						Alpha:        9.9,
						Form:         "Pellet",
						Method:       "Boil",
						AdditionTime: 0,
					},
					Hop{
						Name:         "Cascade",
						Amount:       1.0,
						Alpha:        6.7,
						Form:         "Pellet",
						Method:       "Boil",
						AdditionTime: 0,
					},
					Hop{
						Name:   "Citra",
						Amount: 1.0,
						Alpha:  12.0,
						Form:   "Pellet",
						Method: "Dry Hop",
					},
					Hop{
						Name:   "Galaxy",
						Amount: 1.0,
						Alpha:  13.0,
						Form:   "Pellet",
						Method: "Dry Hop",
					},
				},
				Fermentables: []Fermentable{
					Fermentable{
						Name:      "2 Row",
						Amount:    23.35,
						Potential: 1.036,
						Yield:     77.9,
						Lovibond:  2,
					},
					Fermentable{
						Name:      "Vienna Malt",
						Amount:    1.6,
						Potential: 1.036,
						Yield:     77.9,
						Lovibond:  4,
					},
					Fermentable{
						Name:      "White Wheat",
						Amount:    1.0,
						Potential: 1.04,
						Yield:     86.7,
						Lovibond:  2,
					},
				},
			}
		})
		It("Calculates BoilSpecificGravity", func() {
			Ω(Round(recipe.BoilSpecificGravity(), .5, 3)).Should(Equal(1.085))
		})
		It("Calculates OriginalGravity", func() {
			Ω(Round(recipe.OriginalGravity(), .5, 3)).Should(Equal(1.07))
		})
		It("Calculates 0.0 OriginalGravity when no efficiency provided", func() {
			recipe.Efficiency = 0
			Ω(Round(recipe.OriginalGravity(), .5, 3)).Should(Equal(0.0))
		})
		It("Calculates 0.0 OriginalGravity when no Batch size", func() {
			recipe.Batch = 0
			Ω(Round(recipe.OriginalGravity(), .5, 3)).Should(Equal(0.0))
		})
		It("Calculates EstimatedPreBoilVolume", func() {
			Ω(Round(recipe.EstimatedPreBoilVolume(), .5, 1)).Should(Equal(13.2))
		})
		It("Calculates InternationalBitteringUnits", func() {
			Ω(Round(recipe.InternationalBitteringUnits(), .5, 1)).Should(Equal(31.1))
		})
		It("Calculates Color SRM", func() {
			Ω(Round(recipe.Color(), .5, 1)).Should(Equal(9.4))
		})
		It("Calculates Color SRM of 0 when no fermentables", func() {
			recipe.Fermentables = []Fermentable{}
			Ω(Round(recipe.Color(), .5, 1)).Should(Equal(0.0))
		})
	})

})
