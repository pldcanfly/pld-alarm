package main

import "github.com/pldcanfly/pld-alarm/server"

func main() {
	s  := server.NewServer(":8080")
	s.Run()

}
