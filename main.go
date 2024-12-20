package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"sync"

	"github.com/go-rod/rod"
)

func isValidProductURL(url string) bool {
	productPatterns := []string{
		`/product/`,
		`/item/`,
		`/p/`,
		`/products/`,
	}

	for _, pattern := range productPatterns {
		matched, _ := regexp.MatchString(pattern, url)
		if matched {
			return true
		}
	}
	return false
}

func scrapeWithRod(domain string) []string {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage(fmt.Sprintf("https://%s", domain))
	page.MustWaitLoad()

	links := page.MustElements("a[href]")
	productURLs := []string{}

	for _, link := range links {
		url := link.MustProperty("href").String()
		if isValidProductURL(url) {
			productURLs = append(productURLs, url)
		}
	}
	return productURLs
}

func main() {
	domains := []string{
		"tirabeauty.com",
		"beminimalist.co",
		"houseofbrahma.in",
	}

	results := make(map[string][]string)
	var mutex sync.Mutex
	var wg sync.WaitGroup

	for _, domain := range domains {
		wg.Add(1)
		go func(d string) {
			defer wg.Done()
			urls := scrapeWithRod(d)
			mutex.Lock()
			results[d] = urls
			mutex.Unlock()
		}(domain)
	}

	wg.Wait()

	file, err := os.Create("dynamic_product_urls.json")
	if err != nil {
		log.Fatalf("Failed to save results: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(results)
	if err != nil {
		log.Fatalf("Failed to encode results: %v", err)
	}

	fmt.Println("Results saved to dynamic_product_urls.json")
}
