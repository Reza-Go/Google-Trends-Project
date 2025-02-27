package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	Title    string `xml:"title"`
	ItemList []Item `xml:"item"`
}

type Item struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	Traffic   string `xml:"approx_traffic"`
	NewsItems []News `xml:"news_item"`
}

type News struct {
	Headline     string `xml:"news_item_title"`
	HeadlineLink string `xml:"news_item_url"`
}

func main() {
	var r RSS
	data := readGoogleTrends()
	err := xml.Unmarshal(data, &r)

	if err != nil {
		fmt.Println("error:", err)
	}

	//Print Out
	fmt.Println("\n Below are all the google search trends for today")
	fmt.Println("---------------------------------------------------")

	for i := range r.Channel.ItemList {
		rank := (i + 1)
		fmt.Println("#", rank)
		fmt.Println("Search Term:", r.Channel.ItemList[i].Title)
		fmt.Println("Link to the Trend", r.Channel.ItemList[i].Link)
		fmt.Println("Headline:", r.Channel.ItemList[i].NewsItems[0].Headline)
		fmt.Println("LInk to Article:", r.Channel.ItemList[i].NewsItems[0].HeadlineLink)
		fmt.Println("-----------------------------------------------")
	}
}

func readGoogleTrends() []byte {
	resp := getGoogleTrends()
	data, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return data
}

func getGoogleTrends() *http.Response {
	resp, err := http.Get("https://trends.google.com/trending/rss?geo=US")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return resp
}
