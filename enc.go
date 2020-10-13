package main

import (
	"fmt"
	"math/big"

	"github.com/hanesbarbosa/phe"
)

func enc(args []string) {
	flags := args[1:]
	// Check number of flags
	if len(flags) == 4 {
		if flags[0] == "-kr" && flags[2] == "-m" && checkIntegerFormat(flags[3]) {
			kr := readKeyRing(flags[1])

			m := big.NewInt(0)
			m.SetString(flags[3], 10)

			sk := kr.lastSecretKey()
			pk := kr.lastPublicKey()

			c := phe.Encrypt(sk, pk, m)

			fmt.Println(c.ToString())
		}
	} else {
		// Abort giving the name of the originating function
		abort(args[0])
	}
}

func printEncMenu() {
	fmt.Println(SystemName + " -enc -kr <key ring file> -m <positive integer>")
}
