package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gocolly/colly/v2"
)

func main() {
	os.Exit(_main())
}

func _main() int {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s url\n", os.Args[0])
		return 1
	}

	url := os.Args[1]
	c := colly.NewCollector()

	var cookies []*http.Cookie
	cookies = append(cookies, &http.Cookie{
		Name:   "age_check_done",
		Value:  "1",
		Path:   "/",
		Domain: ".dmm.co.jp",
	})

	if err := c.SetCookies("https://www.dmm.co.jp", cookies); err != nil {
		fmt.Println(err)
		return 1
	}

	c.OnHTML("head title", func(e *colly.HTMLElement) {
		fmt.Printf("Title='%s'\n", e.Text)
	})

	c.Visit(url)
	return 0
}
