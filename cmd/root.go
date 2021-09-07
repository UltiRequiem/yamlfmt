package cmd

import (
	"fmt"
	"os"
)

func Init() {
	overwrite, indent, log, multipleArgs, files := getParams()

	if multipleArgs {
		for _, file := range files {
			formatFile(file, *indent, *overwrite)
			if log {
				fmt.Printf("%s formatted!\n", file)
			}
		}
	} else {
		formatStream(os.Stdin, os.Stdout, *indent)
	}
}
