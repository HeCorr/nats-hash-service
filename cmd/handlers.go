package main

import (
	"crypto"
	"encoding/hex"
	"log"

	"github.com/nats-io/nats.go/micro"
)

func sumHash(hash crypto.Hash, input []byte) []byte {
	h := hash.New()
	h.Write(input)
	output := make([]byte, hex.EncodedLen(len(h.Sum(nil))))
	hex.Encode(output, h.Sum(nil))
	if !quiet {
		log.Printf("'%s' -[%s]-> %s", input, hash, output)
	}
	return output
}

func md5Handler(req micro.Request) {
	req.Respond(sumHash(crypto.MD5, req.Data()))
}

func sha1Handler(req micro.Request) {
	req.Respond(sumHash(crypto.SHA1, req.Data()))
}

func sha256Handler(req micro.Request) {
	req.Respond(sumHash(crypto.SHA256, req.Data()))
}

func sha512Handler(req micro.Request) {
	req.Respond(sumHash(crypto.SHA512, req.Data()))
}

func blake2bHandler(req micro.Request) {
	req.Respond(sumHash(crypto.BLAKE2b_256, req.Data()))
}
