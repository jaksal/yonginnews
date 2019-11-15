package main

import (
	"bytes"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Item list of board title
type Item struct {
	No     string
	Link   string
	Title  string
	Writer string
	Date   string
}

func parseList(r []byte) ([]Item, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(r))
	if err != nil {
		return nil, err
	}

	var result []Item
	doc.Find(".t_list > table > tbody > tr").Each(func(i int, s *goquery.Selection) {
		f1 := s.Children().First()
		f2 := f1.Next()
		f3 := f2.Next()
		f4 := f3.Next()
		f5 := f4.Next()

		var item Item
		item.No = strings.TrimSpace(f1.Text())
		item.Link, _ = f2.Find("a").Attr("href")
		item.Title = strings.TrimSpace(f2.Text())
		item.Writer = strings.TrimSpace(f4.Text())
		item.Date = strings.TrimSpace(f5.Text())

		result = append(result, item)
	})
	return result, nil
}

func parseContents(r []byte) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(r))
	if err != nil {
		return nil, err
	}

	var result []string
	doc.Find(".video > iframe").Each(func(i int, s *goquery.Selection) {
		if link, exist := s.Attr("src"); exist {
			log.Println("find contents ...", link)
			temps := strings.Split(link, "?")
			result = append(result, temps[0])
		}
	})
	return result, nil
}
