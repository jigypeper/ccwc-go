package ccwc

import (
	"strings"
	"unicode/utf8"
)

type Count struct {
	bytes, chars, words, lines int // Changed the field name from 'byte' to 'bytes'
}

func (c *Count) GetStats(s string) { // Changed to pointer receiver and idiomatic method name
	c.bytes = len(s)
	c.chars = utf8.RuneCountInString(s)
	c.words = len(strings.Fields(s))
	c.lines = len(strings.Split(s, "\n"))
}
