package main

import (
	"flag"
	"fmt"
	"go_hw_02/pkg/crawler"
	"go_hw_02/pkg/index"
	"strings"
)

func main() {
	sFlag := flag.String("s", "", "help message for flag s")
	flag.Parse()

	u := []string{"https://go.dev/", "https://golang.org/", "https://rubyonrails.org"}

	i, ds := index.Index(u)

	for i, v := range i {
		c := strings.Contains(i, *sFlag)
		if c && *sFlag != "" {
			for i := 0; i < len(v); i++ {
				fmt.Println("Вхождения найдены в документах: ", binarySearch(v[i], ds))
			}
		}
	}

}

func binarySearch(needle int, haystack []crawler.Document) crawler.Document {

	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2

		if haystack[median].ID < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low].ID != needle {
		return haystack[low]
	}

	return haystack[low]
}
