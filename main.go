package main

import (
	"fmt"
	"log"
	"net/url"	
	"os"
	"encoding/json"
	"github.com/ChimeraCoder/anaconda"
)

type Config struct {
	Twitter struct {
		Key string `json:"key"`
		Secret string `json:"secret"`
	} `json:"twitter"`
}


func loadConf(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	configFile.Close()
	return config
}


func main() {
  conf := loadConf("config.json")
		
	anaconda.SetConsumerKey(conf.Twitter.Key)
	anaconda.SetConsumerSecret(conf.Twitter.Secret)

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
