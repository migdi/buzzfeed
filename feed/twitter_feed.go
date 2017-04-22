package feed

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

type TwitterFeed struct {
	ApiClient TwitterApi
}

type TwitterApi interface {
	GetUserTimeline(params url.Values) ([]anaconda.Tweet, error)
}

type Tweet struct {
	Title string
}

func NewTwitterFeed(consumerKey, consumerSecret string) *TwitterFeed {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	return &TwitterFeed{
		ApiClient: anaconda.NewTwitterApi("", ""),
	}
}

func (f *TwitterFeed) GetLatest(screenName string) (Tweet, error) {
	tweets, err := f.ApiClient.GetUserTimeline(url.Values{
		"screen_name": []string{screenName},
		"count":       []string{"1"},
	})

	if err != nil {
		return Tweet{}, err
	}

	return Tweet{Title: tweets[0].Text}, nil
}
