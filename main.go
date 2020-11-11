package main

import (
	"./decrypt"
	"./encrypt"
	"./keygen"
	"os"
)

func main() {
	args := os.Args
	readArgs(args)
}

func readArgs(args []string) {
	if len(args) <= 1 {
		os.Exit(1)
	}

	switch args[1] {
	case "encrypt":
		encrypt.Encode(args[1:])
		break
	case "decrypt":
		decrypt.Decode(args[1:])
		break
	case "generate":
		keygen.Generate()
		break
	default:
		os.Exit(1)
	}
}
