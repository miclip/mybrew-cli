package ingredients_test

import (
	"time"

	"github.com/miclip/mybrew/fakes"
	. "github.com/miclip/mybrew/ingredients"
	"github.com/miclip/mybrew/ui"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Yeast", func() {
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
			y := Yeast{
				Name:        "California Ale",
				Attenuation: 77.9,
			}
			y.Print(ui)
			Ω(bOut).To(gbytes.Say("California Ale Attenuation: 77.9"))
		})
		It("successfully creates a yeast interactively", func() {
			ui.SetMaxInvalidInput(1)
			bIn.Write([]byte("Test Yeast\n77.9\n"))
			yeast, err := CreateYeastInteractively(ui)
			Ω(bErr).ToNot(gbytes.Say("Error reading input."))
			Ω(yeast).ShouldNot(BeNil())
			Ω(err).Should(Succeed())
			Ω(yeast.Name).Should(Equal("Test Yeast"))
			Ω(yeast.Attenuation).Should(Equal(77.9))
		})
		It("unsuccessfully creates a yeast interactively due to invalue attenuation", func() {
			ui.SetMaxInvalidInput(1)
			bIn.Write([]byte(nil))
			yeast, err := CreateYeastInteractively(ui)
			Ω(bErr).To(gbytes.Say("Error reading input."))
			Ω(yeast).Should(BeNil())
			Ω(err).ShouldNot(Succeed())
		})
		It("unsuccessfully creates a yeast interactively due to invalue attenuation", func() {
			ui.SetMaxInvalidInput(1)
			bIn.Write([]byte("Test Yeast\nxx\n"))
			yeast, err := CreateYeastInteractively(ui)
			Ω(bErr).To(gbytes.Say("Invalid float, please enter a valid value."))
			Ω(yeast).Should(BeNil())
			Ω(err).ShouldNot(Succeed())
		})
	})
})
