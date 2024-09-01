package main

import (
	"flag"
	"go-network-programming/ch06/tftp"
	"log"
	"os"
)

var (
	address = flag.String("a", "127.0.0.1:8081", "listen address")
	payload = flag.String("p", "payload", "file to serve to clients")
)

func main() {
	flag.Parse()
	p, err := os.ReadFile(*payload)
	if err != nil {
		log.Fatal(err)
	}

	s := tftp.Server{Payload: p}
	log.Fatal(s.ListenAndServe(*address))
}
