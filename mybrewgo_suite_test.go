package mybrewgo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMybrewgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mybrewgo Suite")
}
