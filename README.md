# Concurrent Web Crawler

A simple, concurrent web crawler implemented in Go. This project demonstrates the use of goroutines, mutexes, and WaitGroups to create an efficient, thread-safe web crawler.

## Features

- Concurrent crawling using goroutines
- Depth-limited crawling
- Thread-safe tracking of visited URLs
- Fake fetcher for testing and demonstration

## Getting Started

### Prerequisites

- Go 1.11 or higher

### Running the Crawler

1. Clone this repository:
   ```
   git clone https://github.com/yourusername/concurrent-web-crawler.git
   cd concurrent-web-crawler
   ```

2. Run the main program:
   ```
   go run main.go
   ```

## How It Works

The crawler starts from a given URL and explores links up to a specified depth. It uses:

- Goroutines for concurrent crawling
- A mutex to ensure thread-safe access to the visited URLs map
- A WaitGroup to synchronize the completion of all crawling goroutines

The `fakeFetcher` simulates web page fetching for testing purposes, allowing the crawler to run without making actual HTTP requests.

## Customization

To use the crawler with real web pages, implement the `Fetcher` interface with actual HTTP requests instead of the `fakeFetcher`.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
