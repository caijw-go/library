package auth

import (
    "time"
)

var config Config

type Config struct {
    RedisName string
    RedisKey  string //需要预留出%s用于存储token
    RedisTtl  time.Duration
}

func Init(conf Config) {
    config = conf
}
