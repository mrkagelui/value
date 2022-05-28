# value.Of anything!

> Check out [my ptr package](https://github.com/mrkagelui/ptr) to do ptr.Of()!

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mrkagelui/value)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/mrkagelui/value/testing)
[![codecov](https://codecov.io/gh/mrkagelui/value/branch/master/graph/badge.svg?token=A2SCYKTWFZ)](https://codecov.io/gh/mrkagelui/value)
[![Go Report Card](https://goreportcard.com/badge/github.com/mrkagelui/value)](https://goreportcard.com/report/github.com/mrkagelui/value)
[![GitHub license](https://badgen.net/github/license/mrkagelui/ptr)](https://github.com/mrkagelui/ptr/blob/master/LICENSE)
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
	v = *ptr1
}
if ptr2 != nil {
	v = *ptr2
}
```
With this utility, you can
```go
v := value.Of(ptr)
```
```go
v := value.OfFirstNotNil(ptr1, ptr2)
```
If you need a non-zero default value, there are `OfOrDefault` and `OfFirstNotNilOrDefault` variants as well!
