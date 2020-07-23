package cmd

import (
	"path/filepath"
	"time"

	"github.com/miclip/mybrew/fakes"
	"github.com/miclip/mybrew/recipe"
	"github.com/miclip/mybrew/ui"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Recipes Cmd", func() {
	Context("recipes commands", func() {
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
		It("lists all recipes in the local repository", func() {
			path, err = filepath.Abs("../test_data/accidental-ipa.yml")
			Ω(err).Should(Succeed())
			handleAdd(nil, ui)
			Ω(bOut).To(gbytes.Say("Add Recipe..."))
			handleRecipes(nil, ui)
			Ω(bOut).To(gbytes.Say("\nRecipes:\n0. Accidental IPA"))
		})

	})
})
