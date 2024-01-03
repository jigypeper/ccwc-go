package ccwc

import "testing"

func TestCount(t *testing.T) {
	words := "Hello, world."
	want := Count{
		bytes: 13,
		chars: 10,
		words: 2,
		lines: 1,
	}
	var got Count
	if got.GetStats(words); got != want {
		t.Errorf("get_stats() = %q, want %q", got, want)
	}
}
