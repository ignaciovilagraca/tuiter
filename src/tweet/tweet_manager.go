package tweetManager

import (
	"fmt"
	"github.com/ignaciovila/tuiter/src/user"
	)

var tweets []Tweet
var tweetsByUser map[string] []Tweet

func PublishTweet(t Tweet) (int, error) {
	if len(t.GetUser()) < 1 {
		return -1, fmt.Errorf("user is required")
	}

	if len(t.GetText()) < 1 {
		return -1, fmt.Errorf("text is required")
	}

	if len(t.GetText()) > 140 {
		return -1, fmt.Errorf("max length is 140")
	}

	if !userManager.ExistsUser(t.GetUser()) {
		return -1, fmt.Errorf("invalid user")
	}

	tweets = append(tweets, t)

	currentList := tweetsByUser[t.GetUser()]
	tweetsByUser[t.GetUser()] = append(currentList, t)

	return len(tweets) -1, nil
}

func GetTweetById(index int) Tweet {
	return tweets[index]
}

func CountTweetsByUser(usr string) int {
	count := 0

	for i := 0; i < len(tweets); i++ {
		if tweets[i].GetUser() == usr {
			count++
		}
	}

	return count
}

func GetTweetByUser(usr string) []Tweet {
	return tweetsByUser[usr]
}

func GetTweets() []Tweet {
	return tweets
}

func NewTweetManager() {
	tweets = make([]Tweet, 0)
	tweetsByUser = make(map[string] []Tweet)
}
