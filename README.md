# BranchDecoder

![GitHub](https://img.shields.io/github/license/jmizell/BranchDecoder.svg?color=00ff00)
[![GoDoc](https://godoc.org/github.com/jmizell/BranchDecoder/decoder?status.svg)](https://godoc.org/github.com/jmizell/BranchDecoder/decoder)

BranchDecoder is a recursive, broadly searching, string encoding detection, 
and decoding library. It attempts to walk through every potential string 
encoding and assembles an EncodingTree of matched encodings and optionally 
their intermediate values.

**This is in early active development, and not a complete project.**

## Docs
https://godoc.org/github.com/jmizell/BranchDecoder/decoder

## Development Requirements

* [go version >= 1.12.4](https://golang.org/dl/)
* [ragel version == 6.9](http://www.colm.net/open-source/ragel/)
