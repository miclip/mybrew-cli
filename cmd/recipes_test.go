package cmd

import (
	"path/filepath"
	"time"

	"github.com/miclip/mybrewgo/fakes"
	"github.com/miclip/mybrewgo/ui"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Recipes Cmd", func() {
	Context("recipes commands", func() {
		var (
			bOut *gbytes.Buffer
			bErr *gbytes.Buffer
			err  error
			ui   ui.UI
		)
		BeforeEach(func() {
			bOut, bErr = gbytes.NewBuffer(), gbytes.NewBuffer()
			_, _ = gbytes.TimeoutWriter(bOut, time.Second), gbytes.TimeoutWriter(bOut, time.Second)
			ui = fakes.NewFakeUI(bOut, bErr)
		})
		It("lists all recipes in the local repository", func() {
			path, err = filepath.Abs("../test_data/accidental-ipa.yml")
			Ω(err).Should(Succeed())
			handleAdd(nil, ui)
			Ω(bOut).To(gbytes.Say("Adding Recipe..."))
			handleRecipes(nil, ui)
			Ω(bOut).To(gbytes.Say("\nRecipes:\n0. Accidental IPA"))
		})

	})
})
