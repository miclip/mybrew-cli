package mybrewgo

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hoputilization", func() {
	It("Finds the UpperHopMinutes", func() {
		h := NewHopUtilizations()
		Ω(h.findHopUpperMinutes(41)).Should(Equal(45))
	})
	It("Finds the LowerHopMinutes", func() {
		h := NewHopUtilizations()
		Ω(h.findHopLowerMinutes(41)).Should(Equal(30))
	})
	It("Finds Hop Utilizations 41 min addition for 1.07 OG", func() {
		h := NewHopUtilizations()
		l, u, err := h.FindByAdditionTime(41, 1.07, "Pellet")
		Ω(err).Should(Succeed())
		Ω(l).ShouldNot(BeNil())
		Ω(l.Minutes).Should(Equal(30))
		Ω(l.Gravity).Should(Equal(1.07))
		Ω(u).ShouldNot(BeNil())
		Ω(u.Minutes).Should(Equal(45))
		Ω(u.Gravity).Should(Equal(1.07))
	})
	It("Finds Hop Utilizations 91 min addition for 1.07 OG", func() {
		h := NewHopUtilizations()
		l, u, err := h.FindByAdditionTime(91, 1.07, "Pellet")
		Ω(err).Should(Succeed())
		Ω(l).ShouldNot(BeNil())
		Ω(l.Minutes).Should(Equal(90))
		Ω(l.Gravity).Should(Equal(1.07))
		Ω(u).ShouldNot(BeNil())
		Ω(u.Minutes).Should(Equal(90))
		Ω(u.Gravity).Should(Equal(1.07))
	})
	It("Finds Hop Utilizations 1 min addition for 1.07 OG", func() {
		h := NewHopUtilizations()
		l, u, err := h.FindByAdditionTime(1, 1.07, "Pellet")
		Ω(err).Should(Succeed())
		Ω(l).ShouldNot(BeNil())
		Ω(l.Minutes).Should(Equal(5))
		Ω(l.Gravity).Should(Equal(1.07))
		Ω(u).ShouldNot(BeNil())
		Ω(u.Minutes).Should(Equal(5))
		Ω(u.Gravity).Should(Equal(1.07))
	})

})
