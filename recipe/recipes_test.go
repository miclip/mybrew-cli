package recipe_test

import (
	. "github.com/miclip/mybrewgo/recipe"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Recipes", func() {
	It("Creates a new instance of Recipes", func() {
		recipes := NewRecipes()
		Ω(recipes).ShouldNot(BeNil())
		Ω(len(recipes.Recipes)).Should(Equal(0))
	})
	It("Creates a new instance of Recipes and loads recipes", func() {
		recipes := NewRecipes()
		fileName := "../test_data/accidental-ipa.yml"
		recipe, err := OpenRecipe(fileName)
		Ω(err).Should(Succeed())
		recipes.AddRecipe(recipe)
	})
	Context("Recipes loaded in local repository", func() {
		var recipes Recipes
		BeforeEach(func() {
			recipes = NewRecipes()
			fileName := "../test_data/accidental-ipa.yml"
			recipe, _ := OpenRecipe(fileName)
			Ω(recipe).ShouldNot(BeNil())
			recipes.AddRecipe(recipe)
		})
		It("Finds a recipe using the key", func() {
			result := recipes.FindByKey("Accidental IPA", 0)
			Ω(result).ShouldNot(BeNil())
			Ω(result.Name).Should(Equal("Accidental IPA"))
		})
		It("Search a recipe using part of the name", func() {
			matches := recipes.SearchByName("acci")
			Ω(matches).ShouldNot(BeNil())
			Ω(len(matches)).Should(Equal(1))
		})
	})
})
