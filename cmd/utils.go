package cmd

import (
	"flag"
)

func getParams() (*bool, *int, bool, []string) {
	overwrite := flag.Bool("w", false, "Overwrite the input file")
	indent := flag.Int("indent", 2, "Default indent")
	flag.Parse()

	multipleArgs := flag.NArg() > 0

	return overwrite, indent, multipleArgs, flag.Args()

}
