package core

import "os"

const appUri string = "wurbs"
const AuthPath string = "/auth"

var port string = "8080"
var url string = "http://localhost:8080"

func InitConfig() {
	p := os.Getenv("PORT")
	if p != "" {
		port = p
	}
	u := os.Getenv("WURBS_URL")
	if u != "" {
		url = u
	}
}

func Port() string {
	return port
}

func Url() string {
	return url
}
