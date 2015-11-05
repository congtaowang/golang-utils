package crypt

import (
	"encoding/base64"
	"fmt"
	"testing"
)

var desKey 	 []byte		= []byte("12345678")
var aesKey1 []byte 	= []byte("1234567890123456")
var aesKey2 []byte 	= []byte("012345678901234567890123")
var aesKey3 []byte		= []byte("01234567890123456789012345678901")

var orignalData	 []byte	= []byte("粗大的顶顶顶顶顶")

var result string
var err		error

func TestEncrypt(t *testing.T){
	
	testDESEncrpty()
	testAescrypt(1)
	testAescrypt(2)
	testAescrypt(3)
}

func testAescrypt(index int){
	
	var key []byte
	switch index {
		case 1:key = aesKey1
		case 2:key = aesKey2
		case 3:key = aesKey3
		
	}
	result,err = Encrypt(AES,orignalData,key)
	if isNotNil(err) {
		print(err)
		return
	}
	fmt.Println(result)
	
	resultbyte,_ := base64.StdEncoding.DecodeString(result)
	result,err = Decrypt(AES,resultbyte,key)
	if isNotNil(err) {
		print(err)
		return
	}
	fmt.Println(result)
}

func testDESEncrpty(){
	result,err = Encrypt(DES,orignalData,desKey)
	
	if isNotNil(err) {
		print(err)
		return
	}
	fmt.Println(result)
	
	resultbyte,_ := base64.StdEncoding.DecodeString(result)
	result,err = Decrypt(DES,resultbyte,desKey)
	if isNotNil(err) {
		print(err)
		return
	}
	fmt.Println(result)
}

func isNotNil(obj interface{})bool{
	return !isNil(obj)
}

func isNil(obj interface{})bool{
	return obj==nil
}

func print(obj ...interface{}){
	fmt.Println(obj...)
}