package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Reddit struct {
	Data struct {
		After    string `json:"after"`
		Children []struct {
			RedditData `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

type RedditData struct {
	Author string  `json:"author"`
	Over18 bool    `json:"over_18"`
	Score  float64 `json:"score"`
	Title  string  `json:"title"`
	URL    string  `json:"url"`
}

func awww(after string) ([]RedditData, string, error) {
	resp, err := http.Get("https://www.reddit.com/r/awww.json?after=%s", after)
	if err != nil {
		return nil, after, err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)

	var r Reddit
	err = dec.Decode(&r)
	if err != nil {
		return nil, after, err
	}

	var rd []RedditData
	for _, d := range r.Data.Children {
		rd = append(rd, d.RedditData)
	}

	return rd, r.Data.After, nil
}

func main() {
	data, after, err := awww("")
	fmt.Println(data, after)
	time.Sleep(2 * time.Second)

	for err != nil {
		data, after, err = awww(after)
		fmt.Println(data, after)

		time.Sleep(2 * time.Second)
	}
}
