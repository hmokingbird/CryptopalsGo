package sets

import (
	"encoding/hex"
	"fmt"
	"log"
)

var answer []string

func P3() {
	const s = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	var y = "X"
	decode1, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	arr1 := []byte(decode1)
	xor2 := xor2(arr1, y)
	fmt.Println(xor2)
}

func xor2(a []byte, b string) []string {
	answer = make([]string, len(b))
	index := len(a)
	for y := 0; y < len(b); y++ {
		result := make([]byte, index)
		for i := 0; i < index; i++ {
			result[i] = a[i] ^ b[y]
		}
		answer = append(answer, hex.EncodeToString(result))
	}
	return answer
}
