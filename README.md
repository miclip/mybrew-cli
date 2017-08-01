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
