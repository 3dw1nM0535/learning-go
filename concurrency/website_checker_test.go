package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func slowWebsiteChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func mockWebsiteChecker(url string) bool {
	if url == "waat://allmighty.god" {
		return false
	}
	return true
}

func TestWebsiteChecker(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://yahoo.com",
		"https://github.com",
		"waat://allmighty.god",
	}

	want := map[string]bool{
		"https://google.com":   true,
		"https://yahoo.com":    true,
		"https://github.com":   true,
		"waat://allmighty.god": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v expected %v", got, want)
	}
}

func BenchmarkWebsiteChecker(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowWebsiteChecker, urls)
	}
}
