package cmd

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func getParams() (*bool, *int, bool, bool, []string) {
	overwrite := flag.Bool("w", false, "Overwrite the input file")
	indent := flag.Int("indent", 2, "Default indent")
	log := flag.Bool("l", false, "Log when a file is formatted")

	flag.Parse()

	return overwrite, indent, *log, flag.NArg() > 0, flag.Args()

}

func formatFile(file string, indent int, overwrite bool) {
	r, err := os.Open(file)

	defer r.Close()

	if err != nil {
		log.Fatalf("Error while reading the file %s: '%s'", file, err)
	}

	var out bytes.Buffer

	if err := formatStream(r, &out, indent); err != nil {
		log.Fatalf("Failed formatting YAML stream: '%s'", err)
	}

	if e := dumpStream(&out, file, overwrite); e != nil {
		log.Fatalf("Cannot overwrite %s: '%s'", file, e)
	}
}

func formatStream(r io.Reader, out io.Writer, indent int) error {

	d := yaml.NewDecoder(r)
	in := yaml.Node{}
	err := d.Decode(&in)

	for err == nil {
		e := yaml.NewEncoder(out)
		e.SetIndent(indent)
		if err := e.Encode(&in); err != nil {
			log.Fatal(err)
		}
		e.Close()

		if err = d.Decode(&in); err == nil {
			fmt.Fprintln(out, "---")
		}
	}

	if err != nil && err != io.EOF {
		return err
	}

	return nil
}

func dumpStream(out *bytes.Buffer, f string, overwrite bool) error {
	if overwrite {
		return os.WriteFile(f, out.Bytes(), 0744)
	}
	_, err := io.Copy(os.Stdout, out)
	return err
}
