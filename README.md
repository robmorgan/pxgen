# pxgen

[![Build Status](https://travis-ci.org/robmorgan/pxgen.svg?branch=master)](https://travis-ci.org/robmorgan/pxgen)
[![Coverage Status](https://coveralls.io/repos/robmorgan/pxgen/badge.png?branch=master)](https://coveralls.io/r/robmorgan/pxgen?branch=master)
[![GoDoc](https://godoc.org/github.com/robmorgan/pxgen?status.png)](https://godoc.org/github.com/robmorgan/pxgen)

A simple, secure password generator based on: [password-generator-go](https://github.com/cmiceli/password-generator-go) and
[uniuri](https://github.com/dchest/uniuri).

Generates 20 character passwords by default.

## Getting Started

### Usage

```bash
pxgen --help
usage: pxgen [length] [count]

$ pxgen
-ce5Ef3@}4,AgIo,p=U

$ pxgen 10
T8%R90[{F:

$ pxgen 30 2
E!MhL&bI)ahW#n+/pebv#GuRbT5fOX
~*}i{_=o+Q`/0jPoDQsrxKV4=AkFI?
```

### Getting pxgen

You can get pxgen using the standard go workflow - ensure `GOPATH` is set appropriately then:

```sh
go get github.com/robmorgan/pxgen
```

which will generate a binary at `$GOPATH/bin/pxgen`

### Testing

To run tests for pxgen, use `make test`.
