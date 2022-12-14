package tool

import (
    "github.com/caijw-go/library/crypt/md5"
    "github.com/gofrs/uuid"
)

func GenUUID() string {
    return md5.MD5(uuid.Must(uuid.NewV4()).String())
}
