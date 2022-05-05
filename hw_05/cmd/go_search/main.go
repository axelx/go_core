package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_hw_03/pkg/crawler"
	"go_hw_03/pkg/crawler/spider"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	if fs, err := os.Stat("file.txt"); errors.Is(err, os.ErrNotExist) || fs.Size() == 0 {

		urls := []string{"https://go.dev/", "https://golang.org/", "https://rubyonrails.org"}
		docs := parseDocs(urls)

		dByte, err := json.Marshal(docs)
		if err != nil {
			log.Fatal(err)
		}

		f, err := os.Create("./file.txt")
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		err = ioutil.WriteFile(f.Name(), []byte(dByte), 0666)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("result------file does not exist----", docs)

	} else {
		// вариант чтения с помощью io.Reader
		var file *os.File
		file, err = os.Open(fs.Name())
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		buf := make([]byte, fs.Size())
		n, err := file.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		buf = buf[:n]

		var docs []crawler.Document
		json.Unmarshal(buf, &docs)

		fmt.Println("result------file exists----", docs)
	}

}

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
