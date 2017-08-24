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

var _ = Describe("Recipe Cmd", func() {
	Context("recipe commands", func() {
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
			_, _ = gbytes.TimeoutWriter(bOut, time.Second), gbytes.TimeoutWriter(bOut, time.Second)
			_ = gbytes.TimeoutReader(bIn, time.Second)
			ui = fakes.NewFakeUI(bOut, bErr, bIn)
			r = recipe.NewRecipes(ui)
		})
		AfterEach(func() {
			r.DeleteRecipes()
		})
		It("finds a recipe by name and version", func() {
			path, err = filepath.Abs("../test_data/accidental-ipa.yml")
			Ω(err).Should(Succeed())
			handleAdd(nil, ui)
			name = "Accidental IPA"
			version = 0
			handleRecipe(nil, ui)
			Ω(bOut).To(gbytes.Say("Add Recipe...\nRecipe: Accidental IPA Version: 0\nStyle: American IPA\nBatch Size: 11 Boil Time: 90\nOG: 1.07 FG: 1.016 IBU: 37.8 ABV: 7.1 SRM: 9.4\n"))
		})
		It("fails when a name/version cannot be found", func() {
			path, err = filepath.Abs("../test_data/accidental-ipa.yml")
			Ω(err).Should(Succeed())
			handleAdd(nil, ui)
			name = "Doesn't Exist"
			version = 0
			handleRecipe(nil, ui)
			Ω(bErr).To(gbytes.Say("Recipe 'Doesn't Exist' version 0 not found."))
		})
	})
})
