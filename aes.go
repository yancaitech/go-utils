package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
)

// PKCS7Padding func
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7UnPadding func
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// AesEncrypt func
func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// AesDecrypt func
func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

// AesEncryptString func
func AesEncryptString(data, key string) (en string, err error) {
	h := sha256.New()
	h.Write([]byte(key))
	bskey := h.Sum(nil)

	bs, err := AesEncrypt([]byte(data), bskey)
	if err != nil {
		return "", err
	}
	en, err = BytesToBase(bs)
	if err != nil {
		return "", err
	}
	return en, nil
}

// AesDecryptString func
func AesDecryptString(en, key string) (data string, err error) {
	h := sha256.New()
	h.Write([]byte(key))
	bskey := h.Sum(nil)

	bs, err := BaseToBytes(string(en))
	if err != nil {
		return "", err
	}
	bs, err = AesDecrypt(bs, bskey)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
