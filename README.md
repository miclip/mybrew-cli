# MyBrewGo

Command Line Interface to manage Homebrew beer recipes and calculations

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
mybrewgo recipes add ./test_data/accidental-ipa.yml

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
Safale American Attenution: 77
```

Note: Recipe will be added to a local YAML file name `mybrewgo_recipes.yml`


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
Safale American Attenution: 77```
