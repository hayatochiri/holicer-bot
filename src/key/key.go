package key

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
)

func GetTwitterApi(consumer_key, consumer_secret, access_token, access_secret string) *anaconda.TwitterApi {
	anaconda.SetConsumerKey(consumer_key)
	anaconda.SetConsumerSecret(consumer_secret)
	api := anaconda.NewTwitterApi(access_token, access_secret)
	return api
}
