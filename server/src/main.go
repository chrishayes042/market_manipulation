package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

// structure for credentials. Var name and datatype
type Credentials struct {
	ConsumerKey			string
	ConsumerSecret 		string
	AccessToken 		string
	AccessTokenSecret 	string
}

type TweetData struct {
	Handle			string
	Tweet 			string
	LikeCount		int
	RetweetCount	int
}

func main(){
	// using the pkg. if it doesn't load, throw error
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	// setting the creds variable to the Credential structure
	creds := Credentials {
		ConsumerKey: 		os.Getenv("CONSUMER_KEY"),
		ConsumerSecret: 	os.Getenv("CONSUMER_KEY_SECRET"),
		AccessToken: 		os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: 	os.Getenv("ACCESS_TOKEN_SECRET"),
	}
	client, err := getClient(&creds)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	searchTweet(client)
}

func getClient(creds *Credentials) (*twitter.Client, error){

	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)

	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	verify := &twitter.AccountVerifyParams{
		SkipStatus:	  twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	user, _, err := client.Accounts.VerifyCredentials(verify)
		if err != nil{
			fmt.Println(err)
			return nil, err
		}
	fmt.Println("Successfully authenticated using the following account : ", user.ScreenName)
	return client, nil
}

func searchTweet(client *twitter.Client) error {
	search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: "hello",
		Lang: "en",
	})

	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, v := range search.Statuses {
		tweet := TweetData{
			Handle: 		"@"+v.User.ScreenName,
			Tweet:			v.Text,
			LikeCount: 		v.FavoriteCount,
			RetweetCount: 	v.RetweetCount,
		}
		fmt.Printf("%+v\n\n ", tweet)
	}
	
	return nil
}