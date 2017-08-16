package cmd

import (
	"time"

	"github.com/miclip/mybrewgo/fakes"
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
			//err  error
			ui ui.UI
		)
		BeforeEach(func() {
			bOut, bErr = gbytes.NewBuffer(), gbytes.NewBuffer()
			_, _ = gbytes.TimeoutWriter(bOut, time.Second), gbytes.TimeoutWriter(bOut, time.Second)
			ui = fakes.NewFakeUI(bOut, bErr)
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

	})
})
