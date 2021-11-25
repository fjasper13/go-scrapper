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

type Intern struct {
	Profile string `json:"profile"`
	Name    string `json:"name"`
	Place   string `json:"place"`
	Date    string `json:"date"`
	// Duration string `json:"duration"`
	Stipend string `json:"stipend"`
	Apply   string `json:"apply"`
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
	mainUrl := "https://internshala.com/internships/page-1"

	r, err := http.Get(mainUrl)
	check(err)
	if r.StatusCode > 400 {
		fmt.Println("Status Code :", r.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(r.Body)
	check(err)

	// MAKE HTML FILE FROM HTML
	// river, err := doc.Find("#list_container").Html()
	// check(err)
	// writeHtmlFile(river, "index.html")

	// MAKE JSON FILE
	arrTech := make([]Intern, 0)

	doc.Find(".internship_meta").Each(func(i int, s *goquery.Selection) {
		profile := strings.TrimSpace(s.Find(".profile").Text())
		name := strings.TrimSpace(s.Find(".company_name").Text())
		place := strings.TrimSpace(s.Find(".location_link").Text())
		date := strings.TrimSpace(s.Find("#start-date-first").Text())
		// duration := strings.TrimSpace(s.Find(".other_detail_item").Find(".item_body").Text())
		stipend := strings.TrimSpace(s.Find(".stipend").Text())
		apply := strings.TrimSpace(s.Find(".apply_by").Find(".item_body").Text())

		//JSON
		tech := Intern{
			Profile: profile,
			Name:    name,
			Place:   place,
			Date:    date,
			// Duration: duration,
			Stipend: stipend,
			Apply:   apply,
		}

		arrTech = append(arrTech, tech)

	})
	// JSON
	writeJSON(arrTech)

}

func writeJSON(data []Intern) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create json file")
		return
	}

	_ = ioutil.WriteFile("test.json", file, 0644)
}
