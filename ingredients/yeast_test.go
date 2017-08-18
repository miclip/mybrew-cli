package ingredients_test

import (
	"time"

	"github.com/miclip/mybrewgo/fakes"
	. "github.com/miclip/mybrewgo/ingredients"
	"github.com/miclip/mybrewgo/ui"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Yeast", func() {
	Context("UI", func() {
		var (
			bOut *gbytes.Buffer
			bErr *gbytes.Buffer
			ui   ui.UI
		)
		BeforeEach(func() {
			bOut, bErr = gbytes.NewBuffer(), gbytes.NewBuffer()
			_, _ = gbytes.TimeoutWriter(bOut, time.Second), gbytes.TimeoutWriter(bOut, time.Second)
			ui = fakes.NewFakeUI(bOut, bErr, nil)
		})
		It("Prints a Hop", func() {
			y := Yeast{
				Name:        "California Ale",
				Attenuation: 77.9,
			}
			y.Print(ui)
			Ω(bOut).To(gbytes.Say("California Ale Attenuation: 77.9"))
		})
	})
})
