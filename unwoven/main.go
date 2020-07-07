package main

import (
	"fmt"
	"net/http"
)

// - check all are urls
func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://binance.com/", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Request.URL)
}
