package main

import (
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"strings"

	"github.com/gocolly/colly/v2"
)

var htmlTemplate = `
<a href="{{.URL}}" target="_blank">
<img src="{{.Image}}" alt="{{.Title}}" />
</a>

<p>
</p>
`

type Data struct {
	Title string
	URL   string
	Image string
}

func main() {
	os.Exit(_main())
}

func _main() int {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s url\n", os.Args[0])
		return 1
	}

	t, err := template.New("test").Parse(htmlTemplate)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	url := os.Args[1]
	c := colly.NewCollector()

	var d Data
	d.URL = url

	c.OnHTML("a.sokmil-lightbox-jacket", func(e *colly.HTMLElement) {
		if d.Image == "" {
			d.Image = e.Attr("href")
		}
	})
	c.OnHTML("h1.page-title", func(e *colly.HTMLElement) {
		d.Title = e.Text
	})

	if err := c.Visit(url); err != nil {
		fmt.Println(err)
		return 1
	}

	if err := t.Execute(os.Stdout, &d); err != nil {
		fmt.Println(err)
		return 1
	}

	if err := copyToClipboard(d.Title); err != nil {
		fmt.Printf("failed to copy title to clipboard %v\n", err)
	}

	return 0
}

func copyToClipboard(text string) error {
	cmd := exec.Command("xsel", "--input", "--clipboard")
	cmd.Stdin = strings.NewReader(text)
	return cmd.Run()
}
