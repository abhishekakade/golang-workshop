package main

import "fmt"

// GlobalLimit is max limit for cache
const GlobalLimit = 100

// MaxCacheSize is max cache size for cache
const MaxCacheSize int = 10 * GlobalLimit

// CacheKeyBook is prefix for books
// CacheKeyCD is prefix for CDs
const (
	CacheKeyBook = "book_"
	CacheKeyCD   = "cd_"
)

var cache map[string]string

func cacheGet(key string) string {
	return cache[key]
}

func cacheSet(key, val string) {
	if len(cache)+1 >= MaxCacheSize {
		return
	}
	cache[key] = val
}

func SetBook(isbn string, name string) {
	cacheSet(CacheKeyBook+isbn, name)
}

func GetBook(isbn string) string {
	return cacheGet(CacheKeyBook + isbn)
}

func SetCD(sku string, title string) {
	cacheSet(CacheKeyCD+sku, title)
}

func GetCD(sku string) string {
	return cacheGet(CacheKeyCD + sku)
}

func main() {
	cache = make(map[string]string)

	SetBook("1234-5678", "Get Ready To Go")
	SetCD("1234-5678", "Get Ready To Go Audio Book")

	fmt.Println("Book :", GetBook("1234-5678"))
	fmt.Println("CD   :", GetCD("1234-5678"))
}
