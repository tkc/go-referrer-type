package main

import (
	"github.com/tkc/go-referrer-type"
)

func main() {

	var referrerUrl = "https://www.google.co.jp/"
	var serverUrl = "http://example.com/"

	r, err := go_referrer_type.Create(referrerUrl, serverUrl)

	if err == nil {
		res := r.GetType()
		if res == go_referrer_type.TYPE_REFERRER_ORGANIC_SEARCH_GOOGLE {
			log.Print("TYPE_REFERRER_ORGANIC_SEARCH_GOOGLE")
		}
	}
}
