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
    PayConfig            *PayConfig
}

func Init(conf Config) {
    once.Do(func() {
        config = conf
        if conf.PayConfig != nil {
            initPay(conf.PayConfig)
        }
    })
}

func GetAppid() string {
    return config.Appid
}
