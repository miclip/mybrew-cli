package mybrewgo

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hoputilization", func() {
	It("Finds the UpperHopMinutes", func() {
		h := NewHopUtilizations()
		Ω(h.findUpperMinutes(41)).Should(Equal(45))
	})
	It("Finds the LowerHopMinutes", func() {
		h := NewHopUtilizations()
		Ω(h.findLowerMinutes(41)).Should(Equal(30))
	})
	It("Finds Hop Utilizations 41 min addition for 1.07 OG", func() {
		h := NewHopUtilizations()
		u, err := h.FindHopUtilization(41, 1.070, "Pellet")
		Ω(err).Should(Succeed())
		Ω(u).ShouldNot(BeNil())
		Ω(u.Minutes).Should(Equal(41))
		Ω(u.Gravity).Should(Equal(1.07))
		Ω(u.Percentage).Should(Equal(17))
	})
	It("Finds Hop Utilizations 91 min addition for 1.07 OG", func() {
		h := NewHopUtilizations()
		u, err := h.FindHopUtilization(91, 1.07, "Pellet")
		Ω(err).Should(Succeed())
		Ω(u).ShouldNot(BeNil())
		Ω(u.Minutes).Should(Equal(91))
		Ω(u.Gravity).Should(Equal(1.07))
		Ω(u.Percentage).Should(Equal(25))
	})
	It("Finds Hop Utilizations 1 min addition for 1.07 OG", func() {
		h := NewHopUtilizations()
		u, err := h.FindHopUtilization(1, 1.07, "Pellet")
		Ω(err).Should(Succeed())
		Ω(u).ShouldNot(BeNil())
		Ω(u.Minutes).Should(Equal(1))
		Ω(u.Gravity).Should(Equal(1.07))
		Ω(u.Percentage).Should(Equal(0))
	})

})
