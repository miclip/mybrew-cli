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

var _ = Describe("Add", func() {
	Context("Add recipes from external files", func() {
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
	})

})
