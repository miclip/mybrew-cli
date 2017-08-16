package ingredients_test

import (
	"time"

	"github.com/miclip/mybrewgo/fakes"
	"github.com/miclip/mybrewgo/hoputils"
	. "github.com/miclip/mybrewgo/ingredients"
	"github.com/miclip/mybrewgo/ui"
	"github.com/miclip/mybrewgo/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Hop", func() {
	It("Calculates International Bittering Units ", func() {
		h := Hop{
			Name:         "Galaxy",
			Amount:       1.25,
			Alpha:        13,
			Form:         "Pellet",
			Method:       "Boil",
			AdditionTime: 60,
		}
		hopUtils := hoputils.NewHopUtilizations()
		ibu := h.InternationalBitteringUnits(hopUtils, 11, 1.07)
		Ω(utils.Round(ibu, .5, 1)).Should(Equal(28.8))
	})
	It("Calculates zero International Bittering Units with a non boil addition ", func() {
		h := Hop{
			Name:         "Galaxy",
			Amount:       1.25,
			Alpha:        13,
			Form:         "Pellet",
			Method:       "Dry Hop",
			AdditionTime: 12,
		}
		hopUtils := hoputils.NewHopUtilizations()
		ibu := h.InternationalBitteringUnits(hopUtils, 11, 1.07)
		Ω(ibu).Should(Equal(0.0))
	})
	Context("UI", func() {
		var (
			bOut *gbytes.Buffer
			bErr *gbytes.Buffer
			ui   ui.UI
		)
		BeforeEach(func() {
			bOut, bErr = gbytes.NewBuffer(), gbytes.NewBuffer()
			_, _ = gbytes.TimeoutWriter(bOut, time.Second), gbytes.TimeoutWriter(bOut, time.Second)
			ui = fakes.NewFakeUI(bOut, bErr)
		})
		It("Prints a Hop", func() {
			h := Hop{
				Name:         "Galaxy",
				Amount:       1.25,
				Alpha:        13,
				Form:         "Pellet",
				Method:       "Dry Hop",
				AdditionTime: 12,
			}
			h.Print(ui)
			Ω(bOut).To(gbytes.Say("Galaxy Amount: 1.25 Time: 12 Alpha: 13 Form: Pellet Method: Dry Hop"))
		})
	})
})
