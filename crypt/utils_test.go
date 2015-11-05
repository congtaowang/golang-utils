package crypt

import (
	"fmt"
	"testing"
)

var desKey  	[]byte = []byte("12345678")//must be 8bits
var des3Key  	[]byte = []byte("123456781234567812345678")//must be 24bits
var aesKey1 	[]byte = []byte("1234567812345678")//must be 16bits
var aesKey2 	[]byte = []byte("123456781234567812345678")//must be 24bits
var aesKey3 	[]byte = []byte("12345678123456781234567812345678")//must be 32bits

var originalData string = "Im's a string waiting for encrypt"

var desEncrypt string
var des3Encrypt string
var aesEncrypt1 string
var aesEncrypt2 string
var aesEncrypt3 string

func TestEncrypt(t *testing.T){
	desEncrypt,_ = Encrypt(DES,originalData,desKey)
	fmt.Println(desEncrypt)
	des3Encrypt,_ = Encrypt(TRIPLEDES,originalData,des3Key)
	fmt.Println(des3Encrypt)
	aesEncrypt1,_ = Encrypt(AES,originalData,aesKey1)
	fmt.Println(aesEncrypt1)
	aesEncrypt2,_ = Encrypt(AES,originalData,aesKey2)
	fmt.Println(aesEncrypt2)
	aesEncrypt3,_ = Encrypt(AES,originalData,aesKey3)
	fmt.Println(aesEncrypt3)
}

func TestDecrypt(t *testing.T){
	var decrypt string
	decrypt,_ = Decrypt(DES,desEncrypt,desKey)
	fmt.Println(decrypt)
	decrypt,_ = Decrypt(TRIPLEDES,des3Encrypt,des3Key)
	fmt.Println(decrypt)
	decrypt,_ = Decrypt(AES,aesEncrypt1,aesKey1)
	fmt.Println(decrypt)
	decrypt,_ = Decrypt(AES,aesEncrypt2,aesKey2)
	fmt.Println(decrypt)
	decrypt,_ = Decrypt(AES,aesEncrypt3,aesKey3)
	fmt.Println(decrypt)
}