package misc

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"encoding/hex"
	"io"
	"os"
)

func CalcFileMD5(path string) string {
	file, err := os.Open(path)
	if err == nil {
		md5 := crypto.MD5.New()
		_, err := io.Copy(md5, file)
		if err == nil {
			sum := md5.Sum(nil)
			return hex.EncodeToString(sum)
		}
	}
	return ""
}

func padOriginalData(originalData []byte, blockSize int) []byte {
	length := len(originalData)
	toPadLength := blockSize - (4+length)%blockSize
	data := make([]byte, 4)
	binary.BigEndian.PutUint32(data, uint32(length))
	data = append(data, originalData...)
	data = append(data, make([]byte, toPadLength)...)
	return data
}

func AESEncrypt(originalData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCEncrypter(block, key[0:blockSize])
	originalData = padOriginalData(originalData, blockSize)
	encryptedData := make([]byte, len(originalData))
	blockMode.CryptBlocks(encryptedData, originalData)
	return encryptedData, nil
}

func unPadOriginalData(originalData []byte) []byte {
	length := binary.BigEndian.Uint32(originalData)
	return originalData[4 : 4+length]
}

func AESDecrypt(encryptedData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[0:blockSize])
	originalData := make([]byte, len(encryptedData))
	blockMode.CryptBlocks(originalData, encryptedData)
	originalData = unPadOriginalData(originalData)
	return originalData, nil
}
