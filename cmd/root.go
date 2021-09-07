package cmd

import "os"

func Init() {
	overwrite, indent, multipleArgs, files := getParams()

	if multipleArgs {
		for _, file := range files {
			formatFile(file, *indent, *overwrite)
		}
	} else {
		formatStream(os.Stdin, os.Stdout, *indent)
	}
}
