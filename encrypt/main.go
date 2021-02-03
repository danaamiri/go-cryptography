package encrypt

import (
	"../utils"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Encode(args []string) {
	filename, pubFile, err := readArgs(args)
	errorCheck(err)

	orgName, fileType := getFileName(filename)

	dataFile := readFile(filename)
	publicKeyFile := readFile(pubFile)

	publicBlock, _ := pem.Decode(publicKeyFile)
	publicKey, err := x509.ParsePKCS1PublicKey(publicBlock.Bytes)
	errorCheck(err)

	//privateBlock, _ := pem.Decode(privateKeyFile)
	//privateKey, err := x509.ParsePKCS1PrivateKey(privateBlock.Bytes)
	//errorCheck(err)
	err = os.RemoveAll("encrypted")
	errorCheck(err)

	err = os.Mkdir("encrypted", 0775)
	errorCheck(err)

	f, err := os.Create("encrypted/data.p")
	errorCheck(err)

	j := len(dataFile)
	chunk := 214
	var temp []byte

	finalFileName := orgName
	if len(fileType) > 0 {
		finalFileName += "." + fileType
	}

	encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(finalFileName))
	errorCheck(err)
	_, err = f.Write(encrypted)

	n := j / chunk
	pp := n / 100
	p := 0
	percent := 1
	utils.ShowLoading("Encrypting file "+filename, 0)
	for i := 0; i < j; i += chunk {
		temp = dataFile[i : i+chunk]
		encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, temp)
		errorCheck(err)
		_, err = f.Write(encrypted)
		errorCheck(err)
		if p == pp {

			utils.ShowLoading("Encrypting file "+filename, percent)
			percent++
			p = 0
		} else {
			p++
		}

	}

}

func readArgs(args []string) (filename string, publickey string, err error) {
	if len(args) < 5 {
		err := fmt.Errorf("at least 2 args and their values are nessacary. \n usage: \n\t -f [filename] \n\t -p [publickey] \n\t -k [privatekey]")
		return "", "", err
	}

	var fn string
	var pub string
	for f := range args {
		if args[f] == "-f" {
			if f+1 < len(args) {
				if args[f+1] != "-k" {
					fn = args[f+1]
				}
			}
		}
		if args[f] == "-k" {
			if f+1 < len(args) {
				if args[f+1] != "-f" {
					pub = args[f+1]
				}
			}
		}
	}

	if pub == "" || fn == "" {
		err := fmt.Errorf("at least 2 args and their values are nessacary. \n usage: \n\t -f [filename] \n\t -p [publickey] \n\t -k [privatekey]")
		return "", "", err
	}

	return fn, pub, nil
}

func errorCheck(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

func readFile(filename string) (file []byte) {
	dat, err := ioutil.ReadFile(filename)
	errorCheck(err)
	return dat
}

func getFileName(fileCompleteName string) (filename string, fileType string) {
	fileOrgNameSpl := strings.Split(fileCompleteName, "/")
	fileOrgName := fileOrgNameSpl[len(fileOrgNameSpl)-1]
	splitted := strings.Split(fileOrgName, ".")
	if len(splitted) > 1 {
		println(splitted[0])
		return splitted[0], splitted[len(splitted)-1]
	} else {
		return splitted[0], ""
	}
}
