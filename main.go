package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
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

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: %v [options] <regex string> <regex replacement string> [files]*\n", os.Args[0])
		flag.PrintDefaults()
	}

	dryRun := flag.Bool("d", false, "dry run only")
	flag.Parse()
	remainingArgs := flag.Args()

	if len(remainingArgs) < 2 {
		flag.Usage()
		os.Exit(2)
	}

	regex := remainingArgs[0]
	regexReplace := remainingArgs[1]
	files := remainingArgs[2:]

	err := mvr(regex, regexReplace, files, *dryRun)
	if err != nil {
		os.Exit(1)
	} else {
		os.Exit(0)
	}

}
