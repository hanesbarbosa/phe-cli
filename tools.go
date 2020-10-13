package main

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"io/ioutil"
	"os"
)

func fileExists(f string) bool {
	a, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return !a.IsDir()
}

// serializeKeyRing serializes struct to base64
func serializeKeyRing(kr *KeyRing) string {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(kr)
	if err != nil {
		panic("gob encoding failed: " + err.Error())
	}
	return base64.StdEncoding.EncodeToString(b.Bytes())
}

// unserializeKeyRing unserializes struct from base64
func unserializeKeyRingFromBase64(str string) *KeyRing {
	kr := &KeyRing{}
	by, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		panic("base64 decoding failed: " + err.Error())
	}
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	err = d.Decode(&kr)
	if err != nil {
		panic("gob decoding failed: " + err.Error())
	}
	return kr
}

// writeKeyRing writes a serialized key ring to a file.
func writeFile(fn string, s string) {
	b := []byte(s)
	err := ioutil.WriteFile(fn, b, 0600)
	if err != nil {
		panic("Error creating file: " + err.Error())
	}
}

// readKeyRing reads a serialized key ring from a file.
func readFile(fn string) ([]byte, error) {
	b, err := ioutil.ReadFile(fn)
	return b, err
}

func writeKeyRing(kr *KeyRing, filename string) string {
	skr := serializeKeyRing(kr)
	if filename == "" {
		filename = "keyring.kr"
	}
	writeFile(filename, skr)
	return filename
}

func readKeyRing(fn string) *KeyRing {
	b, err := readFile(fn)
	if err != nil {
		panic("Error reading file: " + err.Error())
	}
	kr := unserializeKeyRingFromBase64(string(b))
	return kr
}
