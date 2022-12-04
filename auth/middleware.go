package auth

import (
    "errors"
    "fmt"
    "github.com/caijw-go/library/base"
    "github.com/gin-gonic/gin"
    "net/http"
    "strings"
)

const contextUserKey = "loggedUser"

func Middleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token, err := getTokenFromHeader(c)
        if err != nil {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }
        if userUniqId, err := resolveUserUniqId(token); err != nil {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        } else {
            c.Set(contextUserKey, userUniqId)
            c.Next()
        }
    }
}

// getTokenFromHeader 从请求头中获得token
func getTokenFromHeader(c *gin.Context) (string, error) {
    header := c.GetHeader("Authorization")
    tmp := strings.Split(header, " ")
    if len(tmp) != 2 || tmp[0] != "Basic" || len(tmp[1]) != 32 {
        return "", errors.New("user not login")
    }
    return tmp[1], nil
}

// resolveUserUniqId 通过token解析出用户唯一id，这个唯一id是字符串，格式由调用者自行组织
func resolveUserUniqId(token string) (string, error) {
    if len(token) != 32 {
        return "", errors.New("resolveUserUniqId token error")
    }
    bytes, err := base.Redis(config.RedisName).Get(fmt.Sprintf(config.RedisKey, token)).Bytes()
    if err != nil {
        return "", errors.New("resolveUserUniqId redis get error" + err.Error())
    }
    return string(bytes), nil
}
