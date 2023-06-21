package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "https://") == false {
			url = "https://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("fetch:%v\n", err)
			os.Exit(1)
		}
		b := resp.Status
		fmt.Printf("%s", b)

	}
}
