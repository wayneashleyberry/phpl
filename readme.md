## phpl

> Inspired by [phplint](https://github.com/wayneashleyberry/phplint), phpl is a concurrent php linter written in [Go](https://golang.org/).

[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/phpl)](https://goreportcard.com/report/github.com/wayneashleyberry/phpl)
[![GoDoc](https://godoc.org/github.com/wayneashleyberry/phpl?status.svg)](https://godoc.org/github.com/wayneashleyberry/phpl)

#### Motivation

PHP's native linter can only lint one file at a time, and that's slow.

#### Installation

The easiest way to install phpl is to download the [latest
binary](https://github.com/wayneashleyberry/phpl/releases/latest) and place it
in your `$PATH`.

Alternatively if you have a [go
workspace](https://golang.org/doc/code.html#Workspaces), you can use `go get`.

```sh
go get github.com/wayneashleyberry/phpl
```

#### Usage

```
cd /path/to/php
phpl
```

#### License

MIT © [Wayne Ashley Berry](https://wayne.cloud)
