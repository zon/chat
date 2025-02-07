package main

const protocol string = "http"
const host string = "localhost"
const port string = "8080"
const idUrl string = "http://localhost:8081"

func rootUrl() string {
	return protocol +"://"+ host + ":" + port
}