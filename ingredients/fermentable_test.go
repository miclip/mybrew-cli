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

var _ = Describe("Fermentable", func() {
	Context("Fermentable Calculations", func() {
		var f Fermentable
		BeforeEach(func() {
			f = Fermentable{
				Name:      "2 Row",
				Amount:    23.35,
				Potential: 1.036,
				Yield:     77.9,
				Lovibond:  2,
			}
		})
		It("Calculates the Points()", func() {
			Ω(f.Points()).Should(Equal(35.834))
		})
		It("Calculates the PointsByAmount()", func() {
			Ω(f.PointsByAmount()).Should(Equal(836.7239000000001))
		})
		It("Calculates the ColorMCU()", func() {
			Ω(f.ColorMCU()).Should(Equal(46.7))
		})
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
		It("Prints a fermentable", func() {
			f := Fermentable{
				Name:      "2 Row",
				Amount:    23.35,
				Potential: 1.036,
				Yield:     77.9,
				Lovibond:  2,
				Type:      "Grain",
			}
			f.Print(ui)
			Ω(bOut).To(gbytes.Say("2 Row Amount: 23.4 Yield: 77.9 Potential: 1.036 Lovibond: 2 Type: Grain"))
		})
		It("successfully creates a fermentable interactively", func() {
			ui.SetMaxInvalidInput(1)
			bIn.Write([]byte("Test Fermentable\n12.0\n1.036\n77.9\n2\nGrain\n"))
			fermentable, err := CreateFermentableInteractively(ui)
			Ω(bErr).ToNot(gbytes.Say("Error reading input."))
			Ω(fermentable).ShouldNot(BeNil())
			Ω(err).Should(Succeed())
			Ω(fermentable.Name).Should(Equal("Test Fermentable"))
			Ω(fermentable.Amount).Should(Equal(12.0))
			Ω(fermentable.Potential).Should(Equal(1.036))
			Ω(fermentable.Yield).Should(Equal(77.9))
			Ω(fermentable.Lovibond).Should(Equal(2.0))
			Ω(fermentable.Type).Should(Equal("Grain"))
		})
		It("unsuccessfully creates a fermentable interactively due to invalid name", func() {
			ui.SetMaxInvalidInput(1)
			bIn.Write([]byte(nil))
			fermentable, err := CreateFermentableInteractively(ui)
			Ω(bErr).To(gbytes.Say("Error reading input."))
			Ω(fermentable).Should(BeNil())
			Ω(err).ShouldNot(Succeed())
		})
		It("unsuccessfully creates a fermentable interactively due to invalid amount", func() {
			ui.SetMaxInvalidInput(1)
			bIn.Write([]byte("Test Fermentable\nxx\n1.036\n77.9\n2\nGrain\n"))
			fermentable, err := CreateFermentableInteractively(ui)
			Ω(bErr).To(gbytes.Say("Invalid float, please enter a valid value."))
			Ω(fermentable).Should(BeNil())
			Ω(err).ShouldNot(Succeed())
		})
		It("unsuccessfully creates a fermentable interactively due to invalid potential", func() {
			ui.SetMaxInvalidInput(1)
			bIn.Write([]byte("Test Fermentable\n12.0\nxx\n77.9\n2\nGrain\n"))
			fermentable, err := CreateFermentableInteractively(ui)
			Ω(bErr).To(gbytes.Say("Invalid float, please enter a valid value."))
			Ω(fermentable).Should(BeNil())
			Ω(err).ShouldNot(Succeed())
		})
		It("unsuccessfully creates a fermentable interactively due to invalid yield", func() {
			ui.SetMaxInvalidInput(1)
			bIn.Write([]byte("Test Fermentable\n12.0\n1.036\nxx\n2\nGrain\n"))
			fermentable, err := CreateFermentableInteractively(ui)
			Ω(bErr).To(gbytes.Say("Invalid float, please enter a valid value."))
			Ω(fermentable).Should(BeNil())
			Ω(err).ShouldNot(Succeed())
		})
		It("unsuccessfully creates a fermentable interactively due to invalid lovibond", func() {
			ui.SetMaxInvalidInput(1)
			bIn.Write([]byte("Test Fermentable\n12.0\n1.036\n77.9\nxx\nGrain\n"))
			fermentable, err := CreateFermentableInteractively(ui)
			Ω(bErr).To(gbytes.Say("Invalid float, please enter a valid value."))
			Ω(fermentable).Should(BeNil())
			Ω(err).ShouldNot(Succeed())
		})
		It("unsuccessfully creates a fermentable interactively due to invalid type", func() {
			ui.SetMaxInvalidInput(1)
			bIn.Write([]byte("Test Fermentable\n12.0\n1.036\n77.9\n2\n"))
			fermentable, err := CreateFermentableInteractively(ui)
			Ω(bErr).To(gbytes.Say("Error reading input."))
			Ω(fermentable).Should(BeNil())
			Ω(err).ShouldNot(Succeed())
		})
	})
})
