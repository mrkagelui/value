# value.Of anything!

> Check out [my ptr package](https://github.com/mrkagelui/ptr) to do ptr.Of()!

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mrkagelui/value)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/mrkagelui/value/testing)
[![codecov](https://codecov.io/gh/mrkagelui/value/branch/master/graph/badge.svg?token=A2SCYKTWFZ)](https://codecov.io/gh/mrkagelui/value)
[![Go Report Card](https://goreportcard.com/badge/github.com/mrkagelui/value)](https://goreportcard.com/report/github.com/mrkagelui/value)
![GitHub](https://img.shields.io/github/license/mrkagelui/value)
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/)
![gopher_unboxed](gopher_unboxed.png "gopher_unboxed")

## Why this utility? 

Because I bet you are wary of
```go
var v T
if ptr != nil {
	v = *ptr
}
```
or
```go
var v T
if ptr1 != nil {
	return *ptr1
}
if ptr2 != nil {
	return *ptr2
}
return v
```
With this utility, you can
```go
v := value.Of(ptr)
```
```go
return value.OfFirstNotNil(ptr1, ptr2)
```
If you need a non-zero default value, there are `OfOrDefault` and `OfFirstNotNilOrDefault` variants as well!

If you want to receive an error if all pointers are nil, there is `OfFirstNotNilOrError` variant available.

## How to use

> By the spirit of "a little copying is better than a little dependency", I encourage you to
> simply copy the `Of`/`OfFirstNotNil` functions in your project. However, if you don't mind
> having this dependency:

__As this implementation relies on generics, you need to be using Go >1.18.__

1. With Go installed, run
```commandline
go get -u github.com/mrkagelui/value
```
2. Import this:
```go
import "github.com/mrkagelui/value"
```
3. Start using it!
```go
package main

import "github.com/mrkagelui/value"

func main() {
	var ptr *int
	print(value.Of(ptr))
}
```

## Frequently Asked Question

### Why is there not `OfOrError`?

Because it would look like
```go
if v, err := value.OfOrError(p); err != nil {
	// handle error
}
```
which is hardly shorter than
```go
if p == nil {
	// handle nil pointer
}
```
