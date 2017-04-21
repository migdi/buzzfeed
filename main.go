package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_CONSUMER_SECRET"))

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
