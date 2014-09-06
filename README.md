# pxgen

A simple, strong password generator based on: [password-generator-go](https://github.com/cmiceli/password-generator-go).

Generates 20 character passwords by default.

## Getting Started

### Usage

```bash
pxgen --help
usage: pxgen [length]

$ pxgen
-ce5Ef3@}4,AgIo,p=U

$ pxgen 10
T8%R90[{F:
```

### Getting pxgen

You can get pxgen using the standard go workflow - ensure `GOPATH` is set appropriately, and then:

```sh
go get github.com/robmorgan/pxgen
```

which will generate a binary at `$GOPATH/bin/pxgen`
