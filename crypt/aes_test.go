package crypt

import (
	"testing"
	"fmt"
	"encoding/base64"
)

func TestAes(t *testing.T) {
	key := []byte("aaaaaaaaaaaaaaaa")
	result, err := AesEncrypt([]byte("polaris@studygolang"), key)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, err := AesDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}