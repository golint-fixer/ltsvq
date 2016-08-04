package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var labelsStr string
	flag.StringVar(&labelsStr, "l", "", "labels (ex. time,url,status)")
	var filename string
	flag.StringVar(&filename, "f", "-", "filename")
	flag.Parse()

	labels := bytes.Split([]byte(labelsStr), []byte{','})
	if len(labels) == 0 {
		fmt.Fprintf(os.Stderr, "labels must be specified with -l option\n")
		os.Exit(1)
	}

	var file io.ReadCloser
	if filename == "-" {
		file = os.Stdin
	} else {
		var err error
		file, err = os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ltsvq fails to open input file: %s\n", err)
			os.Exit(1)
		}
		defer file.Close()
	}

	labelAndColons := make([][]byte, 0, len(labels))
	for _, label := range labels {
		labelAndColon := append(label, ':')
		labelAndColons = append(labelAndColons, labelAndColon)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		terms := bytes.Split(scanner.Bytes(), []byte{'\t'})

		var b []byte
		for _, labelAndColon := range labelAndColons {
			for _, term := range terms {
				if bytes.HasPrefix(term, labelAndColon) {
					if len(b) > 0 {
						b = append(b, '\t')
					}
					b = append(b, term...)
					break
				}
			}
		}
		fmt.Println(string(b))
	}
	err := scanner.Err()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ltsvq fails to read record: %s\n", err)
		os.Exit(1)
	}
}
