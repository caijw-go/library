package auth

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/caijw-go/library/base"
    "github.com/caijw-go/library/tool"
    "github.com/gin-gonic/gin"
    "net/http"
    "strings"
)

const contextUserKey = "loggedUser"

func getTokenFromHeader(c *gin.Context) (string, error) { //从请求头中获得token
    header := c.GetHeader("Authorization")
    tmp := strings.Split(header, " ")
    if len(tmp) != 2 || tmp[0] != "Basic" || len(tmp[1]) != 32 {
        return "", errors.New("user not login")
    }
    return tmp[1], nil
}

func saveUser[T any](token string, u T) error { //将用户信息保存到redis里
    str, err := json.Marshal(u)
    if err != nil {
        return err
    }
    if err := base.Redis(config.RedisName).Set(fmt.Sprintf(config.RedisKey, token), str, config.RedisTtl).Err(); err != nil {
        return err
    }
    return nil
}

func GetUserByToken[T any](token string, expireWhenExist bool) (T, error) {
    bytes, err := base.Redis(config.RedisName).Get(fmt.Sprintf(config.RedisKey, token)).Bytes()
    if err != nil {
        return nil, err
    }
    var u T
    if err = json.Unmarshal(bytes, &u); err != nil {
        return nil, err
    }
    if expireWhenExist {
        if err = base.Redis(config.RedisName).Expire(fmt.Sprintf(config.RedisKey, token), config.RedisTtl).Err(); err != nil {
            return nil, err
        }
    }
    return u, nil
}

func Middleware[T any]() gin.HandlerFunc { //校验登录中间件
    return func(c *gin.Context) {
        token, err := getTokenFromHeader(c)
        if err != nil {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }
        u, err := GetUserByToken[T](token, false)
        if err != nil {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }
        c.Set(contextUserKey, u)
        c.Next()
    }
}

func Login[T any](u T) (string, error) { //登录
    token := tool.GenUUID()
    if err := saveUser[T](token, u); err != nil {
        return "", err
    }
    return token, nil
}

func GetUser[T any](c *gin.Context) T { //从Context中取出User
    u, _ := c.Get(contextUserKey)
    return u.(T)
}

func Change[T any](c *gin.Context, u T) error { //修改用户信息，用户在登录过程中如果有需要修改的数据，需要进行修改
    token, err := getTokenFromHeader(c)
    if err != nil {
        return err
    }
    if err = saveUser[T](token, u); err != nil {
        return err
    }
    c.Set(contextUserKey, u)
    return nil
}

func Logout(c *gin.Context) bool { //退出登录
    token, err := getTokenFromHeader(c)
    if err != nil {
        return false
    }
    RemoveToken(token)
    return true
}

func RemoveToken(token string) {
    base.Redis(config.RedisName).Del(fmt.Sprintf(config.RedisKey, token))
}
