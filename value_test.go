package value

import (
	"testing"
)

func equals[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestOf(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		num := 24
		equals(t, Of(&num), num)
	})
	t.Run("nil int ptr", func(t *testing.T) {
		var num *int
		equals(t, Of(num), 0)
	})
	t.Run("int8", func(t *testing.T) {
		var num int8 = 24
		equals(t, Of(&num), num)
	})
	t.Run("int16", func(t *testing.T) {
		var num int16 = 24
		equals(t, Of(&num), num)
	})
	t.Run("int32", func(t *testing.T) {
		var num int32 = 24
		equals(t, Of(&num), num)
	})
	t.Run("int64", func(t *testing.T) {
		var num int64 = 24
		equals(t, Of(&num), num)
	})
	t.Run("uint", func(t *testing.T) {
		var num uint = 24
		equals(t, Of(&num), num)
	})
	t.Run("uint8", func(t *testing.T) {
		var num uint8 = 24
		equals(t, Of(&num), num)
	})
	t.Run("uint16", func(t *testing.T) {
		var num uint16 = 24
		equals(t, Of(&num), num)
	})
	t.Run("uint32", func(t *testing.T) {
		var num uint32 = 24
		equals(t, Of(&num), num)
	})
	t.Run("uint64", func(t *testing.T) {
		var num uint64 = 24
		equals(t, Of(&num), num)
	})
	t.Run("uintptr", func(t *testing.T) {
		var num uintptr = 24
		equals(t, Of(&num), num)
	})
	t.Run("float32", func(t *testing.T) {
		var num float32 = 24
		equals(t, Of(&num), num)
	})
	t.Run("float64", func(t *testing.T) {
		var num float64 = 24
		equals(t, Of(&num), num)
	})
	t.Run("complex64", func(t *testing.T) {
		var num complex64 = 24 + 3i
		equals(t, Of(&num), num)
	})
	t.Run("complex128", func(t *testing.T) {
		num := 24.42 - 4i
		equals(t, Of(&num), num)
	})
	t.Run("byte", func(t *testing.T) {
		var b byte = 'a'
		equals(t, Of(&b), b)
	})
	t.Run("rune", func(t *testing.T) {
		var r rune = '义'
		equals(t, Of(&r), r)
	})
	t.Run("string", func(t *testing.T) {
		title := "maidenless"
		equals(t, Of(&title), title)
	})
	t.Run("nil string ptr", func(t *testing.T) {
		var title *string
		equals(t, Of(title), "")
	})
	t.Run("boolean", func(t *testing.T) {
		b := true
		equals(t, Of(&b), b)
	})
	t.Run("nil boolean ptr", func(t *testing.T) {
		var b *bool
		equals(t, Of(b), false)
	})
	t.Run("array", func(t *testing.T) {
		var arr = [9]int{32, 921}
		equals(t, Of(&arr), arr)
	})
	t.Run("struct", func(t *testing.T) {
		type player struct {
			title string
		}
		p := player{"foul tarnished"}
		equals(t, Of(&p), p)
	})
	t.Run("nil struct ptr", func(t *testing.T) {
		type player struct {
			title string
		}
		var p *player
		equals(t, Of(p), player{})
	})
	t.Run("channel", func(t *testing.T) {
		m := make(chan bool)
		equals(t, Of(&m), m)
	})
}

func TestOfOrDefault(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		num := 24
		equals(t, OfOrDefault(&num, 19), num)
	})
	t.Run("nil int ptr", func(t *testing.T) {
		var num *int
		equals(t, OfOrDefault(num, 18), 18)
	})
	t.Run("string", func(t *testing.T) {
		title := "maidenless"
		equals(t, OfOrDefault(&title, "a ditch somewhere"), title)
	})
	t.Run("nil string ptr", func(t *testing.T) {
		var title *string
		equals(t, OfOrDefault(title, "a ditch somewhere"), "a ditch somewhere")
	})
	t.Run("boolean", func(t *testing.T) {
		b := true
		equals(t, OfOrDefault(&b, false), b)
	})
	t.Run("nil boolean ptr", func(t *testing.T) {
		var b *bool
		equals(t, OfOrDefault(b, true), true)
	})
	t.Run("struct", func(t *testing.T) {
		type player struct {
			title string
		}
		p := player{"foul tarnished"}
		equals(t, OfOrDefault(&p, player{"thy strength befits a crown"}), p)
	})
	t.Run("nil struct ptr", func(t *testing.T) {
		type player struct {
			title string
		}
		var p *player
		equals(t, OfOrDefault(p, player{"thy strength befits a crown"}), player{"thy strength befits a crown"})
	})
}

func TestOfFirstNotNil(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		num := 24
		equals(t, OfFirstNotNil[int](nil, &num), num)
	})
	t.Run("nil int ptr", func(t *testing.T) {
		var numPtr *int
		equals(t, OfFirstNotNil(numPtr, numPtr, nil), 0)
	})
	t.Run("string", func(t *testing.T) {
		title := "maidenless"
		equals(t, OfFirstNotNil[string](nil, &title), title)
	})
	t.Run("nil string ptr", func(t *testing.T) {
		var title *string
		equals(t, OfFirstNotNil(title, title, nil), "")
	})
	t.Run("boolean", func(t *testing.T) {
		b := true
		equals(t, OfFirstNotNil[bool](nil, &b), b)
	})
	t.Run("nil boolean ptr", func(t *testing.T) {
		var b *bool
		equals(t, OfFirstNotNil(b, b, nil), false)
	})
	t.Run("struct", func(t *testing.T) {
		type player struct {
			title string
		}
		p := player{"foul tarnished"}
		equals(t, OfFirstNotNil[player](nil, &p, &player{"thy strength befits a crown"}), p)
	})
	t.Run("nil struct ptr", func(t *testing.T) {
		type player struct {
			title string
		}
		var p *player
		equals(t, OfFirstNotNil(p, nil, nil), player{""})
	})
}

func TestOfFirstNotNilOrDefault(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		num := 24
		equals(t, OfFirstNotNilOrDefault(1, nil, &num), num)
	})
	t.Run("nil int ptr", func(t *testing.T) {
		var numPtr *int
		equals(t, OfFirstNotNilOrDefault(92, numPtr, nil), 92)
	})
	t.Run("string", func(t *testing.T) {
		title := "maidenless"
		equals(t, OfFirstNotNilOrDefault("tarnished", nil, &title), title)
	})
	t.Run("nil string ptr", func(t *testing.T) {
		var title *string
		equals(t, OfFirstNotNilOrDefault("traveler from beyond the fog", title, title, nil), "traveler from beyond the fog")
	})
	t.Run("boolean", func(t *testing.T) {
		b := true
		equals(t, OfFirstNotNilOrDefault(false, nil, &b), b)
	})
	t.Run("nil boolean ptr", func(t *testing.T) {
		var b *bool
		equals(t, OfFirstNotNilOrDefault(true, b, b, nil), true)
	})
	t.Run("struct", func(t *testing.T) {
		type player struct {
			title string
		}
		p := player{"foul tarnished"}
		equals(t, OfFirstNotNilOrDefault(player{"I have never known defeat"}, &p, &player{"thy strength befits a crown"}), p)
	})
	t.Run("nil struct ptr", func(t *testing.T) {
		type player struct {
			title string
		}
		var p *player
		equals(t, OfFirstNotNilOrDefault(player{"I have never known defeat"}, p, nil, nil), player{"I have never known defeat"})
	})
}
