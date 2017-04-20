package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	consumerKey := ""
	consumerSecret := ""

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)

	api := anaconda.NewTwitterApi("", "")

	tweets, err := api.GetUserTimeline(url.Values{
		"screen_name": []string{"ge_botafogo"},
		"count":       []string{"1"},
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tweets[0].Text)

}
