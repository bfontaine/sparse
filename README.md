# sparse

**sparse** is a Go library for memory-efficient and thread-safe sparse data
structures.

[![Build Status](https://travis-ci.org/bfontaine/sparse.svg?branch=master)](https://travis-ci.org/bfontaine/sparse)
[![GoDoc](https://godoc.org/github.com/bfontaine/sparse?status.svg)](https://godoc.org/github.com/bfontaine/sparse)

## Install

    go get github.com/bfontaine/sparse

## Content

Only two data structures are provided for now:

* `BoolSlice` is a compact representation of `[]bool`
* `BoolMatrix` is based on `BoolSlice` to provide a compact representation of a
  booleans matrix

These data structures are more efficient in memory but less efficient in time.
Read operations on both are in `O(n)` even if in practice they’re much lower if
the structures contain sparse data. The write operations are costly due to the
internal data representation. You shouldn’t use these if you want small slices
with frequent write access. They become practical with large data structures
consisting of *sparse* data.

For example, consider a slice of 1M `false` values with only 1000 `true`
values in it. A normal `bool` slice will need 1M cells while a `BoolSlice` will
use between 2 and 2001 cells depending on where the `true` values are. Append
another 1M false values: the `bool` slice will double its size while the
`BoolSlice` internal size won’t change.

## Internal representation

A `BoolSlice` has an internal `int64` slice `m` representing the contiguous
sequences of boolean values. The first sequence is always `false` by
convention. Each cell of `m` contains the length of a contiguous sequence.

Examples:

| content              | `m`             |
|----------------------|-----------------|
| `[]`                 | `[]`            |
| `[F]`                | `[1]`           |
| `[F, F, F]`          | `[3]`           |
| `[F, F, T]`          | `[2, 1]`        |
| `[T, T]`             | `[0, 2]`        |
| `[F, F, T, F, F, F]` | `[2, 1, 3]`     |
| `[T, 20 × F, T]`     | `[0, 1, 20, 1]` |


## Limitations

* The `Set` method isn’t supported at the moment. This means the only way to
  write on a `BoolSlice` after its initialization is with `.Append`.
  I’m looking for an efficient *and* elegant way to solve this issue.
