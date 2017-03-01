package go_referrer_type

type ReferrerSearchYahoo struct {
	Param string
}

type ReferrerSearchGoogle struct {
	Param string
}

type ReferrerSocial struct {
	Param string
}

type ReferrerSocialLink struct {
	Param string
}

type ReferrerOganic struct {
	Param string
}

type ReferrerCp struct {
	Param string
}

type ReferrerDisplay struct {
	Param string
}

type ReferrerCpCampaign struct {
	Param string
}

type ReferrerAd struct {
	Param string
}

var ReferrerSearchYahoos = []ReferrerSearchYahoo{
	{
		Param:"yahoo",
	},
	{
		Param:"yahoo",
	},
}

var ReferrerSearchGoogles = []ReferrerSearchGoogle{
	{
		Param:"google",
	},
	{
		Param:"Google",
	},
}

var ReferrerSocialLinks = []ReferrerSocialLink{
	{
		Param:"twitter",
	},
	{
		Param:"Twitter",
	},
	{
		Param:"t.co",
	},
	{
		Param:"facebook",
	},
	{
		Param:"instagramk",
	},
	{
		Param:"Instagram",
	},
}

var ReferrerOganics = []ReferrerOganic{
	{
		Param:"organic",
	},
	{
		Param:"Organic",
	},
}

var ReferrerCps = []ReferrerCp{
	{
		Param:"cpc",
	},
	{
		Param:"ppc",
	},
}

var ReferrerDisplays = []ReferrerDisplay{
	{
		Param:"display",
	},
}

var ReferrerCpCampaigns = []ReferrerCpCampaign{
	{
		Param:"remarketing",
	},
}

var ReferrerAds = []ReferrerAd{
	{
		Param:"affiliate",
	},
	{
		Param:"email",
	},
}
