package base

import (
    "github.com/caijw-go/library/types/convert"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "log"
)

var _listMysql = make(map[string]*gorm.DB)

func initMysql() {
    res := Config().GetStringMap("mysql")
    if len(res) == 0 {
        return
    }

    if _, ok := res["address"]; ok { //如果是单链接，构建成二维数组的形式
        res = map[string]interface{}{"default": res}
    }

    for name, value := range res {
        _listMysql[name] = createMysql(value)
        log.Printf("base.mysql [%s] connection success", name)
    }
}

func createMysql(cfg interface{}) *gorm.DB {
    maps := cfg.(map[string]interface{})
    config := &gorm.Config{}
    if convert.MustBool(maps["logmode"]) { //info
        config.Logger = logger.Default.LogMode(logger.Info)
    } else { //silent
        config.Logger = logger.Default.LogMode(logger.Silent)
    }
    conn, err := gorm.Open(mysql.Open(maps["address"].(string)), config)
    if err != nil {
        log.Fatalf("base mysql open error : %s", err.Error())
        return nil
    }
    sqlDB, err := conn.DB()
    if err != nil {
        log.Fatalf("base mysql connection error : %s", err.Error())
        return nil
    }
    if maxIdle, ok := maps["maxidle"]; ok {
        sqlDB.SetMaxIdleConns(convert.MustInt(maxIdle))
    }
    if maxOpen, ok := maps["maxopen"]; ok {
        sqlDB.SetMaxOpenConns(convert.MustInt(maxOpen))
    }
    return conn
}

func DB(name ...string) *gorm.DB {
    realName := "default"
    if len(name) > 0 {
        realName = name[0]
    }
    if i, ok := _listMysql[realName]; ok {
        return i
    }
    return nil
}
