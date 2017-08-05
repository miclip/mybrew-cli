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
<<<<<<< Updated upstream
=======

### Usage

Add recipe and display basic recipe calculations:

```sh
mybrewgo recipe add ./test_data/accidental-ipa.yml

Recipe Add...

Reading recipe file .../mybrewgo/test_data/accidental-ipa.yml

Recipe: Accidental IPA
OG: 1.07 FG: 1.016 IBU: 37.8 ABV: 7.1 SRM: 9.4
```

Recipe will be added to a local YAML file name `mybrewgo_recipes.yml`
>>>>>>> Stashed changes
