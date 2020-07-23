package utils_test

import (
	. "github.com/miclip/mybrew/utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Utils", func() {
	It("Rounds", func() {
		Ω(Round(123.555555, .5, 3)).Should(Equal(123.556))
		Ω(Round(123.558, .5, 3)).Should(Equal(123.558))
	})
})
