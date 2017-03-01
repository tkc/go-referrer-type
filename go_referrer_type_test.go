package go_referrer_type

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

const exampleHost = "http://example.com/"

func TestReferrerTyp_GetType0(t *testing.T) {
	referrerUrl := ""
	serverUrl := exampleHost
	r, err := Create(referrerUrl, serverUrl)

	if err == nil {
		res := r.GetType()
		assert.Equal(t, TYPE_REFERRER_NONE, res)
	} else {
		assert.Fail(t, "error occer")
	}
}

func TestReferrerTyp_GetType2(t *testing.T) {
	referrerUrl := exampleHost
	serverUrl := exampleHost
	r, err := Create(referrerUrl, serverUrl)

	if err == nil {
		res := r.GetType()
		assert.Equal(t, TYPE_REFERRER_LINK_INNER, res)
	} else {
		assert.Fail(t, "error occer")
	}
}

func TestReferrerTyp_GetType3(t *testing.T) {
	referrerUrl := "http://example.com/test?name=1$address=japan"
	serverUrl := exampleHost
	r, err := Create(referrerUrl, serverUrl)

	if err == nil {
		res := r.GetType()
		assert.Equal(t, TYPE_REFERRER_LINK_INNER, res)
	} else {
		assert.Fail(t, "error occer")
	}
}

func TestReferrerTyp_GetType4(t *testing.T) {
	referrerUrl := "http://www.yahoo.co.jp"
	serverUrl := exampleHost
	r, err := Create(referrerUrl, serverUrl)

	if err == nil {
		res := r.GetType()
		assert.Equal(t, TYPE_REFERRER_ORGANIC_SEARCH_YAHOO, res)
	} else {
		assert.Fail(t, "error occer")
	}
}

func TestReferrerTyp_GetType5(t *testing.T) {
	referrerUrl := "https://www.google.co.jp/"
	serverUrl := exampleHost
	r, err := Create(referrerUrl, serverUrl)

	if err == nil {
		res := r.GetType()
		assert.Equal(t, TYPE_REFERRER_ORGANIC_SEARCH_GOOGLE, res)
	} else {
		assert.Fail(t, "error occer")
	}
}

func TestReferrerTyp_GetType6(t *testing.T) {
	referrerUrl := "https://www.bing.com/"
	serverUrl := exampleHost
	r, err := Create(referrerUrl, serverUrl)

	if err == nil {
		res := r.GetType()
		assert.Equal(t, TYPE_REFERRER_ORGANIC_SEARCH_ETC, res)
	} else {
		assert.Fail(t, "error occer")
	}
}

func TestReferrerTyp_GetType7(t *testing.T) {
	referrerUrl := "https://www.facebook.com/"
	serverUrl := exampleHost
	r, err := Create(referrerUrl, serverUrl)

	if err == nil {
		res := r.GetType()
		assert.Equal(t, TYPE_REFERRER_SOCIAL_FACEBOOK, res)
	} else {
		assert.Fail(t, "error occer")
	}
}

func TestReferrerTyp_GetType8(t *testing.T) {
	referrerUrl := "https://www.instagram.com/lasselom/"
	serverUrl := exampleHost
	r, err := Create(referrerUrl, serverUrl)

	if err == nil {
		res := r.GetType()
		assert.Equal(t, TYPE_REFERRER_SOCIAL_INSTAGRAM, res)
	} else {
		assert.Fail(t, "error occer")
	}
}

func TestReferrerTyp_GetType9(t *testing.T) {
	referrerUrl := "https://twitter.com/"
	serverUrl := exampleHost
	r, err := Create(referrerUrl, serverUrl)

	if err == nil {
		res := r.GetType()
		assert.Equal(t, TYPE_REFERRER_SOCIAL_TWITTER, res)
	} else {
		assert.Fail(t, "error occer")
	}
}

func TestReferrerTyp_GetType10(t *testing.T) {
	referrerUrl := "http://www.example_outer.com/"
	serverUrl := exampleHost
	r, err := Create(referrerUrl, serverUrl)

	if err == nil {
		res := r.GetType()
		assert.Equal(t, TYPE_REFERRER_LINK_OUTER, res)
	} else {
		assert.Fail(t, "error occer")
	}
}

