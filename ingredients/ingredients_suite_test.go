package ingredients_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestIngredients(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ingredients Suite")
}
