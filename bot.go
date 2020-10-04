package main

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	fmt.Println("Starting bot...")
	s := `Montenegro - a small and placid country with beautific scenic hills 
	#country #world #nation #learn
	https://en.wikipedia.org/wiki/Montenegro`
	tweetCurrent(s)
}

func tweetCurrent(o string) {
	fmt.Println("Preparing to tweet...")
	// api := anaconda.NewTwitterApiWithCredentials(config.Token, config.TokenSecret, config.ConsumerKey, config.ConsumerSecret)
	api := anaconda.NewTwitterApiWithCredentials("1311355556448186368-ypEFB6gZRmW7SL2oQfEdgP0br0Od1N",
		"oEGamPqmK5eN81NNMnCt0EhaU1I1xCwvgIdHfI9024BHZ",
		"bkDXjE0qF5QuhnfZ5S0jxy1MQ",
		"nLmey0RtnfhwzDpKRrElSsW0YjY5f0wulWRqssmmOB0phDp63B")
	tweet, err := api.PostTweet(o, nil)
	if err != nil {
		fmt.Println("update error:", err)
	} else {
		fmt.Println("Tweet posted:")
		fmt.Println(tweet.Text)
	}
}
