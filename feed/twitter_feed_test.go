package feed_test

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/migdi/buzzfeed/feed"
	"github.com/migdi/buzzfeed/feed/feedfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TwitterFeed", func() {
	var twitter *feed.TwitterFeed
	var fakeTwitterApi *feedfakes.FakeTwitterApi

	BeforeEach(func() {
		fakeTwitterApi = new(feedfakes.FakeTwitterApi)
		fakeTwitterApi.GetUserTimelineReturns([]anaconda.Tweet{
			{Text: "Tweet text"},
		}, nil)
		twitter = &feed.TwitterFeed{
			ApiClient: fakeTwitterApi,
		}
	})

	Describe("GetLatest", func() {
		It("pulls latest tweets from the user", func() {
			resp, err := twitter.GetLatest("twitteruser")

			Expect(err).NotTo(HaveOccurred())
			Expect(resp.Title).To(Equal("Tweet text"))
			Expect(fakeTwitterApi.GetUserTimelineArgsForCall(0)).To(Equal(url.Values{
				"screen_name": []string{"twitteruser"},
				"count":       []string{"1"},
			}))
		})
	})
})
