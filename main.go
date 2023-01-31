package main

import "dns-go/server"

func main() {
	s := server.New()
	s.Start()
	defer s.Close()
}
