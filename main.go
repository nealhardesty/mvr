package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func mvr(regexString string, regexReplaceString string, files []string, dryRun bool) error {
	var lastErr error = nil
	r := regexp.MustCompile(regexString)
	for _, filename := range files {
		replacementFilename := r.ReplaceAllString(filename, regexReplaceString)
		fmt.Printf("'%s' > '%s' ... ", filename, replacementFilename)

		if !dryRun {
			err := os.Rename(filename, replacementFilename)
			if err != nil {
				fmt.Printf("%v.\n", err)
				lastErr = err
			} else {
				fmt.Printf("done.\n")
			}
		} else {
			fmt.Printf("dry run.\n")
		}
	}
	return lastErr
}

func mvr_noregex(str string, replaceStr string, files []string, dryRun bool) error {
	var lastErr error = nil
	for _, filename := range files {
		replacementFilename := strings.ReplaceAll(filename, str, replaceStr)
		fmt.Printf("'%s' > '%s' ... ", filename, replacementFilename)

		if !dryRun {
			err := os.Rename(filename, replacementFilename)
			if err != nil {
				fmt.Printf("%v.\n", err)
				lastErr = err
			} else {
				fmt.Printf("done.\n")
			}
		} else {
			fmt.Printf("dry run.\n")
		}
	}
	return lastErr
}

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: %v [options] <regex string> <regex replacement string> [files]*\n\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Println(`
Regex Quick Reference (Go/RE2 syntax):
  .        any character              ^        start of string
  *        zero or more (greedy)       $        end of string
  +        one or more (greedy)       ?        zero or one
  [abc]    character class            [^abc]   negated class
  \d       digit [0-9]                \w       word char [A-Za-z0-9_]
  \s       whitespace                 (...)    capturing group
  $1, $2   backreferences in replacement string`)
	}

	dryRun := flag.Bool("d", false, "dry run only")
	noRegex := flag.Bool("x", false, "disable regex, string replacement only")
	showHelp := flag.Bool("h", false, "show help message")
	flag.Parse()
	remainingArgs := flag.Args()

	if *showHelp {
		flag.Usage()
		os.Exit(0)
	}
	if len(remainingArgs) < 2 {
		flag.Usage()
		os.Exit(2)
	}

	regex := remainingArgs[0]
	regexReplace := remainingArgs[1]
	files := remainingArgs[2:]

	var err error
	if *noRegex {
		err = mvr_noregex(regex, regexReplace, files, *dryRun)
	} else {
		err = mvr(regex, regexReplace, files, *dryRun)
	}
	if err != nil {
		os.Exit(1)
	} else {
		os.Exit(0)
	}

}
