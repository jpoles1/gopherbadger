# GopherBadger
### Generate coverage badge images using Go!

<img src="https://raw.githubusercontent.com/jpoles1/gopherbadger/master/badge.png" width="100"/>

One day, I noticed that there was no easy way to generate coverage badges for my Golang projects. So I made one (see above)! 

This program can be placed in any project which contains a makefile with a "cover" command. This should "cover" command should generate an output containg a line like: "coverage: 84%". From this command, the software will generate a badge image in the same folder called "badge.png".

A makefile for Go might look like the following:

```
build:
	go build
run:
	make build && ./gocoverbadger
cover:
	go test -coverprofile=coverage.out && go tool cover -html=coverage.out -o=coverage.html
```
