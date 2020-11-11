package keygen

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func Generate(){
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	errorCheck(err)

	publickey:= &privatekey.PublicKey
	err = os.Mkdir("keypair", 0755)
	errorCheck(err)
    // private key dump
	var privateKeyBytes []byte = x509.MarshalPKCS1PrivateKey(privatekey)
	privateKeyBlock := &pem.Block{
		Type:    "RSA PRIVATE KEY",
		Bytes:   privateKeyBytes,
	}

	privatePem, err := os.Create("keypair/private.pem")
	errorCheck(err)

	err = pem.Encode(privatePem, privateKeyBlock)
	errorCheck(err)
	// public key dump
	var publicKeyBytes []byte = x509.MarshalPKCS1PublicKey(publickey)
	publicKeyBlock := &pem.Block{
		Type:    "PUBLIC KEY",
		Bytes:   publicKeyBytes,
	}

	publicPem, err := os.Create("keypair/public.pem")
	errorCheck(err)

	err = pem.Encode(publicPem, publicKeyBlock)
	errorCheck(err)

	fmt.Println("RAS Public and Private Keys Generated! (private.pem, public.pem)")
	os.Exit(0)
}

func errorCheck(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}