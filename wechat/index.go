package wechat

import (
    "sync"
)

const domain = "https://api.weixin.qq.com"

var once sync.Once
var config Config

type Config struct {
    Appid                string
    Secret               string
    AccessTokenRedisName string
    AccessTokenRedisKey  string
}

func Init(conf Config) {
    once.Do(func() {
        config = conf
    })
}

func GetAppid() string {
    return config.Appid
}
