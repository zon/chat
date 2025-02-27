package main

import "os"

const appUri string = "chat"
const protocol string = "http"
const hostDefault string = "localhost"
const port string = "8080"
const authProtocol string = "http"
const authHostDefault string = "localhost"
const authPort string = "8081"

func host() string {
	h := os.Getenv("WURDS_HOST")
	if h == "" {
		h = hostDefault
	}
	return h
}

func authHost() string {
	host := os.Getenv("WHO_THIS_HOST")
	if host == "" {
		host = authHostDefault
	}
	return host
}

func url() string {
	return protocol +"://"+ host() +":"+ port
}

func authUrl() string {
	return authProtocol + "://" + authHost() + ":" + authPort + authPath + "/" + appUri
}
