package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
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

// parseIp validates and returns the ip from raw
func parseIp(raw string) (string, error) {
	if ip := net.ParseIP(raw); ip != nil {
		return raw, nil
	} else {
		return "", errors.New(fmt.Sprintf("input provided(%s) is not a valid IP", raw))
	}
}

// ipsFromScanner ...
func ipsFromScanner(scanner *bufio.Scanner) ([]string, error) {
	ips := make([]string, 0)
	for scanner.Scan() {
		newip, err := parseIp(scanner.Text())
		if err != nil {
			return nil, err
		}
		// drop dups
		for _, ip := range ips {
			if newip == ip {
				continue
			}
		}
		ips = append(ips, newip)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return ips, nil
}
