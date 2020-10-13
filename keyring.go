package main

import (
	"fmt"
	"strconv"

	"github.com/hanesbarbosa/phe"
)

// KeyRing stores a track of secret and public keys along with tokens
type KeyRing struct {
	SK []*phe.SecretKey
	PK []*phe.PublicKey
	TK []*phe.Token
}

func newKeyRing(sk *phe.SecretKey, pk *phe.PublicKey, tk *phe.Token) *KeyRing {
	kr := &KeyRing{}
	kr.SK = append(kr.SK, sk)
	kr.PK = append(kr.PK, pk)
	kr.TK = append(kr.TK, tk)
	return kr
}

func addKeyRing(kr *KeyRing, sk *phe.SecretKey, pk *phe.PublicKey, tk *phe.Token) *KeyRing {
	kr.SK = append(kr.SK, sk)
	kr.PK = append(kr.PK, pk)
	kr.TK = append(kr.TK, tk)
	return kr
}

func (kr *KeyRing) lastSecretKey() *phe.SecretKey {
	return kr.SK[len(kr.SK)-1]
}

func (kr *KeyRing) lastPublicKey() *phe.PublicKey {
	return kr.PK[len(kr.PK)-1]
}

func (kr *KeyRing) lastToken() *phe.Token {
	return kr.TK[len(kr.TK)-1]
}

// ToString ...
func (kr *KeyRing) ToString() {
	i := 0
	for i < len(kr.SK) {
		is := strconv.Itoa(i)
		fmt.Println()
		fmt.Println("sk[" + is + "].K1: " + kr.SK[i].K1.ToString())
		fmt.Println("sk[" + is + "].K2: " + kr.SK[i].K2.ToString())
		fmt.Println("sk[" + is + "].G: " + kr.SK[i].G.String())
		fmt.Println("pk[" + is + "].B: " + strconv.FormatInt(kr.PK[i].B, 10))
		fmt.Println("pk[" + is + "].Q: " + kr.PK[i].Q.String())
		if kr.TK[i].T1 != nil && kr.TK[i].T2 != nil {
			fmt.Println("tk[" + is + "].TK1: " + kr.TK[i].T1.ToString())
			fmt.Println("tk[" + is + "].TK2: " + kr.TK[i].T2.ToString())
		}
		i = i + 1
	}
}
