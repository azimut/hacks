package main

import (
	"bufio"
	"errors"
	"os"
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

func urlsFromScanner(scanner *bufio.Scanner) ([]string, error) {
	nmaps := make([]string, 0)
	client := initClient()
	for scanner.Scan() {
		newnmap, err := getFinalUrl(scanner.Text(), client)
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
