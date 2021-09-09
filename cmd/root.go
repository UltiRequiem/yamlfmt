package cmd

import (
	"log"
	"os"
)

// yamlfmt entry point
func Init() {
	overwrite, indent, logFiles, multipleArgs, files := getParams()

	if multipleArgs {
		channel := make(chan string)

		for _, file := range files {
			go formatFile(file, *indent, *overwrite, channel)
		}

		for i := 0; i < len(files); i++ {
			logFile := <-channel

			if logFiles {
				log.Print(logFile)
			}

		}

	} else {
		formatStream(os.Stdin, os.Stdout, *indent)
	}
}
