package ingredients_test

import (
	. "github.com/miclip/mybrewgo/ingredients"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Fermentable", func() {
	It("Prints the fermentable details", func() {
		f := Fermentable{
			Name:      "2 Row",
			Amount:    23.35,
			Potential: 1.036,
			Yield:     77.9,
			Lovibond:  2,
		}
		Ω(f).ShouldNot(BeNil())

	})
})
