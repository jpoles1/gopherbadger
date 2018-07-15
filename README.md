# GopherBadger
### Generate coverage badge images using Go!

![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-73%25-brightgreen.svg?longCache=true&style=flat)

One day, I noticed that there was no easy way to generate coverage badges for my Golang projects. So I made one (see above)! 

## Getting Started 

To install the executeable (ensure your $PATH contains $GOPATH/bin):

```
go get github.com/jpoles1/gopherbadger
```

This program can be run in any Go project that is compatible with the built-in go cover tool, which will extract a coverage percentage based upon files from all of your Go code (even that in sub-packages). Just run gopherbadger in your project root directory.

## Quick Start:

<hr>

To prevent saving of a .png badge:

`gopherbadger -png=false`

<hr>

To update a .md file badge (note: comma-separated):

`gopherbadger -md="README.md,coverage.md"`

<hr>

Manually set the coverage value (note: do not include %):

`gopherbadger -md="README.md,coverage.md" -manualcov=95`

<hr>

## Confused?

Try running:

```
gopherbadger -h
```
