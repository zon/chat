package core

const ProxyPort string = "7331"

func ProxyUrl() string {
	return Protocol + "://" + Host() + ":" + ProxyPort
}