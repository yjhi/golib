package jutils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
)

func Md5(content string) string {
	w := md5.New()
	io.WriteString(w, content)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

func Sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func SHA1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%X", t.Sum(nil))
}

func MD5(content string) string {
	w := md5.New()
	io.WriteString(w, content)
	md5str := fmt.Sprintf("%X", w.Sum(nil))
	return md5str
}

func _fPKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func _fPKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//AES加密,CBC
func AesBase64Encode(origstr string, keystr string) (string, error) {

	key := []byte(keystr)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	origData := []byte(origstr)
	blockSize := block.BlockSize()

	origData = _fPKCS7Padding(origData, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])

	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)

	return base64.StdEncoding.EncodeToString(crypted), nil
}

//AES解密
func AesBase64Decode(data string, keystr string) (string, error) {
	key := []byte(keystr)

	crypted, err := base64.StdEncoding.DecodeString(data)

	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))

	blockMode.CryptBlocks(origData, crypted)
	origData = _fPKCS7UnPadding(origData)

	return string(origData), nil
}
