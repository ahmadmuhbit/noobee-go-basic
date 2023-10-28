package main

import (
	"fmt"
	"net/url"
)

func main() {
	urlString := "https://noobee.id/search?kind=Course"

	url, err := url.Parse(urlString)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("URL:", url)
		fmt.Println("Protocol:", url.Scheme)
		fmt.Println("Host:", url.Host)
		fmt.Println("Path:", url.Path)
		fmt.Println("Raw Query:", url.RawQuery)
		fmt.Println("Query:", url.Query())
	}
}