func TestReferrerTyp_GetType11(t *testing.T) {
	referrerUrl := "http://www.example.com/download.html?utm_source=yahoo&utm_medium=cpc&utm_term=%E3%81%A7%E3%82%82&utm_content=%E5%BA%83%E5%91%8A%E3%81%AE%E5%8C%BA%E5%88%A5%E7%94%A8&utm_campaign=%E3%83%97%E3%83%AD%E3%83%A2%E3%83%BC%E3%82%B7%E3%83%A7%E3%83%B3%20%E3%82%B3%E3%83%BC%E3%83%89"
	serverUrl := exampleHost
	r, err := Create(referrerUrl, serverUrl)

	if err == nil {
		res := r.GetType()
		assert.Equal(t, TYPE_REFERRER_AD_YAHOO_CPC, res)
	} else {
		assert.Fail(t, "error occer")
	}
}

func TestReferrerTyp_GetType12(t *testing.T) {
	referrerUrl := "http://www.example_outer.com/download.html?utm_source=yahoo&utm_medium=ppc&utm_term=%E3%81%A7%E3%82%82&utm_content=%E5%BA%83%E5%91%8A%E3%81%AE%E5%8C%BA%E5%88%A5%E7%94%A8&utm_campaign=%E3%83%97%E3%83%AD%E3%83%A2%E3%83%BC%E3%82%B7%E3%83%A7%E3%83%B3%20%E3%82%B3%E3%83%BC%E3%83%89"
	serverUrl := exampleHost
	r, err := Create(referrerUrl, serverUrl)

	if err == nil {
		res := r.GetType()
		assert.Equal(t, TYPE_REFERRER_AD_YAHOO_CPC, res)
	} else {
		assert.Fail(t, "error occer")
	}
}

func TestReferrerTyp_GetType13(t *testing.T) {
	referrerUrl := "http://www.example_outer.com/download.html?utm_source=Google&utm_medium=cpc&utm_term=%E3%81%A7%E3%82%82&utm_content=%E5%BA%83%E5%91%8A%E3%81%AE%E5%8C%BA%E5%88%A5%E7%94%A8&utm_campaign=%E3%83%97%E3%83%AD%E3%83%A2%E3%83%BC%E3%82%B7%E3%83%A7%E3%83%B3%20%E3%82%B3%E3%83%BC%E3%83%89"
	serverUrl := exampleHost
	r, err := Create(referrerUrl, serverUrl)

	if err == nil {
		res := r.GetType()
		assert.Equal(t, TYPE_REFERRER_AD_GOOGLE_CPC, res)
	} else {
		assert.Fail(t, "error occer")
	}
}

func TestReferrerTyp_GetType14(t *testing.T) {
	referrerUrl := "http://www.example_outer.com/download.html?utm_source=yahoo&utm_medium=display&utm_campaign=display"
	serverUrl := exampleHost
	r, err := Create(referrerUrl, serverUrl)

	if err == nil {
		res := r.GetType()
		assert.Equal(t, TYPE_REFERRER_AD_YAHOO_DISPLAY, res)
	} else {
		assert.Fail(t, "error occer")
	}
}

func TestReferrerTyp_GetType15(t *testing.T) {
	referrerUrl := "http://www.example_outer.com/download.html?utm_source=Google&utm_medium=display&utm_campaign=display"
	serverUrl := exampleHost
	r, err := Create(referrerUrl, serverUrl)
	if err == nil {
		res := r.GetType()
		assert.Equal(t, TYPE_REFERRER_AD_GOOGLE_DISPLAY, res)
	} else {
		assert.Fail(t, "error occer")
	}
}

func TestReferrerTyp_GetType16(t *testing.T) {
	referrerUrl := "http://www.example_outer.com/download.html?utm_source=Google&utm_medium=display&utm_campaign=remarketing"
	serverUrl := exampleHost
	r, err := Create(referrerUrl, serverUrl)
	if err == nil {
		res := r.GetType()
		assert.Equal(t, TYPE_REFERRER_AD_GOOGLE_REMARKETING, res)
	} else {
		assert.Fail(t, "error occer")
	}
}

