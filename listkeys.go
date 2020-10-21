package main

import (
	"fmt"
)

// listKeys shows all the stored keys.
func listKeys(args []string) {
	flags := args[1:]
	// Check number of flags
	if len(flags) == 2 && flags[0] == "-kr" && fileExists(flags[1]) {
		kr := readKeyRing(flags[1])
		kr.List()
	} else {
		// Abort giving the name of the originating function
		abort(args[0])
	}
}

func printListKeysMenu() {
	fmt.Println(SystemName + " --listkeys -kr <key ring file>")
}
