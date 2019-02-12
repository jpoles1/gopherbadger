# GopherBadger
### Generate coverage badge images using Go!

<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-68%25-brightgreen.svg?longCache=true&style=flat)</a>

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

Either enter a Markdown file that does not already exist, or a Markdown file (like your README.md) that contains the following tag somewhere in the contents:

!`[gopherbadger-tag-do-not-edit]()`

This tag will be replaced by the image for your coverage badge. 

<hr>

Manually set the coverage value (note: do not include %):

`gopherbadger -md="README.md,coverage.md" -manualcov=95`

<hr>

Test code coverage using build tags:

`gopherbadger -tags "unit"`

<hr>

## Confused?

Try running:

```
gopherbadger -h
```
