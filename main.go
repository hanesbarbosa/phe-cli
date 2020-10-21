package main

import (
	"os"
)

// SystemName defines the CLI command name for messages.
const SystemName = "phe-cli"

func main() {

	if len(os.Args) < 2 {
		abort("main")
	}

	// Flags comming from the user
	checkInputs(os.Args[1:])
}

func checkInputs(args []string) {
	// TODO: check the case where nothing is provided.
	switch args[0] {
	case "--keygen":
		// Generate key pair
		keyGen(args)
	case "--encrypt":
		// Encrypt
		enc(args)
	case "--decrypt":
		// Decrypt
		dec(args)
	case "--tokgen":
		// Generate token
		tokGen(args)
	case "--listkeys":
		// List all keys
		listKeys(args)
	default:
		abort("main")
	}
}

func printMenu(function string) {
	switch function {
	case "--keygen":
		printKeyGenMenu()
	case "--encrypt":
		printEncMenu()
	case "--decrypt":
		printDecMenu()
	case "--tokgen":
		printTokGenMenu()
	case "--listkeys":
		printListKeysMenu()
	case "main":
		printMainMenu()
	}
}

func abort(function string) {
	printMenu(function)
	os.Exit(1)
}

func printMainMenu() {
	printKeyGenMenu()
	printEncMenu()
	printDecMenu()
	printTokGenMenu()
	printListKeysMenu()
}
