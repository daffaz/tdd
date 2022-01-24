package concurrency

import (
	"testing"
	"time"
)

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i += 1 {
		urls[i] = "an url"
	}
	b.ResetTimer()
	CheckWebsites(slowStubWebsiteChecker, urls)

	for i := 0; i < b.N; i += 1 {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
