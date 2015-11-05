package crypt

import (
	"encoding/base64"
)

const(
	DES			= 0x00000001
	TRIPLEDES	= 0x00000002
	AES			= 0x00000003
)

// 
func Encrypt(method int,original string,key string)(result string,err error) {
	
	err		= nil
	var data []byte
	switch method {
		case DES:		data,err = DesEncrypt([]byte(original),[]byte(key))
		case TRIPLEDES:data,err = TripleDesEncrypt([]byte(original),[]byte(key))
		case AES:		data,err = AesEncrypt([]byte(original),[]byte(key))
	}
	
	result = base64.StdEncoding.EncodeToString(data)
	return
}

func Decrypt(method int,encrypted string,key string)(result string,err error) {
	
	data,err := base64.StdEncoding.DecodeString(encrypted)
	if err!=nil {
		return
	}
	switch method {
		case DES:		data,err = DesDecrypt(data,[]byte(key))
		case TRIPLEDES:data,err = TripleDesDecrypt(data,[]byte(key))
		case AES:		data,err = AesDecrypt(data,[]byte(key))
	}
	return string(data),err
}


