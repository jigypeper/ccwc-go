package ccwc

import (
	"strings"
	"unicode/utf8"
)

type Count struct {
	byte, chars, words, lines int
}

func (c Count) get_stats(s string) Count {
	c.byte = len(s)
	c.chars = utf8.RuneCountInString(s)
	c.words = len(strings.Fields(s))
	c.lines = len(strings.Split(s, "\n"))

	return c
}
