package cmd

import (
	"path/filepath"
	"time"

	"github.com/miclip/mybrewgo/fakes"
	"github.com/miclip/mybrewgo/recipe"
	"github.com/miclip/mybrewgo/ui"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Add", func() {
	Context("Add recipes from external files", func() {
		var (
			bOut *gbytes.Buffer
			bErr *gbytes.Buffer
			bIn  *gbytes.Buffer
			err  error
			ui   ui.UI
			r    recipe.Recipes
		)
		BeforeEach(func() {
			bOut, bErr, bIn = gbytes.NewBuffer(), gbytes.NewBuffer(), gbytes.NewBuffer()
			_, _ = gbytes.TimeoutWriter(bOut, time.Second), gbytes.TimeoutWriter(bErr, time.Second)
			_ = gbytes.TimeoutReader(bIn, time.Second)
			ui = fakes.NewFakeUI(bOut, bErr, bIn)
			r = recipe.NewRecipes(ui)
		})
		AfterEach(func() {
			r.DeleteRecipes()
		})
		It("fails when path argument not valid", func() {
			path = "../invalid/error.json"
			handleAdd(nil, ui)
			Ω(bErr).To(gbytes.Say("Error adding recipe: failed to read file with:"))
		})
		It("fails to add a duplicate recipe", func() {
			path, err = filepath.Abs("../test_data/accidental-ipa.yml")
			Ω(err).Should(Succeed())
			handleAdd(nil, ui)
			handleAdd(nil, ui)
			Ω(bErr).To(gbytes.Say("already exists, increment the version number."))
		})
		It("adds a new recipe interactively", func() {
			path = ""
			ui.SetMaxInvalidInput(1)
			bIn.Write([]byte("Test Name\n1\n11.0\n61.5\n71.0\nAll Grain\nTest Style\nf\nTest Fermentable\n12.0\n1.036\n77.9\n2\nGrain\nh\nTest Hop\n1.25\n12\n35\nPellet\nBoil\ny\nTest Yeast\n77.9\ns\n"))
			handleAdd(nil, ui)
			r = recipe.NewRecipes(ui)
			Ω(bErr).ToNot(gbytes.Say("Error reading input."))
			recipe := r.FindByKey("Test Name", 1)
			Ω(recipe).ShouldNot(BeNil())
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
