package core

import "os"

const appUri string = "wurds"
const Protocol string = "http"
const hostDefault string = "localhost"
const Port string = "8080"
const authProtocol string = "http"
const authHostDefault string = "localhost"
const authPort string = "8081"
const AuthPath string = "/auth"

func Host() string {
	h := os.Getenv("WURDS_HOST")
	if h == "" {
		h = hostDefault
	}
	return h
}

func AuthHost() string {
	host := os.Getenv("WHO_THIS_HOST")
	if host == "" {
		host = authHostDefault
	}
	return host
}

func Url() string {
	return Protocol + "://" + Host() + ":" + Port
}

func AuthUrl() string {
	return authProtocol + "://" + AuthHost() + ":" + authPort + AuthPath + "/" + appUri
}
