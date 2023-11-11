package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"log"

	"github.com/nats-io/nats.go/micro"
	"golang.org/x/crypto/blake2b"
)

func md5Handler(req micro.Request) {
	hash := md5.New()
	hash.Write(req.Data())
	resp := make([]byte, hex.EncodedLen(len(hash.Sum(nil))))
	hex.Encode(resp, hash.Sum(nil))
	req.Respond(resp)
	log.Printf("'%s' -[md5]-> %s", req.Data(), resp)
}

func sha1Handler(req micro.Request) {
	hash := sha1.New()
	hash.Write(req.Data())
	resp := make([]byte, hex.EncodedLen(len(hash.Sum(nil))))
	hex.Encode(resp, hash.Sum(nil))
	req.Respond(resp)
	log.Printf("'%s' -[sha1]-> %s", req.Data(), resp)
}

func sha256Handler(req micro.Request) {
	hash := sha256.New()
	hash.Write(req.Data())
	resp := make([]byte, hex.EncodedLen(len(hash.Sum(nil))))
	hex.Encode(resp, hash.Sum(nil))
	req.Respond(resp)
	log.Printf("'%s' -[sha256]-> %s", req.Data(), resp)
}

func sha512Handler(req micro.Request) {
	hash := sha512.New()
	hash.Write(req.Data())
	resp := make([]byte, hex.EncodedLen(len(hash.Sum(nil))))
	hex.Encode(resp, hash.Sum(nil))
	req.Respond(resp)
	log.Printf("'%s' -[sha512]-> %s", req.Data(), resp)
}

func blake2bHandler(req micro.Request) {
	hash, _ := blake2b.New256(nil)
	hash.Write(req.Data())
	resp := make([]byte, hex.EncodedLen(len(hash.Sum(nil))))
	hex.Encode(resp, hash.Sum(nil))
	req.Respond(resp)
	log.Printf("'%s' -[blake2b]-> %s", req.Data(), resp)
}
