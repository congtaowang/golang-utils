package crypt

import (
	"encoding/base64"
)

func Encrypt(mode int,orignalData,key []byte)(string,error){
	
	var result []byte 	= orignalData
	var err 	error	= nil

	switch{
		case mode==AES:result,err = AesEncrypt(orignalData,key)
		case mode==DES:result,err = DesEncrypt(orignalData,key)
		default:
	}

	return base64.StdEncoding.EncodeToString(result),err
}

func Decrypt(mode int,encryptData,key []byte)(string,error){
	
	var result []byte 	= encryptData
	var err 	error	= nil

	switch{
		case mode==AES:result,err = AesDecrypt(encryptData,key)
		case mode==DES:result,err = DesDecrypt(encryptData,key)
		default:
	}

	return string(result),err
}

