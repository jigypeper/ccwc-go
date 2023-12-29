package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"io"
	"bytes"
	"os"
	"flag"
)

var countFlag = flag.String("Count options", "all", "choose from bytes, characters, words, and lines for count")

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

func main {
	buf := &bytes.Buffer{}
	n, err := io.Copy(buf, os.Stdin)
	if err != nil {
		log.Fatalln(err)
	} else if n <= 1 { // buffer always contains '\n'
		log.Fatalln("no input provided")
	}
}
