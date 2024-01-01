package main

import (
	"fmt"
	"io"
	"bytes"
	"os"
	"flag"
	"github.com/jigypeper/ccwc-go/src"
)

var countFlag = flag.String("Count options", "all", "choose from bytes, characters, words, and lines for count")

func main {
	buf := &bytes.Buffer{}
	n, err := io.Copy(buf, os.Stdin)
	if err != nil {
		log.Fatalln(err)
	} else if n <= 1 { // buffer always contains '\n'
		log.Fatalln("no input provided")
	}
}
