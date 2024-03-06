package concurrency

import (
	"reflect"
	"testing"
	"time"
)

// mockWebsiteChecker simulates checking a website, for testing purposes.
func mockWebsiteChecker(url string) bool {
	if url == "http://test.com" {
		return false
	}
	return true
}

// TestCheckWebsites tests the CheckWebsites function to ensure it behaves correctly.
func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"https://www.python.org",
		"http://test.com",
	}

	want := map[string]bool{
		"http://google.com":      true,
		"https://www.python.org": true,
		"http://test.com":        false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}

// slowStubWebsiteChecker simulates a slow response from checking a website.
func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

// BenchmarkCheckWebsites benchmarks the CheckWebsites function for performance.
func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "http://test.com"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
