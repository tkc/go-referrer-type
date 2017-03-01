package go_referrer_type

import (
	"errors"
	"strings"
	"net/url"
)

const paramTypeNone = 0

type ReferrerType struct {
	ReferrerHost             string
	ReferrerParam            string
	ReferrerParamUtmCampaign string
	ReferrerParamUtmSource   string
	ReferrerParamUtmMedium   string
	ReferrerPath             string
	ServerHost               string
	ResultType               ResultType
}

func Create(referrerUrl string, serverUrl string) (ReferrerType, error) {

	referrerType := ReferrerType{
		ResultType:CreateResultType(),
	}
	err := referrerType.SetParam(referrerUrl, serverUrl)
	if err != nil {
		return ReferrerType{}, errors.New("arge is not url .")
	}
	return referrerType, nil
}

func (r *ReferrerType) SetResultType(resultType ResultType) *ReferrerType {
	r.ResultType = resultType
	return r
}

func (r *ReferrerType) SetParam(referrerUrl string, hostUrl string) error {

	ReParse, err := url.Parse(referrerUrl)
	if err != nil {
		return err
	}

	r.ReferrerHost = ReParse.Host
	r.ReferrerPath = ReParse.Path
	r.ReferrerParam = ReParse.RawQuery

	hostParse, err := url.Parse(hostUrl)
	if err != nil {
		return err
	}

	r.ServerHost = hostParse.Host
	return nil
}

func (r *ReferrerType) GetType() int {

	if r.ReferrerHost == "" {
		return r.ResultType.TYPE_REFERRER_NONE
	}

	if len(r.ReferrerParam) > 0 {
		paramType := r.GetParam()
		if paramType != paramTypeNone {
			return paramType
		}
	}

	if r.ReferrerHost == r.ServerHost {
		return r.ResultType.TYPE_REFERRER_LINK_INNER
	}

	if strings.Contains(r.ReferrerHost, "yahoo.") {
		return r.ResultType.TYPE_REFERRER_ORGANIC_SEARCH_YAHOO
	}

	if strings.Contains(r.ReferrerHost, "google.") {
		return r.ResultType.TYPE_REFERRER_ORGANIC_SEARCH_GOOGLE
	}

	if strings.Contains(r.ReferrerHost, "bing.") {
		return r.ResultType.TYPE_REFERRER_ORGANIC_SEARCH_ETC
	}

	if strings.Contains(r.ReferrerHost, "facebook.") {
		return r.ResultType.TYPE_REFERRER_SOCIAL_FACEBOOK
	}

	if strings.Contains(r.ReferrerHost, "instagram.") {
		return r.ResultType.TYPE_REFERRER_SOCIAL_INSTAGRAM
	}

	if strings.Contains(r.ReferrerHost, "twitter.") {
		return r.ResultType.TYPE_REFERRER_SOCIAL_TWITTER
	}

	if r.ReferrerHost != r.ServerHost {
		return r.ResultType.TYPE_REFERRER_LINK_OUTER
	}

	return r.ResultType.TYPE_REFERRER_NONE
}

func (r *ReferrerType) GetParam() int {

	r.ParseQuery()

	for _, v := range ReferrerDisplays {
		if strings.Contains(r.ReferrerParamUtmCampaign, v.Param) {
			for _, v2 := range ReferrerSearchYahoos {
				if strings.Contains(r.ReferrerParamUtmSource, v2.Param) {
					return r.ResultType.TYPE_REFERRER_AD_YAHOO_DISPLAY
				}
			}
			for _, v2 := range ReferrerSearchGoogles {
				if strings.Contains(r.ReferrerParamUtmSource, v2.Param) {
					return r.ResultType.TYPE_REFERRER_AD_GOOGLE_DISPLAY
				}
			}

		}
	}

	for _, v := range ReferrerCps {
		if strings.Contains(r.ReferrerParam, v.Param) {
			for _, v2 := range ReferrerSearchYahoos {
				if strings.Contains(r.ReferrerParam, v2.Param) {
					return r.ResultType.TYPE_REFERRER_AD_YAHOO_CPC
				}
			}
			for _, v2 := range ReferrerSearchGoogles {
				if strings.Contains(r.ReferrerParam, v2.Param) {
					return r.ResultType.TYPE_REFERRER_AD_GOOGLE_CPC
				}
			}
		}
	}

	for _, v := range ReferrerCpCampaigns {
		if strings.Contains(r.ReferrerParam, v.Param) {
			for _, v2 := range ReferrerSearchYahoos {
				if strings.Contains(r.ReferrerParam, v2.Param) {
					return r.ResultType.TYPE_REFERRER_AD_YAHOO_REMARKETING
				}
			}
			for _, v2 := range ReferrerSearchGoogles {
				if strings.Contains(r.ReferrerParam, v2.Param) {
					return r.ResultType.TYPE_REFERRER_AD_GOOGLE_REMARKETING
				}
			}
		}
	}

	for _, v := range ReferrerSocialLinks {
		if strings.Contains(r.ReferrerParam, v.Param) {
			return r.ResultType.TYPE_REFERRER_SOCIAL_ETC
		}
	}

	for _, v := range ReferrerAds {
		if strings.Contains(r.ReferrerParamUtmMedium, v.Param) {
			return r.ResultType.TYPE_REFERRER_LINK_OUTER
		}
	}

	for _, v := range ReferrerOganics {
		if strings.Contains(r.ReferrerParamUtmMedium, v.Param) {
			return r.ResultType.TYPE_REFERRER_LINK_OUTER
		}
	}

	return paramTypeNone
}

func (r *ReferrerType) ParseQuery() {

	param, err := url.ParseQuery(r.ReferrerParam)
	if (err == nil) {
		if len(param["utm_campaign"]) > 0 {
			r.ReferrerParamUtmCampaign = param["utm_campaign"][0]
		}
		if len(param["utm_source"]) > 0 {
			r.ReferrerParamUtmSource = param["utm_source"][0]
		}
		if len(param["utm_medium"]) > 0 {
			r.ReferrerParamUtmMedium = param["utm_medium"][0]
		}
	}
}
