package ui_test

import (
	"time"

	"github.com/miclip/mybrewgo/fakes"
	"github.com/miclip/mybrewgo/ui"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Askfor", func() {
	Context("when using a console UI", func() {
		var (
			bOut *gbytes.Buffer
			bErr *gbytes.Buffer
			bIn  *gbytes.Buffer
			ui   ui.UI
		)
		BeforeEach(func() {
			bOut, bErr, bIn = gbytes.NewBuffer(), gbytes.NewBuffer(), gbytes.NewBuffer()
			_, _ = gbytes.TimeoutWriter(bOut, time.Second), gbytes.TimeoutWriter(bOut, time.Second)
			_ = gbytes.TimeoutReader(bIn, time.Second)
			ui = fakes.NewFakeUI(bOut, bErr, bIn)
		})
		It("AskForText should return the text value", func() {
			bIn.Write([]byte("testing\n"))
			s, err := ui.AskForText("Please type test:")
			Ω(err).Should(Succeed())
			Ω(bOut).To(gbytes.Say("Please type test:"))
			Ω(s).Should(Equal("testing"))
		})
		It("two calls to AskForText should return the text values", func() {
			bIn.Write([]byte("testing1\ntesting2\n"))
			s, err := ui.AskForText("Please type test 1:")
			Ω(err).Should(Succeed())
			Ω(bOut).To(gbytes.Say("Please type test 1:"))
			Ω(s).Should(Equal("testing1"))
			s, err = ui.AskForText("Please type test 2:")
			Ω(err).Should(Succeed())
			Ω(bOut).To(gbytes.Say("Please type test 2:"))
			Ω(s).Should(Equal("testing2"))
		})
		It("AskForInt should return the int value", func() {
			bIn.Write([]byte("12\n"))
			i, err := ui.AskForInt("Please type an int:")
			Ω(err).Should(Succeed())
			Ω(bOut).To(gbytes.Say("Please type an int:"))
			Ω(i).Should(Equal(12))
		})
		It("AskForInt should error if the value is not an int", func() {
			ui.SetMaxInvalidInput(1)
			bIn.Write([]byte("a\n"))
			i, err := ui.AskForInt("Please type an int:")
			Ω(bOut).To(gbytes.Say("Please type an int:"))
			Ω(err).ShouldNot(Succeed())
			Ω(i).Should(Equal(0))
			Ω(bErr).Should(gbytes.Say("Invalid int, please enter a valid value."))
		})
		It("AskForInt should error if the ReadString fails", func() {
			ui.SetMaxInvalidInput(1)
			bIn.Write(nil)
			i, err := ui.AskForInt("Please type an int:")
			Ω(bOut).To(gbytes.Say("Please type an int:"))
			Ω(err).ShouldNot(Succeed())
			Ω(i).Should(Equal(0))
			Ω(bErr).Should(gbytes.Say("Error reading input."))
		})
		It("AskForFloat should return the float value", func() {
			bIn.Write([]byte("120.0\n"))
			f, err := ui.AskForFloat("Please type a float:")
			Ω(err).Should(Succeed())
			Ω(bOut).To(gbytes.Say("Please type a float:"))
			Ω(f).Should(Equal(120.0))
		})
		It("AskForInt should error if the value is not an int", func() {
			ui.SetMaxInvalidInput(1)
			bIn.Write([]byte("a\n"))
			i, err := ui.AskForFloat("Please type a float:")
			Ω(bOut).To(gbytes.Say("Please type a float:"))
			Ω(err).ShouldNot(Succeed())
			Ω(i).Should(Equal(0.0))
			Ω(bErr).Should(gbytes.Say("Invalid float, please enter a valid value."))
		})
		It("AskForInt should error if the ReadString fails", func() {
			ui.SetMaxInvalidInput(1)
			bIn.Write(nil)
			i, err := ui.AskForFloat("Please type a float:")
			Ω(bOut).To(gbytes.Say("Please type a float:"))
			Ω(err).ShouldNot(Succeed())
			Ω(i).Should(Equal(0.0))
			Ω(bErr).Should(gbytes.Say("Error reading input."))
		})
	})

})
