E-commerce Product URL Crawler
Welcome to the E-commerce Product URL Crawler! This tool is designed to scrape product URLs from dynamic e-commerce websites and save the results into a JSON file.

Features
Handles both static and dynamic websites using headless browsing.
Automatically detects product URLs based on common patterns.
Supports concurrent crawling for multiple domains.
Outputs structured data in dynamic_product_urls.json.
Getting Started
1. Clone the Repository
First, clone the repository to your local machine:

bash
Copy code
git clone <repository-url>
cd web-crawler
2. Install Go
Ensure you have Go installed. If not, download and install it from Go's official website.

Verify installation:

bash
Copy code
go version
3. Install Dependencies
Install the required libraries:

bash
Copy code
go mod tidy
4. Install Chromium
This project uses Rod, which requires Chromium for headless browsing.

macOS:
bash
Copy code
brew install chromium
Linux:
bash
Copy code
sudo apt-get install chromium-browser
Windows: Download Chromium from here.
Ensure Chromium is accessible in your system's PATH:

bash
Copy code
chromium --version
Configuration
Add Domains
Update the domains list in main.go with the e-commerce domains you want to crawl:

go
Copy code
domains := []string{
    "example1.com",
    "example2.com",
    "example3.com",
}
Modify URL Patterns
The crawler detects product URLs using predefined patterns. You can update these patterns in the isValidProductURL function:

go
Copy code
productPatterns := []string{
    `/product/`,
    `/item/`,
    `/p/`,
    `/products/`,
}
How to Run
1. Run the Crawler
Run the project directly using:

bash
Copy code
go run main.go
2. View the Output
The crawler saves the results to dynamic_product_urls.json in the following format:

json
Copy code
{
  "example1.com": [
    "https://example1.com/product/123",
    "https://example1.com/item/456"
  ],
  "example2.com": [
    "https://example2.com/products/789"
  ]
}
Troubleshooting
Chromium Not Found
Ensure Chromium is installed and properly configured in your system's PATH.

No Output in JSON
Verify the domain list.
Check the patterns in isValidProductURL to ensure they match the product URLs for your target websites.