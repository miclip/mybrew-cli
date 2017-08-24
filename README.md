[![Build Status](https://travis-ci.org/miclip/mybrewgo.svg?branch=master)](https://travis-ci.org/miclip/mybrewgo)
[![codecov](https://codecov.io/gh/miclip/mybrewgo/branch/master/graph/badge.svg)](https://codecov.io/gh/miclip/mybrewgo)

# MyBrewGo

MyBrewGo is a very fast command line interface for managing homebrew recipes. MyBrewGo
supports recipes in either YAML, JSON, XML and they can be added directly via the cli.

Recipes are stored local to the executable in the human readable YAML format. This enables
the user to choose a source code repository like github.com to store and backup recipes.

### Developer Getting Started

```sh
mkdir -p $(go env GOPATH)/src/github.com/miclip
cd $(go env GOPATH)/src/github.com/miclip
git clone git@github.com:miclip/mybrewgo
cd mybrewgo
```
At this point you should be able to run the unit tests:

```sh
go test $(go list ./... | grep -v /vendor/)
```
### Usage

#### Add Recipe
Add a recipe and display basic recipe calculations:

```sh
mybrewgo recipes add --path ./test_data/accidental-ipa.yml

Recipe Add...

Reading recipe file .../mybrewgo/test_data/accidental-ipa.yml

Recipe: Accidental IPA
Style: American IPA
Batch Size: 11 Boil Time: 90
OG: 1.07 FG: 1.016 IBU: 37.8 ABV: 7.1 SRM: 9.4
Fermentables:
2 Row Amount: 23.4 Yield: 77.9 Potential: 1.036 Lovibond: 2 Type: Grain
Vienna Malt Amount: 1.6 Yield: 77.9 Potential: 1.036 Lovibond: 4 Type: Grain
White Wheat Amount: 1 Yield: 86.7 Potential: 1.04 Lovibond: 2 Type: Grain
Hops:
Galaxy Amount: 1.25 Time: 60 Alpha: 13 Form: Pellet Method: Boil
Centennial Amount: 1 Time: 10 Alpha: 9.9 Form: Pellet Method: Boil
Cascade Amount: 1 Time: 10 Alpha: 6.7 Form: Pellet Method: Boil
Centennial Amount: 1 Time: 0 Alpha: 9.9 Form: Pellet Method: Boil
Cascade Amount: 1 Time: 0 Alpha: 6.7 Form: Pellet Method: Boil
Citra Amount: 1 Time: 12 Alpha: 12 Form: Pellet Method: Dry Hop
Galaxy Amount: 1.25 Time: 12 Alpha: 13 Form: Pellet Method: Dry Hop
Yeasts:
Safale American Attenuation: 77
```

Note: Recipe will be added to a local YAML file name `mybrewgo_recipes.yml`

Add a recipe interactively via the user interface:

```sh
mybrewgo recipes add

Adding Recipe...
Recipe Name:
```
The command line interface will prompt for each property of a recipe and then
the ingredients. It will save the recipe into the local repository and display the
the recipe details and calculated values.

#### List Recipes
List all the recipes in the local repository

```sh
mybrewgo recipes
Recipes:
0. Accidental IPA\0	1. Czech Pilsner\0	2. Dry Irish Stout\0

Select a recipe: 0

Recipe: Accidental IPA Version: 0
Style: American IPA
Batch Size: 11 Boil Time: 90
OG: 1.07 FG: 1.016 IBU: 37.8 ABV: 7.1 SRM: 9.4
Fermentables:
2 Row Amount: 23.4 Yield: 77.9 Potential: 1.036 Lovibond: 2 Type: Grain
Vienna Malt Amount: 1.6 Yield: 77.9 Potential: 1.036 Lovibond: 4 Type: Grain
White Wheat Amount: 1 Yield: 86.7 Potential: 1.04 Lovibond: 2 Type: Grain
Hops:
Galaxy Amount: 1.25 Time: 60 Alpha: 13 Form: Pellet Method: Boil
Centennial Amount: 1 Time: 10 Alpha: 9.9 Form: Pellet Method: Boil
Cascade Amount: 1 Time: 10 Alpha: 6.7 Form: Pellet Method: Boil
Centennial Amount: 1 Time: 0 Alpha: 9.9 Form: Pellet Method: Boil
Cascade Amount: 1 Time: 0 Alpha: 6.7 Form: Pellet Method: Boil
Citra Amount: 1 Time: 12 Alpha: 12 Form: Pellet Method: Dry Hop
Galaxy Amount: 1.25 Time: 12 Alpha: 13 Form: Pellet Method: Dry Hop
Yeasts:
Safale American Attenuation: 77
```

#### View Recipe
##### By Name and Version:

```sh
mybrewgo recipe -n 'Accidental IPA' -v 1
...
Recipe: Accidental IPA Version: 1
Style: American IPA
Batch Size: 11 Boil Time: 90
OG: 1.07 FG: 1.016 IBU: 37.8 ABV: 7.1 SRM: 9.4
Fermentables:
2 Row Amount: 23.4 Yield: 77.9 Potential: 1.036 Lovibond: 2 Type: Grain
Vienna Malt Amount: 1.6 Yield: 77.9 Potential: 1.036 Lovibond: 4 Type: Grain
White Wheat Amount: 1 Yield: 86.7 Potential: 1.04 Lovibond: 2 Type: Grain
Hops:
Galaxy Amount: 1.25 Time: 60 Alpha: 13 Form: Pellet Method: Boil
Centennial Amount: 1 Time: 10 Alpha: 9.9 Form: Pellet Method: Boil
Cascade Amount: 1 Time: 10 Alpha: 6.7 Form: Pellet Method: Boil
Centennial Amount: 1 Time: 0 Alpha: 9.9 Form: Pellet Method: Boil
Cascade Amount: 1 Time: 0 Alpha: 6.7 Form: Pellet Method: Boil
Citra Amount: 1 Time: 12 Alpha: 12 Form: Pellet Method: Dry Hop
Galaxy Amount: 1.25 Time: 12 Alpha: 13 Form: Pellet Method: Dry Hop
Yeasts:
Safale American Attenuation: 77
```

##### Search By Name:

```sh
mybrewgo recipes search 'c'

Search results for 'c':
0. Accidental IPA
1. Czech Pilsner
Please select a result:
1
...
Recipe: Czech Pilsner Version: 0
Style: Bohemian Pilsner
Batch Size: 10 Boil Time: 60
OG: 1.049 FG: 1.014 IBU: 54.5 ABV: 4.6 SRM: 9.3
Fermentables:
2 Row Amount: 19 Yield: 77.9 Potential: 1.036 Lovibond: 2 Type: Grain
Crystal 10 Amount: 0.5 Yield: 73.6 Potential: 73.6 Lovibond: 10 Type: Grain
Hops:
Perle Amount: 2 Time: 60 Alpha: 8 Form: Pellet Method: Boil
Saaz Amount: 2 Time: 30 Alpha: 4 Form: Pellet Method: Boil
Saaz Amount: 2 Time: 15 Alpha: 4 Form: Pellet Method: Boil
Yeasts:
Pilsner Lager Yeast Attenuation: 72
```

### Tasks/Features

- [x] Add Recipe from yaml
- [x] Store in local repo (yaml)
- [x] Basic Recipe calculations
- [x] Display basic recipe details
- [x] Find recipes
- [x] List recipes in local store
- [ ] Add recipe via cli
- [ ] Modify recipes via cli
- [ ] Mashing calculations
- [ ] Scaling recipes
- [ ] Pulling ingredients from web api
- [ ] Storing recipes in database via web api
- [ ] Brew day instructions and calculations
- [ ] Customize hop Utilizations
