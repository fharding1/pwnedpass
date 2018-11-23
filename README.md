# PwnedPass

PwnedPass is a Go library for accessing the [Pwned Password
API](https://haveibeenpwned.com/API/v2). It currently only supports the V2
version of the API. It's got a dead simple API, tests, and benchmarks.

## Usage

See [godoc](https://godoc.org/github.com/fharding1/pwnedpass) for usage and
documentation.

## Performance

Running on my computer, the benchmark results I get are:

    BenchmarkCount-8                    5000            231893 ns/op
    BenchmarkCountIntegration-8          100          24331000 ns/op

The test with a mock local HTTP server takes about 231893 ns/op, while the
integration tests against the actual API take 24331000 ns/op, meaning the only
real limitation is the speed of the API and your network.
