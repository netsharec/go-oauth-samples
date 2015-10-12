package main

import (
	"fmt"
	"github.com/mrjones/oauth"
	"io/ioutil"
	"log"
	"os"
)

var (
	ConsumerKey       = os.Getenv("ENV_TWITTER_CONSUMER_KEY")
	ConsumerSecret    = os.Getenv("ENV_TWITTER_CONSUMER_SECRET")
	AccessToken       = os.Getenv("ENV_TWITTER_ACCESS_TOKEN")
	AccessTokenSecret = os.Getenv("ENV_TWITTER_ACCESS_TOKEN_SECRET")
)

func main() {
	consumer := oauth.NewConsumer(ConsumerKey,
		ConsumerSecret,
		oauth.ServiceProvider{})
	//NOTE: remove this line or turn off Debug if you don't
	//want to see what the headers look like
	consumer.Debug(true)
	//Roll your own AccessToken struct
	accessToken := &oauth.AccessToken{Token: AccessToken,
		Secret: AccessTokenSecret}
	twitterEndPoint := "https://api.twitter.com/1.1/statuses/mentions_timeline.json"
	response, err := consumer.Get(twitterEndPoint, nil, accessToken)
	if err != nil {
		log.Fatal(err, response)
	}
	defer response.Body.Close()
	fmt.Println("Response:", response.StatusCode, response.Status)
	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(respBody))
}
