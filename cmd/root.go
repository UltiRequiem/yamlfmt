package cmd

import (
	. "github.com/UltiRequiem/yamlfmt/pkg"
	"log"
	"os"
)

func Init() {
	overwrite, indent, logFiles, multipleArgs, files := getParams()

	if multipleArgs {
		channel := make(chan string)

		// Send All Gophers
		for _, file := range files {
			go FormatFile(file, indent, overwrite, channel)
		}

		// Call All Gophers
		for i := 0; i < len(files); i++ {
			message := <-channel

			if logFiles {
				log.Print(message)
			}

		}

		return
	}

	FormatStream(os.Stdin, os.Stdout, indent)
}
