package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
	"io"
	"math/rand"
	"os"
	"bufio"
)

type Result struct {
	URL           string
	Header        string
	StatusCode    int
	ContentLength int64
}

func main() {
	urlPtr := flag.String("url", "", "URL to make requests to")
	headersFilePtr := flag.String("headers", "", "File containing headers for requests")
	proxyPtr := flag.String("proxy", "", "Proxy server IP:PORT (e.g., 127.0.0.1:8080)")
	quietPtr := flag.Bool("q", false, "Suppress banner")
	flag.Parse()
	log.SetFlags(0)
	    // Print tool banner

    if !*quietPtr {
    log.Print(`


	   __               __                      
	  / /  ___ ___  ___/ /__ _______ _    _____ 
	 / _ \/ -_) _ \/ _  / -_) __/ _ \ |/|/ / _ \
	/_//_/\__/\_,_/\_,_/\__/_/ / .__/__,__/_//_/
	                          /_/               
    
`)
	}

	if *urlPtr == "" {
		fmt.Println("Please provide a valid URL using the -url flag")
		return
	}

	if *headersFilePtr == "" {
		fmt.Println("Please provide a valid headers file using the -headers flag")
		return
	}

	headers, err := readHeadersFromFile(*headersFilePtr)
	if err != nil {
		fmt.Println("Error reading headers:", err)
		return
	}

	var wg sync.WaitGroup
	results := make(chan Result)

	for _, header := range headers {
		wg.Add(1)
		go func(header string) {
			defer wg.Done()

			response, err := makeRequest(*urlPtr, header, *proxyPtr)
			if err != nil {
				return
			}

			result := Result{
				URL:           *urlPtr + "?cachebuster=" + generateCacheBuster(),
				Header:        header,
				StatusCode:    response.StatusCode,
				ContentLength: response.ContentLength,
			}
			results <- result
		}(header)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	printResults(results)
}

func readHeadersFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	headers := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		headers = append(headers, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return headers, nil
}

func makeRequest(baseURL, header, proxy string) (*http.Response, error) {
	urlWithBuster := baseURL + "?cachebuster=" + generateCacheBuster()
	headers := parseHeaders(header)

	req, err := http.NewRequest("GET", urlWithBuster, nil)
	if err != nil {
		return nil, err
	}

	for _, h := range headers {
		parts := strings.SplitN(h, ": ", 2)
		if len(parts) == 2 {
			req.Header.Add(parts[0], parts[1])
		}
	}

	client := &http.Client{}
	if proxy != "" {
		proxyURL, err := url.Parse("http://" + proxy)
		if err != nil {
			fmt.Println("Error parsing proxy URL:", err)
			return nil, err
		}
		transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
		client = &http.Client{Transport: transport}
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if response.ContentLength >= 0 {
		return response, nil
	}

	body, err := io.ReadAll(response.Body)
	if err == nil {
		response.ContentLength = int64(len(body))
	}
	return response, nil
}

func parseHeaders(header string) []string {
	return strings.Split(header, "\n")
}

func generateCacheBuster() string {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 10)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func printResults(results <-chan Result) {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	for result := range results {
		statusColorFunc := red
		if result.StatusCode == 200 {
			statusColorFunc = green
		}

		statusOutput := statusColorFunc(fmt.Sprintf("[%d]", result.StatusCode))
		contentLengthOutput := magenta(fmt.Sprintf("[CL: %d]", result.ContentLength))
		headerOutput := cyan(fmt.Sprintf("[%s]", result.Header))

		parsedURL, _ := url.Parse(result.URL)
		query := parsedURL.Query()
		query.Del("cachebuster")
		parsedURL.RawQuery = query.Encode()
		urlOutput := yellow(fmt.Sprintf("[%s]", parsedURL.String()))

		resultOutput := fmt.Sprintf("%s %s %s %s", statusOutput, contentLengthOutput, headerOutput, urlOutput)
		fmt.Println(resultOutput)
	}
}
