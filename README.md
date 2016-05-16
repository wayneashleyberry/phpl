## phpl

> Inspired by [phplint](https://github.com/wayneashleyberry/phplint), phpl is a concurrent php linter written in go.

[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/phpl)](https://goreportcard.com/report/github.com/wayneashleyberry/phpl)

#### Installation

You can download the [latest binary](https://github.com/wayneashleyberry/phpl/releases/latest), or if you have a go workspace:

```sh
go get github.com/wayneashleyberry/phpl
```

#### Usage

```
cd /path/to/php
phpl
```

#### Motivation

PHP's native linter can only lint one file at a time, and that's slow.

#### License

MIT © [Wayne Ashley Berry](http://www.wayneashleyberry.com)
