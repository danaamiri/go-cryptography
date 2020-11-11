package decrypt

import (
	"../utils"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

func Decode(args []string) {
	filename, privateFile, err := readArgs(args)
	errorCheck(err)

	dataFile := readFile(filename)
	privateKeyFile := readFile(privateFile)

	privateBlock, _ := pem.Decode(privateKeyFile)
	privateKey, err := x509.ParsePKCS1PrivateKey(privateBlock.Bytes)
	errorCheck(err)

	//privateBlock, _ := pem.Decode(privateKeyFile)
	//privateKey, err := x509.ParsePKCS1PrivateKey(privateBlock.Bytes)
	//errorCheck(err)
	err = os.RemoveAll("decrypted")
	errorCheck(err)

	err = os.Mkdir("decrypted", 0775)
	errorCheck(err)

	j := len(dataFile)
	chunk := 256
	var temp []byte

	n := math.Floor(float64(j / chunk))
	pp := math.Floor(n / 100)
	p := float64(0)
	percent := 1
	utils.ShowLoading("Decrypting file "+filename, 0)
	var f *os.File
	for i := 0; i < j; i += chunk {
		temp = dataFile[i : i+chunk]
		decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, temp)
		errorCheck(err)
		//err = ioutil.WriteFile("encrypted/data.p", append(encrypted, (byte) "\n"), 0644)
		if i == 0 {
			if string(decrypted) != "" {
				f, err = os.Create("decrypted/data." + string(decrypted))
				errorCheck(err)
			} else {
				f, err = os.Create("decrypted/data")
				errorCheck(err)
			}

		} else {
			_, err = f.Write(decrypted)
			errorCheck(err)
		}

		if p == pp {

			utils.ShowLoading("Decrypting file "+filename, percent)
			percent++
			p = 0
		} else {
			p++
		}

	}
}

func readArgs(args []string) (filename string, privatekey string, err error) {
	if len(args) < 5 {
		err := fmt.Errorf("at least 2 args and their values are nessacary. \n usage: \n\t -f [filename] \n\t -p [publickey] \n\t -k [privatekey]")
		return "", "", err
	}

	var fn string
	var pri string
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
					pri = args[f+1]
				}
			}
		}
	}

	if pri == "" || fn == "" {
		err := fmt.Errorf("at least 2 args and their values are nessacary. \n usage: \n\t -f [filename] \n\t -p [publickey] \n\t -k [privatekey]")
		return "", "", err
	}

	return fn, pri, nil
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
