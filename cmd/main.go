package main

import (
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

func main() {
	nc, _ := nats.Connect("nats://demo.nats.io:4222")

	svc, err := micro.AddService(nc, micro.Config{
		Name:        "HashService",
		Version:     "1.0.0",
		Description: "Hashes input data into md5, sha1, sha256, sha512 or blake2b (256-bit). Output is Hexadecimal. Example: `nats req hash.md5 'hello'`.",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected.")

	hashGroup := svc.AddGroup("hash")

	handlers := map[string]micro.Handler{
		"md5":     micro.HandlerFunc(md5Handler),
		"sha1":    micro.HandlerFunc(sha1Handler),
		"sha256":  micro.HandlerFunc(sha256Handler),
		"sha512":  micro.HandlerFunc(sha512Handler),
		"blake2b": micro.HandlerFunc(blake2bHandler),
	}

	for e, h := range handlers {
		if err = hashGroup.AddEndpoint(e, h); err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Registered endpoints.")

	select {}
}
