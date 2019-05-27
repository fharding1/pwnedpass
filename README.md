# PwnedPass

[![Godoc](https://godoc.org/github.com/fharding1/pwnedpass?status.svg)](http://godoc.org/github.com/fharding1/pwnedpass)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![CI: CircleCI](https://circleci.com/gh/fharding1/pwnedpass.svg?style=svg)](https://circleci.com/gh/fharding1/pwnedpass)

PwnedPass is a Go library for accessing the [Pwned Password
API](https://haveibeenpwned.com/API/v2). It currently only supports the V2
version of the API. It's got a dead simple API, tests, and benchmarks.

## Usage

See [godoc](https://godoc.org/github.com/fharding1/pwnedpass) for usage and
documentation.

## CLI

To install the CLI tool from the latest source run:

```
go install github.com/fharding1/pwnedpass/cmd/pwnedpass
```

Otherwise, you can visit the [releases](https://github.com/fharding1/pwnedpass/releases)
to download the binary of a release to use.

### Usage

Using the CLI tool is fairly simple:

```
pwnedpass password
pwnedpass "a password with spaces"
```

## Performance

Running on my computer, the benchmark results I get are:

```
BenchmarkCount-8              	    5000	    216640 ns/op
BenchmarkCountIntegration-8   	      50	  26431351 ns/op
```

The test with a mock local HTTP server takes about 231893 ns/op, while the
integration tests against the actual API take 24331000 ns/op, meaning the only
real limitation is the speed of the API and your network.
