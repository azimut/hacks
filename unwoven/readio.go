package main

import (
	"bufio"
	"errors"
	"os"
	"strings"
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

// urlEqual checks if urls are the same, meant to not consider trailing slash and keep one
func urlEqual(a, b string) bool {
	return strings.TrimRight(a, "/") == strings.TrimRight(b, "/")
}

func urlsFromScanner(scanner *bufio.Scanner) ([]string, error) {
	var found bool
	nmaps := make([]string, 0)
	client := initClient()

	for scanner.Scan() {
		newnmap, err := getFinalUrl(scanner.Text(), client)
		if err != nil {
			return nil, err
		}
		nmaps = append(nmaps, newnmap)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	urls := make([]string, 0)
	for j := 0; j < len(nmaps); j++ {
		found = false
		for _, value := range nmaps[j+1:] {
			if urlEqual(nmaps[j], value) {
				found = true
				break
			}
		}
		if !found {
			urls = append(urls, nmaps[j])
		}
	}
	return urls, nil
}
