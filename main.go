package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func printUsage() {
	fmt.Println("rematch is intended to be used with piped input, and a valid regex must be provided as its only argument.")
	fmt.Println("Example Usage: fortune | rematch [--option] '^\\d+'")

	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("  %s\t%s\n", f.Name, f.Usage)
	})
}

func main() {
	// Our Flags
	var posix, all bool

	flag.BoolVar(&posix, "posix", false, "Use POSIX regex.  Defaults to PCRE")
	flag.BoolVar(&all, "all", false, "Return all matches. Defaults to the first match")

	flag.Parse()

	if flag.NArg() < 1 {
		printUsage()
		return
	}

	// Check stdin for piped input
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeNamedPipe == 0 {
		printUsage()
		return
	}

	// Capture piped input, capturing runes
	reader := bufio.NewReader(os.Stdin)
	captured := &strings.Builder{}

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		captured.WriteRune(input)
	}

	// Compile our regex argument
	reString := flag.Args()[flag.NArg()-1]

	var re *regexp.Regexp

	if posix {
		re, err = regexp.CompilePOSIX(reString)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	} else {
		re, err = regexp.Compile(reString)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	// Print the first or all matches (as captures)
	if all {
		for _, m := range re.FindAllString(captured.String(), -1) {
			fmt.Println(m)
		}
	} else {
		fmt.Println(re.FindString(captured.String()))
	}
}
