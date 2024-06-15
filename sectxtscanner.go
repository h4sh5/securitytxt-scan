package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"log"
	"sync"
	"strings"
)


func RequestSecurityTxt(domain string) {
	urls := [4]string{}
	urls[0] = fmt.Sprintf("https://%s/.well-known/security.txt", domain)
	urls[1] = fmt.Sprintf("https://%s/security.txt", domain)
	urls[2] = fmt.Sprintf("http://%s/.well-known/security.txt", domain)
	urls[3] = fmt.Sprintf("http://%s/security.txt", domain)

	for _,requestURL := range urls {
		res, err := http.Get(requestURL)
		if err != nil {
			log.Printf("error making http request to %s: %s\n", domain, err)
			continue
		}


		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Printf("client: could not read response body: %s\n", err)
			continue
		}

		contentType := res.Header.Get("Content-Type")

		if res.StatusCode == 200 && strings.Contains(contentType, "text/plain") {
			log.Printf("200 text/plain on: %s\n", requestURL)
			fileName := fmt.Sprintf("raw/%s_security.txt", domain)
			os.WriteFile(fileName, resBody, 0644)
			return
		}
	}
	

}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Missing argument for input file of domains. Usage:\n%s <domains to scan>\n", os.Args[0])
		os.Exit(1)
	}

	var wg sync.WaitGroup

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	for _, domain := range strings.Split(string(data[:]), "\n") {

		fileName := fmt.Sprintf("raw/%s_security.txt", domain)

		if _, err := os.Stat(fileName); os.IsNotExist(err) {
			wg.Add(1) // increment wait group counter
			go func (d string) {
				defer wg.Done()
				RequestSecurityTxt(d)
			} (domain)
		}

		
	}

	// wait for all groups to finish
	wg.Wait()

}