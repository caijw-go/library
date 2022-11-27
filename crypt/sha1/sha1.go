package sha1

import (
    "crypto/sha1"
    "encoding/hex"
    "strings"
)

func Sha1(text string) string {
    ctx := sha1.New()
    _, err := ctx.Write([]byte(text))
    if err != nil {
        return ""
    }
    return strings.ToLower(hex.EncodeToString(ctx.Sum(nil)))
}

func Sha1Upper(text string) string {
    result := Sha1(text)
    return strings.ToUpper(result)
}
