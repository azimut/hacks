package main

import (
	"bufio"
	"errors"
	"io/ioutil"
	"os"

	"github.com/tomsteele/go-nmap"
)

// errorPipeless errors if program is not called with pipe input
func errorPipeless() error {
	finfo, err := os.Stdin.Stat()
	if err != nil {
		return err
	}
	if finfo.Mode()&os.ModeNamedPipe == 0 {
		return errors.New("is not called with pipe data")
	}
	return nil
}

func parseFilePath(filepath string) (*nmap.NmapRun, error) {
	bfile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	parsed, err := nmap.Parse(bfile)
	if err != nil {
		return nil, err
	}
	return parsed, nil
}

func nmapsFromScanner(scanner *bufio.Scanner) ([]*nmap.NmapRun, error) {
	nmaps := make([]*nmap.NmapRun, 0)
	for scanner.Scan() {
		newnmap, err := parseFilePath(scanner.Text())
		if err != nil {
			return nil, err
		}
		// drop dups
		for _, nmap := range nmaps {
			if newnmap == nmap {
				continue
			}
		}
		nmaps = append(nmaps, newnmap)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return nmaps, nil
}
