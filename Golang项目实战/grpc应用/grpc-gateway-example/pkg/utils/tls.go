package utils

import (
	"crypto/tls"
	"golang.org/x/net/http2"
	"io/ioutil"
	"log"
)

func GetTLSConfig(serverPemPath, serverKeyPath string) *tls.Config {
	var certKeyPair *tls.Certificate
	cert, _ := ioutil.ReadFile(serverPemPath)
	key, _ := ioutil.ReadFile(serverKeyPath)

	pair, err := tls.X509KeyPair(cert, key)
	if err != nil {
		log.Println("TLS KeyPair err: %v\n", err)
	}

	certKeyPair = &pair

	return &tls.Config{
		Certificates: []tls.Certificate{*certKeyPair},
		NextProtos:   []string{http2.NextProtoTLS},
	}
}
