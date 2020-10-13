package main

import (
	"os"
)

// SystemName defines the CLI command name for messages.
const SystemName = "phe-cli"

func main() {

	if len(os.Args) <= 3 {
		abort("main")
	}

	// Flags comming from the user
	checkInputs(os.Args[1:])
}

func checkInputs(args []string) {
	// TODO: check the case where nothing is provided.
	switch args[0] {
	case "-keygen":
		// Generate key pair
		keyGen(args)
	case "-enc":
		// Encrypt
		enc(args)
	case "-dec":
		// Decrypt
		dec(args)
	case "-tokgen":
		// Generate token
		tokgen(args)
	default:
		abort("main")
	}
}

func printMenu(function string) {
	switch function {
	case "-keygen":
		printKeyGenMenu()
	case "-enc":
		printEncMenu()
	case "-dec":
		printDecMenu()
	case "-tokgen":
		printTokGenMenu()
	default:
		printKeyGenMenu()
		printEncMenu()
		printDecMenu()
		printTokGenMenu()
	}
}

func abort(function string) {
	printMenu(function)
	os.Exit(1)
}
