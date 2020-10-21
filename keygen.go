package main

import (
	"fmt"
	"strconv"

	"github.com/hanesbarbosa/phe"
)

// keyGen serializes keys and save them in a key ring file.
func keyGen(args []string) {
	flags := args[1:]
	// Check number of flags
	if len(flags) == 4 {
		// TODO: check flags[3] input with checkers.go
		if flags[0] == "-l" && checkIntegerFormat(flags[1]) && flags[2] == "-kr" { // && flags[3] == "" {
			var kr *KeyRing
			n, _ := strconv.Atoi(flags[1])
			l := int64(n)
			sk, pk := phe.GenerateKeys(l)
			if fileExists(flags[3]) {
				rkr := readKeyRing(flags[3])
				kr = addKeyRing(rkr, sk, pk, &phe.Token{})
			} else {
				kr = newKeyRing(sk, pk, &phe.Token{})
			}
			writeKeyRing(kr, flags[3])
			kr.ToString()
		}
	} else {
		// Abort giving the name of the originating function
		abort(args[0])
	}
}

// printKeyGenMenu ...
func printKeyGenMenu() {
	fmt.Println(SystemName + " --keygen -l <positive integer> -kr <key ring file>")
}
