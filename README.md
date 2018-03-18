# GopherBadger
### Generate coverage badge images using Go!

<img src="https://ci.jpoles1.com/api/badges/jpoles1/gopherbadger/status.svg" height="25"/>     <img src="https://raw.githubusercontent.com/jpoles1/gopherbadger/master/coverage_badge.png" height="25"/>

One day, I noticed that there was no easy way to generate coverage badges for my Golang projects. So I made one (see above)! 

## Getting Started 

To install the executeable (ensure your $PATH contains $GOPATH/bin):

```
go get github.com/jpoles1/gopherbadger
```

This program can be run in any project which contains a makefile with a "cover" command. This "cover" command should generate an output containg a line like: "coverage: 84%". From this command, the software will generate a badge image in the same folder called "coverage_badge.png". Simple navigate to the project root and run "gopherbadger"!

## Go Example

A makefile for a Go project might look like the following:

```
build:
	go build
run:
	make build && ./gopherbadger
cover:
	go test -coverprofile=coverage.out && go tool cover -html=coverage.out -o=coverage.html
```

Running "gopherbadger" in a project using this makefile will generate both a badge image "coverage_badge.png" and a coverage report "coverage.html". These can be handy in improving test coverage on your project!
