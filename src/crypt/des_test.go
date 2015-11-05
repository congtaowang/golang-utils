package crypt

import (
	"testing"
	"encoding/base64"
	"fmt"
)


func TestDes(t *testing.T) {
	key := []byte("12345678")
	result, err := DesEncrypt([]byte("It's a english orignal data"), key)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, err := DesDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}

func Test3Des(t *testing.T) {
	key := []byte("sfe023f_sefiel#fi32lf3e!")
	result, err := TripleDesEncrypt([]byte("polaris@studygolang"), key)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, err := TripleDesDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}
