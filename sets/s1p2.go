package sets

import (
	"encoding/hex"
	"fmt"
	"log"
)

func P2() {
	const s, y = "1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965"
	decode1, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	arr1 := []byte(decode1)
	decode2, err := hex.DecodeString(y)
	if err != nil {
		log.Fatal(err)
	}
	arr2 := []byte(decode2)
	xord := xor(arr1, arr2)
	enc := hex.EncodeToString(xord)
	fmt.Println(enc)
}

func xor(a []byte, b []byte) []byte {
	index := len(a)
	result := make([]byte, index)
	for i := 0; i < index; i++ {
		result[i] = a[i] ^ b[i]
	}
	return (result)
}
