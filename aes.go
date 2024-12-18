package iutils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"os"
)

// 加密
// 加密文件，并保存
func AesEncryptFilePkcs5(srcFile, key, iv, destFile string) error {
	// todo 因为ioutil废弃了，改为os操作，这里没有测试
	content, err := os.ReadFile(srcFile)
	if err != nil {
		return err
	}
	d, err := AesEncryptPkcs5(content, key, iv)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(string(d))
	if err != nil {
		return err
	}
	return nil
}

func AesEncryptPkcs5(origData []byte, key string, iv string) ([]byte, error) {
	return AesEncryptPkcs5Byte(origData, []byte(key), []byte(iv))
}

func AesEncryptPkcs5Byte(origData []byte, key []byte, iv []byte) ([]byte, error) {
	return AesEncrypt(origData, key, iv, PKCS5Padding)
}

func AesEncrypt(origData []byte, key []byte, iv []byte, paddingFunc func([]byte, int) []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = paddingFunc(origData, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// 解密
func AesDecryptPkcs5(crypted []byte, key string, iv string) ([]byte, error) {
	return AesDecryptPkcs5Byte(crypted, []byte(key), []byte(iv))
}

func AesDecryptPkcs5Byte(crypted []byte, key []byte, iv []byte) ([]byte, error) {
	return AesDecrypt(crypted, key, iv, PKCS5UnPadding)
}

func AesDecrypt(crypted, key []byte, iv []byte, unPaddingFunc func([]byte) []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = unPaddingFunc(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	if length < unpadding {
		return []byte("unpadding error")
	}
	return origData[:(length - unpadding)]
}
