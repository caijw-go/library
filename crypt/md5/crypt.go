package md5

import (
    "crypto/md5"
    "encoding/hex"
    "strings"
)

// MD5 32位小写MD5
func MD5(text string) string {
    return genByBytes([]byte(text))
}

func genByBytes(bytes []byte) string {
    ctx := md5.New()
    _, err := ctx.Write(bytes)
    if err != nil {
        return ""
    }
    return strings.ToLower(hex.EncodeToString(ctx.Sum(nil)))
}
