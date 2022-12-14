package aes

import (
    "bytes"
    "crypto/aes"
    "crypto/cipher"
    "encoding/base64"
)

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
    padding := blockSize - len(ciphertext)%blockSize
    //填充
    padText := bytes.Repeat([]byte{byte(padding)}, padding)

    return append(ciphertext, padText...)
}

func PKCS5UnPadding(origData []byte) []byte {
    length := len(origData)
    unPadding := int(origData[length-1])
    return origData[:(length - unPadding)]
}

func CBCEncrypt(key, content string) (string, error) {
    keyBytes := []byte(key)
    encodeBytes := []byte(content)
    //根据key 生成密文
    block, err := aes.NewCipher(keyBytes)
    if err != nil {
        return "", err
    }

    //blockSize := block.BlockSize()
    encodeBytes = PKCS5Padding(encodeBytes, 16)

    blockMode := cipher.NewCBCEncrypter(block, keyBytes)
    crypto := make([]byte, len(encodeBytes))
    blockMode.CryptBlocks(crypto, encodeBytes)

    return base64.StdEncoding.EncodeToString(crypto), nil
}

func CBCDecrypt(key, content string) (string, error) {
    keyBytes := []byte(key)
    //先解密base64
    decodeBytes, err := base64.StdEncoding.DecodeString(content)
    if err != nil {
        return "", err
    }

    block, err := aes.NewCipher(keyBytes)
    if err != nil {
        return "", err
    }
    blockMode := cipher.NewCBCDecrypter(block, keyBytes)
    origData := make([]byte, len(decodeBytes))

    blockMode.CryptBlocks(origData, decodeBytes)
    origData = PKCS5UnPadding(origData)
    return string(origData), nil
}
