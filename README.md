# My Golang Job Scraper Project

This project provides a convenient way to search for Golang jobs across multiple websites, including Djinni.co and Dou.ua.

## Getting Started

**Prerequisites**

* Golang (version 1.22) 

**Installation**

1. Clone the repository:
   ```
   git clone https://github.com/fentezi/ScraperVacancy.git
   ```
   
2. Navigate to the project directory:
   ```
   cd ScraperVacancy
   ```

3. Install dependencies:
   ```
   go get -u ./...
   ```

4. By default, the server port is 8080. If you want to use a different port, then flags are used to specify the port.   
   Example:
   ```
   go run cmd/app.go -port=3030
   ```
