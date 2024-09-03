# just

[![Go Reference](https://pkg.go.dev/badge/github.com/kazhuravlev/just.svg)](https://pkg.go.dev/github.com/kazhuravlev/just)
[![License](https://img.shields.io/github/license/kazhuravlev/just?color=blue)](https://github.com/kazhuravlev/just/blob/master/LICENSE)
[![Build Status](https://github.com/kazhuravlev/just/actions/workflows/tests.yml/badge.svg?branch=master)](https://github.com/kazhuravlev/just/actions/workflows/tests.yml?query=branch%3Amaster)
[![Go Report Card](https://goreportcard.com/badge/github.com/kazhuravlev/just)](https://goreportcard.com/report/github.com/kazhuravlev/just)
[![CodeCov](https://codecov.io/gh/kazhuravlev/just/branch/master/graph/badge.svg?token=tNKcOjlxLo)](https://codecov.io/gh/kazhuravlev/just)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#utilities)

This project contains features that help make the noisy stuff in every project.

- Filter slices, maps.
- Applying functions to collections.
- Null[any] for optional fields in API and sql.DB queries.

## Most helpful functions

- [`SliceIter`](https://pkg.go.dev/github.com/kazhuravlev/just#SliceIter) allows you to iterate over slice with
  special `IterContext` that provide some methods like `IsFirst`, `IsLast`
- [`SliceMap`](https://pkg.go.dev/github.com/kazhuravlev/just#SliceMap) and
  [`SliceMapErr`](https://pkg.go.dev/github.com/kazhuravlev/just#SliceMapErr) allow you to map a slice to
  another one. Useful for adapters.
- [`NewPool`](https://pkg.go.dev/github.com/kazhuravlev/just#NewPool) `sync.Pool` with generics
- [`ChanAdapt`](https://pkg.go.dev/github.com/kazhuravlev/just#ChanAdapt) allows to create an adapted version of channel
- [`ContextWithTimeout`](https://pkg.go.dev/github.com/kazhuravlev/just#ContextWithTimeout) runs a function with context
  and timeout
- [`ErrAs`](https://pkg.go.dev/github.com/kazhuravlev/just#ErrAs) helps to handle an errors
- [`SliceChunk`](https://pkg.go.dev/github.com/kazhuravlev/just#SliceChunk) iterates over chunks
- [`SliceChain`](https://pkg.go.dev/github.com/kazhuravlev/just#SliceChain) joins slices to one
- [`StrSplitByChars`](https://pkg.go.dev/github.com/kazhuravlev/just#StrSplitByChars) splits the string by symbols

## Examples

This library contains a bunch of functions. Please
see [pkg.go.dev](https://pkg.go.dev/github.com/kazhuravlev/just#pkg-examples)
for examples.
