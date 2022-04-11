package main

import (
	"flag"
	"fmt"
	"go_hw_03/pkg/crawler"
	"go_hw_03/pkg/crawler/spider"
	"go_hw_03/pkg/index"
	"sort"
	"strings"
)

func main() {
	sFlag := flag.String("s", "", "Введите слова для поиска по флагу -s")

	flag.Parse()

	urls := []string{"https://go.dev/", "https://golang.org/", "https://rubyonrails.org"}

	docs := parseDocs(urls)

	indx := index.Index(docs)

	sort.Sort(byID(docs))

	// поиск по флагу s
	for i, v := range indx {
		if strings.Contains(i, *sFlag) {
			for i := 0; i < len(v); i++ {
				fmt.Println("Вхождения найдены в документах: ", binarySearch(v[i], docs))
			}
		}
	}

}

// Для сортировки документов
type byID []crawler.Document

func (a byID) Len() int           { return len(a) }
func (a byID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byID) Less(i, j int) bool { return a[i].ID < a[j].ID }

func parseDocs(urls []string) []crawler.Document {
	var docs = []crawler.Document{}

	s := spider.New()

	for _, v := range urls {
		doc, err := s.Scan(v, 1)
		if err != nil {
			continue
		}
		docs = append(docs, doc...)
	}

	return docs
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
