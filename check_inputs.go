package main

import (
	"os"
)

func checkInputs(args []string) {
	// TODO: check the case where nothing is provided.
	switch args[0] {
	case "-keygen":
		// Generate key pair
		keyGen(args[1:])
	// case "-enc":
	// 	// Encrypt
	// 	enc(flags)
	// case "-dec":
	// 	// Decrypt
	// 	dec(flags)
	// case "-tokgen":
	// 	// Generate token
	// 	tokgen(flags)
	default:
		// printGeneralMenu()
		// exitProgram()
	}
}

func printGeneralMenu() {
	printKeyGenSignature()
	printEncSignature()
	printDecSignature()
	printTokGenSignature()
}

func exitProgram() {
	os.Exit(1)
}
