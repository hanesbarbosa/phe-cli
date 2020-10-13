package main

import "fmt"

func enc(args []string) {
	flags := args[1:]
	// Check number of flags
	if len(flags) != 6 {
		// Abort with the name of the originating function
		abort(args[0])
	}
	if flags[0] == "-sk" && flags[1] == "./file" &&
		flags[2] == "-pk" && flags[3] == "./file" &&
		flags[4] == "-m" && checkIntegerFormat(flags[5]) {
		fmt.Println("Encrypt")
	} else {
		abort(args[0])
	}
}

func printEncMenu() {
	fmt.Println(SystemName + " -enc -sk <File> -pk <File> -m <Positive Integer>")
}
