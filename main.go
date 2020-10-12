package main

import (
	"os"
)

func main() {

	if len(os.Args) <= 3 {
		printGeneralMenu()
		exitProgram()
	}

	// Flag comming from the user
	checkInputs(os.Args[1:])

	// switch f {
	// case "-k":
	// 	// Generate key pair
	// 	// cmd -k bits
	// 	k := generateKey(os.Args[2])
	// 	fmt.Println(k)
	// case "-e":
	// 	// Encryption
	// 	// cmd -e number s1 s2
	// 	c := encrypt(os.Args[2], os.Args[3], os.Args[4])
	// 	fmt.Println(c)
	// case "-d":
	// 	// Decryption
	// 	// cmd -d c s1 s2
	// 	n := decrypt(os.Args[2], os.Args[3], os.Args[4])
	// 	fmt.Println(n)
	// case "-t":
	// 	// Generate tokens
	// 	tks := generateTokens(os.Args[2], os.Args[3], os.Args[4], os.Args[5])
	// 	fmt.Println("[", tks[0].ToString(), ",", tks[1].ToString(), "]")
	// case "-u":
	// 	// Key Update
	// 	// Should be: cmd -u c t1 t2
	// 	mc := keyUpdate(os.Args[2], os.Args[3], os.Args[4])
	// 	fmt.Println(mc.ToString())
	// default:
	// 	fmt.Println("-1")
	// }
}

// func encrypt(n, s1, s2 string) string {
// 	mn := edche.NumberToMultivector(n)
// 	ms1 := edche.StringToMultivector(s1)
// 	ms2 := edche.StringToMultivector(s2)

// 	em := edche.Encrypt(mn, ms1, ms2)
// 	return em.ToString()
// }

// func decrypt(c, s1, s2 string) string {
// 	mc := edche.StringToMultivector(c)
// 	ms1 := edche.StringToMultivector(s1)
// 	ms2 := edche.StringToMultivector(s2)

// 	dm := edche.Decrypt(mc, ms1, ms2)
// 	n := edche.MultivectorToNumber(dm)
// 	return n.String()
// }

// func generateTokens(s1, s2, s3, s4 string) [2]edche.Multivector {
// 	ms1 := edche.StringToMultivector(s1)
// 	ms2 := edche.StringToMultivector(s2)
// 	ms3 := edche.StringToMultivector(s3)
// 	ms4 := edche.StringToMultivector(s4)
// 	tks := edche.GenerateTokens(ms1, ms2, ms3, ms4)
// 	return tks
// }

// func keyUpdate(c, t1, t2 string) edche.Multivector {
// 	mc1 := edche.StringToMultivector(c)
// 	mt1 := edche.StringToMultivector(t1)
// 	mt2 := edche.StringToMultivector(t2)
// 	mc2 := edche.KeyUpdate(mc1, mt1, mt2)
// 	return mc2
// }
