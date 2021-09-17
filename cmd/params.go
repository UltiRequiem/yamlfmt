package cmd

import (
	"flag"
)

func getParams() (*bool, *int, bool, bool, []string) {
	overwrite := flag.Bool("w", false, "Overwrite the input file")
	indent := flag.Int("indent", 2, "Default indent")
	log := flag.Bool("l", false, "Log when a file is formatted")

	flag.Parse()

	return overwrite, indent, *log, flag.NArg() > 0, flag.Args()

}
