package core

import nurl "net/url"

const ProxyPort string = "7331"

func ProxyUrl() string {
	u, err := nurl.Parse(Url())
	if err != nil {
		panic(err)
	}
	return u.Scheme + "://" + u.Hostname() + ":" + ProxyPort
}
