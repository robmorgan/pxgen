package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Standard characters to generate passwords with
var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+,.?/:;{}[]`~")

func usage() {
	fmt.Fprintf(os.Stderr, "usage: pxgen [length] [count]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	var length = 20
	var count = 1
	flag.Usage = usage
	flag.Parse()
	s := flag.Arg(0)
	if len(s) > 0 {
		// string to int
		var err error
		length, err = strconv.Atoi(s)
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
	}
	i := flag.Arg(1)
	if len(i) > 0 {
		// string to int
		var err error
		count, err = strconv.Atoi(i)
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
	}
	for i := 0; i < count; i++ {
		fmt.Println(randChar(length, StdChars))
	}
}

func randChar(length int, chars []byte) string {
	newPword := make([]byte, length)
	randomData := make([]byte, length+(length/4)) // storage for random bytes.
	clen := byte(len(chars))
	maxrb := byte(256 - (256 % len(chars)))
	i := 0
	for {
		if _, err := io.ReadFull(rand.Reader, randomData); err != nil {
			return ""
		}
		for _, c := range randomData {
			if c >= maxrb {
				continue
			}
			newPword[i] = chars[c%clen]
			i++
			if i == length {
				return string(newPword)
			}
		}
	}
}
