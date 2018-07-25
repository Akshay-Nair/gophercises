//Package crypt provides the methods for encryption and decryption.
package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

//this method is to return a cipher which is hashed with md5.
func getHashedKey(key string) (cipher.Block, error) {
	hasher := md5.New()
	hasher.Write([]byte(key))
	cipherKey := hasher.Sum(nil)
	return aes.NewCipher(cipherKey)
}

//Encrypt method encrypts a text with a key and returns string along with  an error if one occurs.
func Encrypt(key string, text string) (string, error) {
	cipherKey, err := getHashedKey(key)
	if err != nil {
		return "", nil
	}
	ciphertext := make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(cipherKey, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(text))
	return fmt.Sprintf("%x", ciphertext), nil
}

//Decrypt function decrypts a encrypted key with a cipher key provided
//it returns a decrypted string along with an error if one occurs.
func Decrypt(key string, hexCode string) (string, error) {
	block, err := getHashedKey(key)
	if err != nil {
		return "", err
	}
	ciphertext, err := hex.DecodeString(hexCode)
	if err != nil {
		return "", err
	}
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("encrypt: cipher too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil
}
