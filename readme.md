**Go Cryptography Application**

Generate Key Pair:\
`go run main.go generate`

Encrypt File\
`go run main.go encrypt -f [path-to-file] -k [path-to-public-key]`

Decrypt File\
`go run main.go decrypt -f [path-to-ecnrypted-file] -k [path-to-private-key]`