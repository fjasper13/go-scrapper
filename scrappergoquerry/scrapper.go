package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Tech struct {
	Title       string `json:"title"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func writeHtmlFile(data, filename string) {
	file, err := os.Create(filename)
	defer file.Close()
	check(err)

	file.WriteString(data)

}

func main() {
	mainUrl := "https://techcrunch.com/"

	r, err := http.Get(mainUrl)
	check(err)
	if r.StatusCode > 400 {
		fmt.Println("Status Code :", r.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(r.Body)
	check(err)

	// MAKE HTML FILE FROM HTML
	// river, err := doc.Find("div.river").Html()
	// check(err)
	// writeHtmlFile(river, "index.html")

	// MAKE CSV FILE
	// file, err := os.Create("test.csv")
	// check(err)

	// writer := csv.NewWriter(file)

	// MAKE JSON FILE
	arrTech := make([]Tech, 0)

	doc.Find("div.river").Find("div.post-block").Each(func(i int, s *goquery.Selection) {
		h2 := s.Find("h2")
		title := strings.TrimSpace(h2.Text())
		url, _ := h2.Find("a").Attr("href")
		desc := strings.TrimSpace(s.Find("div.post-block__content").Text())

		//JSON
		tech := Tech{
			Title:       title,
			Url:         url,
			Description: desc,
		}

		arrTech = append(arrTech, tech)

		// CSV
		// posts := []string{title, url, desc}
		// writer.Write(posts)
	})
	// JSON
	writeJSON(arrTech)

	//CSV
	// writer.Flush()
}

func writeJSON(data []Tech) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create json file")
		return
	}

	_ = ioutil.WriteFile("test.json", file, 0644)
}
