package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {

	// If there are not enough arguments
	if len(os.Args) < 2 {

		// TODO: In the future, implement REPL mode
		fmt.Println("Usage: sha256 \"some string\"")

	} else {

		// Take the first argument and hash it
		bytes := sha256.Sum256([]byte(os.Args[1]))
		hashHex := hex.EncodeToString(bytes[:])

		// Print the hash in hexadecimal
		fmt.Println(hashHex)

	}

}
