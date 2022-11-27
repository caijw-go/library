package base

import (
    "github.com/caijw-go/library/types/convert"
    "github.com/go-redis/redis"
    "log"
)

var _listRedis = make(map[string]*redis.Client)

func initRedis() {
    res := Config().GetStringMap("redis")
    if len(res) == 0 {
        return
    }

    if _, ok := res["address"]; ok { //如果是单链接，构建成二维数组的形式
        res = map[string]interface{}{"default": res}
    }
    for name, value := range res {
        _listRedis[name] = createRedis(value)
        log.Printf("base.redis [%s] connection success", name)
    }
}

func createRedis(values interface{}) *redis.Client {
    cfg := values.(map[string]interface{})
    conn := redis.NewClient(&redis.Options{
        Addr:     convert.MustString(cfg["address"]),
        Password: convert.MustString(cfg["password"]),
        DB:       convert.MustInt(cfg["db"]),
        PoolSize: convert.MustInt(cfg["poolsize"]),
    })
    if err := conn.Ping().Err(); err != nil {
        log.Fatalf("base redis connection error : %s", err.Error())
    }
    return conn
}

func Redis(name ...string) *redis.Client {
    realName := "default"
    if len(name) > 0 {
        realName = name[0]
    }
    if i, ok := _listRedis[realName]; ok {
        return i
    }
    return nil
}
