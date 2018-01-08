package sets

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

//P1 of the sets
func P1() {
	const s = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	decode, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	arr := []byte(decode)
	enc := base64.StdEncoding.EncodeToString(arr)
	fmt.Println(enc)
}
