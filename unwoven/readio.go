package main

import (
	"bufio"
	"errors"
	"os"
	"sort"
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
	nmaps := make([]string, 0)
	client := initClient()

	for scanner.Scan() {
		newnmap, err := getFinalUrl(scanner.Text(), client)
		if err != nil {
			return nil, err
		}
		// Unique input
		if !inSlice(nmaps, newnmap, func(a, b string) bool {
			return a == b
		}) {
			nmaps = append(nmaps, newnmap)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	sort.Strings(nmaps)

	urls := make([]string, 0)
	for j := 0; j < len(nmaps); j++ {
		if !inSlice(nmaps[j+1:], nmaps[j], urlEqual) {
			urls = append(urls, nmaps[j])
		}
	}
	return urls, nil
}
