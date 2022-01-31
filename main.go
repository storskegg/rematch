package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func printUsage() {
	fmt.Println("rematch is intended to be used with piped input, and a valid regex must be provided as its only argument.")
	fmt.Println("Example Usage: fortune | rematch '^\\d+'")
}

func main() {
	if len(os.Args) != 2 {
		printUsage()
		return
	}

	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeNamedPipe == 0 {
		printUsage()
		return
	}

	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	haystack := string(output)

	first := os.Args[1]

	re, err := regexp.Compile(first)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(re.FindString(haystack))
}
