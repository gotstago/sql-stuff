package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

func replace(re, repl, filename string) {
	regex, err := regexp.Compile(re)
	if err != nil {
		return // there was a problem with the regular expression.
	}

	fh, err := os.Open(filename)
	f := bufio.NewReader(fh)

	if err != nil {
		return // there was a problem opening the file.
	}
	defer fh.Close()

	buf := make([]byte, 1024)
	for {
		buf, _, err = f.ReadLine()
		if err != nil {
			return
		}

		s := string(buf)
		result := regex.ReplaceAllString(s, repl)
		fmt.Print(result + "\n")
	}
}

func main() {
	flag.Parse()
	if flag.NArg() == 3 {
		replace(flag.Arg(0), flag.Arg(1), flag.Arg(2))
	} else {
		fmt.Printf("Wrong number of arguments.\n")
	}
}
