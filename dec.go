package main

import "fmt"

func dec(args []string) {
	flags := args[1:]
	// Check number of flags
	if len(flags) != 6 {
		// Abort with the name of the originating function
		abort(args[0])
	}
	if flags[0] == "-sk" && flags[1] == "./file" &&
		flags[2] == "-pk" && flags[3] == "./file" &&
		flags[4] == "-c" && flags[5] == "Multivector" {
		fmt.Println("Decrypt")
	} else {
		abort(args[0])
	}
}

func printDecMenu() {
	fmt.Println(SystemName + " -dec -sk <File> -pk <File> -c <String Multivector>")
}
