package main

import "github.com/payfazz/fazzlearning-api/http/server"

func main() {
	api := server.CreateAPIServer()
	api.Serve()
}
