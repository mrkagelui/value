// Package value provides utilities to access value of pointers
package value

// Of returns the pointed value of p, if p is nil, returns the zero value
func Of[T any](p *T) T {
	var v T
	if p != nil {
		return *p
	}
	return v
}

// OfOrDefault returns the pointed value of p, if p is nil, returns the given default
func OfOrDefault[T any](p *T, def T) T {
	if p != nil {
		return *p
	}
	return def
}

// OfFirstNotNil returns the pointed value of the first not nil pointer given
// if all pointers are nil, returns the zero value
func OfFirstNotNil[T any](ptrs ...*T) T {
	var v T
	for _, p := range ptrs {
		if p != nil {
			return *p
		}
	}
	return v
}

// OfFirstNotNilOrDefault returns the pointed value of the first not nil pointer given
// if all pointers are nil, returns the given default
func OfFirstNotNilOrDefault[T any](def T, ptrs ...*T) T {
	for _, p := range ptrs {
		if p != nil {
			return *p
		}
	}
	return def
}
