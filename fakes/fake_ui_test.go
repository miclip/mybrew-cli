package fakes_test

import (
	"time"

	"github.com/miclip/mybrewgo/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("FakeUi", func() {
	Context("Fake UI", func() {
		var (
			bOut *gbytes.Buffer
			bErr *gbytes.Buffer
		)
		BeforeEach(func() {
			bOut, bErr = gbytes.NewBuffer(), gbytes.NewBuffer()
			_, _ = gbytes.TimeoutWriter(bOut, time.Second), gbytes.TimeoutWriter(bOut, time.Second)
		})
		It("Creates a fake ui", func() {
			ui := fakes.NewFakeUI(bOut, bErr)
			Ω(ui).ShouldNot(BeNil())
		})
	})

})
