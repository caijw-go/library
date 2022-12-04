package auth

import (
    "errors"
    "fmt"
    "github.com/caijw-go/library/base"
    "github.com/caijw-go/library/tool"
    "github.com/gin-gonic/gin"
)

//Login 登录
func Login(userUniqId string) (string, error) {
    token := tool.GenUUID()
    if err := base.Redis(config.RedisName).Set(fmt.Sprintf(config.RedisKey, token), userUniqId, config.RedisTtl).Err(); err != nil {
        return "", errors.New("adminUserLogin redis.Set error" + err.Error())
    }
    return token, nil
}

//GetUserUniqId 从Context中取出UserUniqId
func GetUserUniqId(c *gin.Context) string {
    username, _ := c.Get(contextUserKey)
    return username.(string)
}

//Logout 退出登录
func Logout(c *gin.Context) bool {
    token, err := getTokenFromHeader(c)
    if err != nil {
        return false
    }
    base.Redis(config.RedisName).Del(fmt.Sprintf(config.RedisKey, token))
    return true
}
