package cmd

import (
	"github.com/UltiRequiem/yamlfmt/pkg"
	"log"
	"os"
)

func Init() {
	overwrite, indent, logFiles, multipleArgs, files := getParams()

	if multipleArgs {
		channel := make(chan string)

		for _, file := range files {
			go yamlfmt.FormatFile(file, *indent, *overwrite, channel)
		}

		for i := 0; i < len(files); i++ {
			logFile := <-channel

			if logFiles {
				log.Print(logFile)
			}

		}

		return
	}

	yamlfmt.FormatStream(os.Stdin, os.Stdout, *indent)
}
