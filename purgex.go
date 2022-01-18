package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/jayateertha043/purgex/pkg/httpclient"
)

func main() {

	THREADS := flag.Int("t", 8, "Enter amount of threads")
	TIMEOUT := flag.Int("timeout", 3, "Enter request timeout in seconds")
	PROXY := flag.String("proxy", "", "Use custom proxy [http://ip:port or https://ip:port]")
	headersF := flag.String("headers", "", "To use Custom Headers headers.json file")
	NOBANNER := flag.Bool("nobanner", false, "Disable Banner")
	NOSTATUS := flag.Bool("nostatus", false, "Outputs only urls with status code between 100-300")
	MAXREQUEST := flag.Int("maxrequest", 1000, "Maximum requests/urls to try")

	flag.Parse()

	urls := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		txt := scanner.Text()
		u, err := url.Parse(txt)
		if err == nil {
			if strings.HasPrefix(u.String(), "http://") || strings.HasPrefix(u.String(), "https://") {
				urls = append(urls, txt)
			}
		}
	}
	showBanner := !*NOBANNER
	maxrequest := *MAXREQUEST
	if showBanner {
		printBanner()
	}

	if *THREADS < 1 {
		fmt.Println("Please supply valid number of threads, exiting...")
		return
	}

	headers := make(map[string]string)
	*headersF = strings.TrimSpace(*headersF)
	if *headersF != "" {
		jsonFile, err := os.Open(*headersF)
		if err != nil {
			fmt.Println("Unable to find " + *headersF)
			return
		}
		defer jsonFile.Close()
		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			fmt.Println("Unable to read")
			return
		}
		err = json.Unmarshal(byteValue, &headers)
		if err != nil {
			fmt.Println("Json format invalid in headers.json")
			return
		}
	}

	if len(urls) < *THREADS {
		*THREADS = len(urls)
	}

	sem := make(chan bool, *THREADS)
	var wg = new(sync.WaitGroup)

	for i := range urls {

		if i+1 >= maxrequest {
			break
		}

		U := urls[i]
		sem <- true
		wg.Add(1)
		go func(string) {
			defer func() {
				<-sem
				defer wg.Done()
			}()
			status_code, err := httpclient.PurgeRequest(U, headers, *TIMEOUT, *PROXY)
			if err == nil {
				if *NOSTATUS {
					if status_code < 300 && status_code > 100 {
						fmt.Println(U)
					}
				} else {
					fmt.Println(U + "\t" + strconv.Itoa(status_code))
				}

			}
		}(U)
	}
	wg.Wait()
}

func printBanner() {
	banner := `.______    __    __  .______        _______  _______ .______      
	|   _  \  |  |  |  | |   _  \      /  _____||   ____||   _  \     
	|  |_)  | |  |  |  | |  |_)  |    |  |  __  |  |__   |  |_)  |    
	|   ___/  |  |  |  | |      /     |  | |_ | |   __|  |      /     
	|  |      |  '--'  | |  |\  \----.|  |__| | |  |____ |  |\  \----.
	| _|       \______/  | _| '._____| \______| |_______|| _| '._____|
																	  `
	fmt.Println(banner)
}
