package main

import "fmt"

func tokgen(args []string) {
	flags := args[1:]
	// Check number of flags
	if len(flags) != 8 {
		// Abort with the name of the originating function
		abort(args[0])
	}
	if flags[0] == "-sk-old" && flags[1] == "./file" &&
		flags[2] == "-sk-new" && flags[3] == "./file" &&
		flags[4] == "-pk-old" && flags[5] == "./file" &&
		flags[6] == "-pk-new" && flags[7] == "./file" {
		fmt.Println("Token Generator")
	} else {
		abort(args[0])
	}
}

func printTokGenMenu() {
	fmt.Println(SystemName + " tokgen -sk-old <File> -sk-new <File> -pk-old <File> -pk-new <File>")
}
