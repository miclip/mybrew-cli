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
			bIn  *gbytes.Buffer
			ui   ui.UI
		)
		BeforeEach(func() {
			bOut, bErr, bIn = gbytes.NewBuffer(), gbytes.NewBuffer(), gbytes.NewBuffer()
			_, _ = gbytes.TimeoutWriter(bOut, time.Second*2), gbytes.TimeoutWriter(bErr, time.Second)
			_ = gbytes.TimeoutReader(bIn, time.Second)
			ui = fakes.NewFakeUI(bOut, bErr, bIn)
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
		It("successfully creates a hop interactively", func() {
			ui.SetMaxInvalidInput(1)
			bIn.Write([]byte("Test Hop\n1.25\n12\n35\nPellet\nBoil\n"))
			hop, err := CreateHopInteractively(ui)
			Ω(bErr).ToNot(gbytes.Say("Error reading input."))
			Ω(hop).ShouldNot(BeNil())
			Ω(err).Should(Succeed())
			Ω(hop.Name).Should(Equal("Test Hop"))
			Ω(hop.Amount).Should(Equal(1.25))
			Ω(hop.Alpha).Should(Equal(12.0))
			Ω(hop.AdditionTime).Should(Equal(35))
			Ω(hop.Form).Should(Equal("Pellet"))
			Ω(hop.Method).Should(Equal("Boil"))
		})
		It("unsuccessfully creates a hop interactively due to invalid name", func() {
			ui.SetMaxInvalidInput(1)
			bIn.Write([]byte(nil))
			hop, err := CreateHopInteractively(ui)
			Ω(bErr).To(gbytes.Say("Error reading input."))
			Ω(hop).Should(BeNil())
			Ω(err).ShouldNot(Succeed())
		})
		It("unsuccessfully creates a hop interactively due to invalid amount", func() {
			ui.SetMaxInvalidInput(1)
			bIn.Write([]byte("Test Hop\nxx\n12\n35\nPellet\nBoil\n"))
			hop, err := CreateHopInteractively(ui)
			Ω(bErr).To(gbytes.Say("Invalid float, please enter a valid value."))
			Ω(hop).Should(BeNil())
			Ω(err).ShouldNot(Succeed())
		})
		It("unsuccessfully creates a hop interactively due to invalid alpha", func() {
			ui.SetMaxInvalidInput(1)
			bIn.Write([]byte("Test Hop\n1.25\nxx\n35\nPellet\nBoil\n"))
			hop, err := CreateHopInteractively(ui)
			Ω(bErr).To(gbytes.Say("Invalid float, please enter a valid value."))
			Ω(hop).Should(BeNil())
			Ω(err).ShouldNot(Succeed())
		})
		It("unsuccessfully creates a hop interactively due to invalid addition time", func() {
			ui.SetMaxInvalidInput(1)
			bIn.Write([]byte("Test Hop\n1.25\n12\nxx\nPellet\nBoil\n"))
			hop, err := CreateHopInteractively(ui)
			Ω(bErr).To(gbytes.Say("Invalid int, please enter a valid value."))
			Ω(hop).Should(BeNil())
			Ω(err).ShouldNot(Succeed())
		})
		It("unsuccessfully creates a hop interactively due to invalid type", func() {
			ui.SetMaxInvalidInput(1)
			bIn.Write([]byte("Test Hop\n1.25\n12\n35\n"))
			hop, err := CreateHopInteractively(ui)
			Ω(bErr).To(gbytes.Say("Error reading input."))
			Ω(hop).Should(BeNil())
			Ω(err).ShouldNot(Succeed())
		})
		It("unsuccessfully creates a hop interactively due to invalid method", func() {
			ui.SetMaxInvalidInput(1)
			bIn.Write([]byte("Test Hop\n1.25\n12\n35\nPellet\n"))
			hop, err := CreateHopInteractively(ui)
			Ω(bErr).To(gbytes.Say("Error reading input."))
			Ω(hop).Should(BeNil())
			Ω(err).ShouldNot(Succeed())
		})
	})
})
