package main

import (
	"fmt"

	"github.com/hanesbarbosa/phe"
)

// tokGen ...
func tokGen(args []string) {
	flags := args[1:]
	// Check number of flags
	if len(flags) == 2 && flags[0] == "-kr" && fileExists(flags[1]) {
		kr := readKeyRing(flags[1])
		skOld := kr.lastSecretKey()
		pkOld := kr.lastPublicKey()

		l := pkOld.B * 8
		skNew, pkNew := phe.GenerateKeys(l)

		tk := phe.GenerateToken(skOld, skNew, pkOld, pkNew)
		kr = addKeyRing(kr, skNew, pkNew, tk)

		writeKeyRing(kr, flags[1])
		kr.ToString()
	} else {
		// Abort giving the name of the originating function
		abort(args[0])
	}
}

func printTokGenMenu() {
	fmt.Println(SystemName + " tokgen -kr <key ring file>")
}
