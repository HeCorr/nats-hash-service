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
		panic(err)
	}

	hashGroup := svc.AddGroup("hash")

	if err = hashGroup.AddEndpoint("md5", micro.HandlerFunc(md5Handler)); err != nil {
		panic(err)
	}

	if err = hashGroup.AddEndpoint("sha1", micro.HandlerFunc(sha1Handler)); err != nil {
		panic(err)
	}

	if err = hashGroup.AddEndpoint("sha256", micro.HandlerFunc(sha256Handler)); err != nil {
		panic(err)
	}

	if err = hashGroup.AddEndpoint("sha512", micro.HandlerFunc(sha512Handler)); err != nil {
		panic(err)
	}

	if err = hashGroup.AddEndpoint("blake2b", micro.HandlerFunc(blake2bHandler)); err != nil {
		panic(err)
	}

	log.Println("Connected.")

	select {}
}
