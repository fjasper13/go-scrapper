package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"log"

	"github.com/gocolly/colly"
)

type Fact struct {
	Title       string `json:"title"`
	Place       string `json:"place"`
	Description string `json:"description"`
}

func main() {
	allFacts := make([]Fact, 0)

	c := colly.NewCollector(
		colly.AllowedDomains("internshala.com"),
	)

	c.OnHTML(".individual_internship_header", func(e *colly.HTMLElement) {
		// placeName := e.ChildText("span")
		title := e..Find("profile").ChildText("a")
		// title := e.ChildAttr("profile", "a")

		// test := e.Attr("id")
		// fmt.Println("test : ", test)

		// factDesc := e.Text

		fact := Fact{
			Title: title,
			// Place: placeName,
			// Description: factDesc,
		}

		allFacts = append(allFacts, fact)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://internshala.com/internships/page-1")

	// enc := json.NewEncoder(os.Stdout)
	// enc.SetIndent("", " ")
	// enc.Encode(allFacts)

	writeJSON(allFacts)
}

func writeJSON(data []Fact) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create json file")
		return
	}

	_ = ioutil.WriteFile("test.json", file, 0644)
}
