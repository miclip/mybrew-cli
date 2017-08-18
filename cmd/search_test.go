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

var _ = Describe("Search Cmd", func() {
	Context("search commands", func() {
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
		It("fails when no arguments are provided", func() {
			handleSearch(nil, ui)
			Ω(bErr).To(gbytes.Say("No search arguments provided."))
		})
		It("fails when no recipes are found", func() {
			args := make([]string, 1)
			args[0] = "test"
			handleSearch(args, ui)
			Ω(bErr).To(gbytes.Say("No recipes found for 'test'."))
		})
		It("finds one recipe", func() {
			args := make([]string, 1)
			path, err = filepath.Abs("../test_data/accidental-ipa.yml")
			Ω(err).Should(Succeed())
			handleAdd(nil, ui)
			Ω(bOut).To(gbytes.Say("Adding Recipe...\n"))
			args[0] = "acc"
			handleSearch(args, ui)
			Ω(bOut).To(gbytes.Say("One recipe found, displaying recipe:\nRecipe: Accidental IPA Version: 0\nStyle: American IPA\nBatch Size: 11 Boil Time: 90"))
		})
		It("finds multiple recipes", func() {
			args := make([]string, 1)
			path, err = filepath.Abs("../test_data/accidental-ipa.yml")
			Ω(err).Should(Succeed())
			handleAdd(nil, ui)
			path, err = filepath.Abs("../test_data/czech-pilsner.yml")
			Ω(err).Should(Succeed())
			handleAdd(nil, ui)
			args[0] = "c"
			handleSearch(args, ui)
			Ω(bOut).To(gbytes.Say("Search results for 'c', 2 recipes found:"))
		})

	})
})
