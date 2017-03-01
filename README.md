# go-referrer-type
determine of url referrer type.

## Usage

```go
var referrerUrl = "https://www.google.co.jp/"
var serverUrl = "http://example.com/"
r, _ := go_referrer_type.Create(referrerUrl, serverUrl)
res := r.GetType()
=>TYPE_REFERRER_ORGANIC_SEARCH_GOOGLE
```

##  Default Result Types

```go
const TYPE_REFERRER_ORGANIC_SEARCH_ETC = 100; // organic etc
const TYPE_REFERRER_ORGANIC_SEARCH_YAHOO = 101; // organic yahoo
const TYPE_REFERRER_ORGANIC_SEARCH_GOOGLE = 102; // organic google
const TYPE_REFERRER_LINK_OUTER = 201; // direct link outer
const TYPE_REFERRER_LINK_INNER = 202; // direct link inner
const TYPE_REFERRER_SOCIAL_ETC = 300; // social outer
const TYPE_REFERRER_SOCIAL_TWITTER = 301; // social titter
const TYPE_REFERRER_SOCIAL_FACEBOOK = 302; // social facebook
const TYPE_REFERRER_SOCIAL_INSTAGRAM = 303; // social instagram
const TYPE_REFERRER_AD_YAHOO_CPC = 401; // adWords yahoo cpc
const TYPE_REFERRER_AD_GOOGLE_CPC = 402; // adWords google cpc
const TYPE_REFERRER_AD_YAHOO_DISPLAY = 403; // adWords yahoo display
const TYPE_REFERRER_AD_GOOGLE_DISPLAY = 404; // adWords google display
const TYPE_REFERRER_AD_YAHOO_REMARKETING = 405; // adWords yahoo remarketing
const TYPE_REFERRER_AD_GOOGLE_REMARKETING = 406; // adWords google remarketing
const TYPE_REFERRER_NONE = 999; // referrer none
```