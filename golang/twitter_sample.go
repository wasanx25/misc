package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Tweets []twitter.Tweet

func (t Tweets) Len() int {
	return len(t)
}

func (t Tweets) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

type ByCreatedAt struct {
	Tweets
}

func (t ByCreatedAt) Less(i, j int) bool {
	return t.Tweets[i].CreatedAt > t.Tweets[j].CreatedAt
}

func main() {
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessSecret := os.Getenv("TWITTER_ACCESS_SECRET")

	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		log.Fatal("You should set consumer key/secret and access token/secret.")
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	var latestId int64 = 0
	for {
		tweets, _, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
			Count:   20,
			SinceID: latestId,
		})

		if err != nil {
			panic(err.Error())
		}

		sort.Sort(sort.Reverse(ByCreatedAt{tweets}))

		for _, t := range tweets {
			fmt.Println(t.Text)
			fmt.Println(t.ID)
			fmt.Println(t.CreatedAt)
			fmt.Println("----------------")
		}

		latestId = tweets[len(tweets)-1].ID
		time.Sleep(20 * time.Second)
	}
}
