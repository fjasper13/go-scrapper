package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

func main() {

	name := "data.csv"
	file, err := os.Create(name)
	if err != nil {
		log.Fatalf("could not create file, err :%q", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(
		colly.AllowedDomains("internshala.com"),
	)
	c.OnHTML(".internship_meta", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText("a"),
			e.ChildText("span"),
		})
	})

	for i := 0; i < 312; i++ {
		fmt.Printf("Scrapping Page : %d\n", i)

		c.Visit("https://internshala.com/internships/page-" + strconv.Itoa(i))
	}

	log.Printf("Scrapping Complete\n")
	log.Println(c)
}
