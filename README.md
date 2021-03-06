# readfq

[![Documentation](https://godoc.org/github.com/sstadick/readfq?status.svg)](http://godoc.org/github.com/sstadick/readfq)
[![Go Report Card](https://goreportcard.com/badge/github.com/sstadick/readfq)](https://goreportcard.com/report/github.com/sstadick/readfq)
[![Maintainability](https://api.codeclimate.com/v1/badges/15d5ff1cb3d07eca8b4f/maintainability)](https://codeclimate.com/github/sstadick/readfq/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/15d5ff1cb3d07eca8b4f/test_coverage)](https://codeclimate.com/github/sstadick/readfq/test_coverage)

A packaged version of readfq implementation for reading fastq and fastq formatted files.

## Synopsis

readfq is a fast and light wieght package for processing FASTA and FASTQ formatted files. This package aims to make it easy to include for your project as a dependancy, or just copy and paste it in. Tests are being added for peace of mind.

## Examples

Please see the [docs](https://godoc.org/github.com/sstadick/readfq)

## Prior Art

This Go version is a port of [drio's][1] version, which can be found [here][2]. He also did a nice write up of
it [here][3]. Additionally, other implementations, as well as the original c implementation can be found
[here][4]. A good discussion of related parsers can be found on biostars [here][5]

## Contributing

Yes! Please do, this includes bug reports, things you tought about, or PR's.

[1]: https://github.com/drio
[2]: https://github.com/drio/drio.go/blob/master/bio/fasta/fasta.go
[3]: http://drio.github.io/2012/10/golang-performance/
[4]: https://github.com/lh3/readfq
[5]: https://www.biostars.org/p/10353/
