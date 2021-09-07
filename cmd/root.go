package cmd

import (
	"log"
	"os"
)

// yamlfmt entry point
func Init() {
	overwrite, indent, logFiles, multipleArgs, files := getParams()

	if multipleArgs {
		for _, file := range files {
			formatFile(file, *indent, *overwrite)
			if logFiles {
				log.Printf("%s formatted!\n", file)
			}
		}
	} else {
		formatStream(os.Stdin, os.Stdout, *indent)
	}
}
