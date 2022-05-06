package concurrency

type WebsiteChecker func(string) bool

// go channel to communicate btwn website check and its goroutines
// to avoid race condition bug
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	resultChannel := make(chan result)
	checks := make(map[string]bool)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		checks[r.string] = r.bool
	}

	return checks
}
