# HL7

![Test](https://github.com/zamariola/hl7reader/workflows/Test/badge.svg)

This is a basic HL7 parser written in Go. There are no external dependencies,
only the standard library is used (except for tests, where
`github.com/testify/assert` is used).

This parser accepts an `io.Reader` as the input, so anything that follows that
interface should be usable here.

This library is tested to work on the following platforms:

- Go versions 1.15.x and 1.14.x
- Latest version of macOS
- Latest version of Windows
- Latest version of Ubuntu

It should probably work without issue in older versions of Go and operating
systems since we are not utilizing any low-level features or anything
particularly modern in Go.

## Installation

Installation is simple:

```bash
$ go get github.com/zamariola/hl7reader
```

## Usage

Usage is meant to mimic the semantics of a reader. Additionally, the HL7 data
is read from an `io.Reader`, so it's relatively trivial to read HL7 data from
a variety of sources (a file, TCP connection, `bytes.Buffer`, etc.). Actual
parsing of the message happens when you fetch a segment. This laziness should
help with large messages, particularly when you find out halfway through the
message that you don't care about it anymore (maybe it doesn't have anything
to do with your business case).

For example usage, see the [example program](example/main.go). There is an
example file with HL7 data in it gathered from
[Wikipedia](https://en.wikipedia.org/wiki/Health_Level_7) and
[Ringholm bv](http://www.ringholm.com/docs/04300_en.htm).

Further usage information can be found
[here](https://pkg.go.dev/github.com/zamariola/hl7).
