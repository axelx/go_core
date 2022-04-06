package index

import (
	"go_hw_02/pkg/crawler"
	"go_hw_02/pkg/crawler/spider"
	"sort"
	"strings"
)

type ByID []crawler.Document

func (a ByID) Len() int           { return len(a) }
func (a ByID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByID) Less(i, j int) bool { return a[i].ID < a[j].ID }

func Index(l []string) (map[string][]int, []crawler.Document) {
	id := 1
	s := parseDocument(l)
	var in = make(map[string][]int)

	for i, v := range s {
		s[i].ID = id
		v.ID = id
		tmp := strings.Fields(v.Title)
		for i := 0; i < len(tmp); i++ {
			k := strings.ToLower(tmp[i])
			in[k] = append(in[k], v.ID)
		}
		id++
	}
	sort.Sort(ByID(s))

	return in, s
}

func parseDocument(u []string) []crawler.Document {
	var ad = []crawler.Document{}

	s := spider.New()

	for _, v := range u {
		list, err := s.Scan(v, 1)
		if err != nil {
			continue
		}
		ad = append(ad, list...)
	}

	return ad
}
