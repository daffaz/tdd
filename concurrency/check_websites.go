package concurrency

type WebsiteCheker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteCheker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i += 1 {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
