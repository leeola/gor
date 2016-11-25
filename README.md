
# GoR(un)

Gor (Gorun) is a simple binary to build and run the given Go package. Just like
`go run`.

## Motivation

`go run` has a limitation in that it does not replicate the binary. It doesn't pass
args and it doesn't pass the exit number. `gor` acts *(as best it can)* like the
given binary.

This basically just runs `go build <pkg> && ./<pkg> ...`, but with less typing and
repetition. Yes, a shell script could achieve the same thing, but why would i want
to write it in Bash? I like Go :)

## Installation

```
go install github.com/leeola/gor
```

## Usage

`gor <go package>`

Example:

```
~/g/s/g/l/gor> gor _examples/hello.go
hello from hello.go
~/g/s/g/l/gor> gor github.com/leeola/gor/_examples/hello-package
hello from hello-package/hello-package.go
```

Note that the binary will be named with the `.gor` extension and is not cleaned up
after running. This allows you to call it repeatedly without recompiling.

## Future Development

Additional features are welcome and will likely be added. Including merger with the
[goscriptify](https://github.com/leeola/goscriptify) project. Until then, this project
is stupidly simple.

## LICENSE

MIT
