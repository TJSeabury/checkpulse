package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	pulse("https://notasubdomain.marketmentorsdev.com/")
}

func pulse(domain string) {
	var client = &http.Client{
		Timeout: time.Second * 10,
	}

	res, err := client.Head(domain)
	if err != nil {
		if os.IsTimeout(err) {
			// timeout
			panic(err)
		} else {
			// panic: Head "https://ctd.marketmentorsdev.com/": dial tcp 45.33.70.122:443: connectex: No connection could be made because the target machine actively refused it.
			panic(err)
		}
	}

	fmt.Println("Status:", res.StatusCode)
	fmt.Println("ContentLength:", res.ContentLength)
	fmt.Println(res)
}
