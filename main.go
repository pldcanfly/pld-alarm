package main

import "github.com/pldcanfly/pld-alarm/server"

func main() {
	server := server.NewServer(":8080")
	server.Run()

}
