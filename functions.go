package main

import (
	"fmt"
	"regexp"
)

// SystemName ...
const SystemName = "phe-cli"

// ...
func keyGen(flags []string) {
	// Check number of arguments
	if len(flags) != 2 {
		printGeneralMenu()
		exitProgram()
	}
	// [0] = "-l", [1] = "256"
	if flags[0] == "-l" && checkIntegerFormat(flags[1]) {
		// Call key generation function
		// n, _ := strconv.Atoi(flags[1])
		// l := int64(n)
		// sk, pk := phe.GenerateKeys(l)
		// fmt.Println(sk)
	}
}

// WriteKey writes an already packed key to a file.
// func writeKey(fn string, s string) error {
// 	b := []byte(s)
// 	err := ioutil.WriteFile(fn, b, 0600)
// 	return err
// }

// func enc(flags []string) {

// }

// func dec(flags []string) {

// }

// func tokgen(flags []string) {

// }

func checkIntegerFormat(s string) bool {
	matched, err := regexp.MatchString("\\b[0-3]+\\b", s)
	if err != nil {
		return false
	}

	return matched
}

func printKeyGenSignature() {
	fmt.Println(SystemName + " -keygen -l <Positive Integer>")
}

func printEncSignature() {
	fmt.Println(SystemName + " -enc <File> -sk <File> -pk <File> -m <Positive Integer>")
}

func printDecSignature() {
	fmt.Println(SystemName + " -dec -sk <File> -pk <File> -c <String Multivector>")
}

func printTokGenSignature() {
	fmt.Println(SystemName + " tokgen -sk-old <File> -sk-new <File> -pk-old <File> -pk-new <File>")
}
