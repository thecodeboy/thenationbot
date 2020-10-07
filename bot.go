package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

// Country conforms to the individual datatype in countries.json
type Country struct {
	ID          int    `json:"id"`
	Name        string `json:"country_name"`
	Code        string `json:"country_code"`
	Description string `json:"description"`
	Wiki        string `json:"wiki"`
}

func main() {
	log.Println("Starting Bot")

	// Reading array of countries as string
	file, err := ioutil.ReadFile("countries.json")
	if err != nil {
		log.Fatalln(err)
	}

	// Making array of Country structs
	countries := make([]Country, 0)
	json.Unmarshal([]byte(file), &countries)

	// Generating ID of country to tweet about
	rand.Seed(time.Now().UnixNano())
	chosenCountryID := rand.Intn(len(countries))
	tweet := prepareTweet(countries[chosenCountryID])
	log.Println(tweet)

	sendTweet(tweet)

	log.Println("Stopping Bot")
}

// Generating the tweet text
// TODO - use string.Builder instead
func prepareTweet(country Country) string {
	log.Println("Preparing tweet about " + country.Name)
	tweet := country.Name
	if len(country.Description) != 0 {
		tweet += " - " + country.Description
	}
	if len(country.Wiki) != 0 {
		tweet += " - " + country.Wiki
	}
	return tweet
}

// Sending prepared tweet
func sendTweet(tweet string) {
	log.Println("Preparing to send tweet")

	os.Getenv("FOO")

	// Bootstrapping API client
	api := anaconda.NewTwitterApiWithCredentials(os.Getenv("TOKEN"), os.Getenv("TOKENSECRET"),
		os.Getenv("CONSUMERKEY"), os.Getenv("CONSUMERSECRET"))

	_, err := api.PostTweet(tweet, nil)
	if err != nil {
		log.Fatalln("Error posting tweet: ", err)
	} else {
		log.Println("Tweet sent successfully")
	}
}
