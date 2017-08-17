package ui_test

import (
	"time"

	"github.com/miclip/mybrewgo/fakes"
	"github.com/miclip/mybrewgo/ui"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("UI", func() {
	Context("Console UI", func() {
		var (
			bOut *gbytes.Buffer
			bErr *gbytes.Buffer
		)
		BeforeEach(func() {
			bOut, bErr = gbytes.NewBuffer(), gbytes.NewBuffer()
			_, _ = gbytes.TimeoutWriter(bOut, time.Second), gbytes.TimeoutWriter(bOut, time.Second)
		})
		It("Creates a console ui", func() {
			ui := ui.NewConsoleUI()
			Ω(ui).ShouldNot(BeNil())
		})
	})
	Context("Console UI", func() {
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
		It("SystemLinef", func() {
			Ω(ui).ShouldNot(BeNil())
			ui.SystemLinef("test line %s", "arg")
			Ω(bOut).To(gbytes.Say("test line arg"))
			Ω(bOut).To(gbytes.Say("\n"))
		})
		It("Systemf", func() {
			Ω(ui).ShouldNot(BeNil())
			ui.Systemf("test line %s", "arg")
			Ω(bOut).To(gbytes.Say("test line arg"))
			Ω(bOut).ToNot(gbytes.Say("\n"))
		})
		It("PrintLinef", func() {
			Ω(ui).ShouldNot(BeNil())
			ui.PrintLinef("test line %s", "arg")
			Ω(bOut).To(gbytes.Say("test line arg"))
			Ω(bOut).To(gbytes.Say("\n"))
		})
		It("Printf", func() {
			Ω(ui).ShouldNot(BeNil())
			ui.Printf("test line %s", "arg")
			Ω(bOut).To(gbytes.Say("test line arg"))
			Ω(bOut).ToNot(gbytes.Say("\n"))
		})
		It("ErrorLinef", func() {
			Ω(ui).ShouldNot(BeNil())
			ui.ErrorLinef("test line %s", "arg")
			Ω(bErr).To(gbytes.Say("test line arg"))
			Ω(bErr).To(gbytes.Say("\n"))
		})
		It("Errorf", func() {
			Ω(ui).ShouldNot(BeNil())
			ui.Errorf("test line %s", "arg")
			Ω(bErr).To(gbytes.Say("test line arg"))
			Ω(bErr).ToNot(gbytes.Say("\n"))
		})
		It("Displays columns", func() {
			Ω(ui).ShouldNot(BeNil())
			var items = make([]string, 6)
			items[0] = "first"
			items[1] = "second"
			items[2] = "third"
			items[3] = "fourth"
			items[4] = "fifth"
			items[5] = "sixth"
			ui.DisplayColumns(items, 3)
			Ω(bOut).To(gbytes.Say(`0\.\sfirst\t1\.\ssecond\t2\.\sthird\t\s3\.\sfourth\t4\.\sfifth\t5\.\ssixth\t\s`))
		})
	})

})
