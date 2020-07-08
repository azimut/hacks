package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// initClient returns an Initialized client
func initClient() *http.Client {
	return &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        10,
			MaxIdleConnsPerHost: 10,
			MaxConnsPerHost:     10,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}

// TODO: clean up url? like port :80 :443 or trailing slash?
func parseUrl(rawurl string) (string, error) {
	_, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}
	return rawurl, nil
}

// getFinalUrl returns the final resolved url after GET the url
func getFinalUrl(rawurl string, client *http.Client) (string, error) {
	url, err := parseUrl(rawurl)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return fmt.Sprint(resp.Request.URL), nil
}
