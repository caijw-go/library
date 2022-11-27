package base64

import (
    "encoding/base64"
)

func Encode(value string) string {
    return base64.StdEncoding.EncodeToString([]byte(value))
}

func Decode(value string) string {
    result, err := base64.StdEncoding.DecodeString(value)
    if err != nil {
        return ""
    }
    return string(result)
}
