package main

const appUri string = "chat"
const protocol string = "http"
const host string = "localhost"
const port string = "8080"
const authProtocol string = "http"
const authHost string = "localhost"
const authPort string = "8081"

func authUrl() string {
	return authProtocol + "://" + authHost + ":" + authPort + authPath + "/" + appUri
}
