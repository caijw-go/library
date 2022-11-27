package aes

import (
    "fmt"
    "testing"
)

func TestAesCBC_Encrypt(t *testing.T) {
    res, err := CBCEncrypt("1234567890123456", "hello")
    if err != nil {
        t.Fatal(err)
    } else {
        fmt.Println(res)
    }
}

func TestAesCBC_Decrypt(t *testing.T) {
    res, err := CBCDecrypt("1234567890123456", "ObBxtb9plyPvM6ZEdBv6MQ==")
    if err != nil {
        t.Fatal(err)
    } else {
        fmt.Println(res)
    }
}
