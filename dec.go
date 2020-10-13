package main

import (
	"fmt"

	"github.com/hanesbarbosa/phe"
)

func dec(args []string) {
	flags := args[1:]
	// Check number of flags
	if len(flags) == 4 {
		if flags[0] == "-kr" && flags[2] == "-c" && checkMultivectorFormat(flags[3]) {
			kr := readKeyRing(flags[1])

			c := phe.StringToMultivector(flags[3])

			sk := kr.lastSecretKey()
			pk := kr.lastPublicKey()

			m := phe.Decrypt(sk, pk, c)

			fmt.Println(m.String())
		}
	} else {
		// Abort giving the name of the originating function
		abort(args[0])
	}
}

func printDecMenu() {
	fmt.Println(SystemName + " -dec -kr <key ring file> -c <string multivector>")
}
