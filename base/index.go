package base

import (
    "runtime"
    "sync"
)

var once sync.Once

type Template struct {
    CpuNum      int      //应用程序所使用的cpu核数，默认为1核
    ReqTimeout  int64    //应用程序内对外发送请求的超时时间，单位：秒
    AppYamlPath []string //应用程序的application.yaml配置文件路径
}

func Init(cfg Template) {
    once.Do(func() {
        //设置cpu核数
        runtime.GOMAXPROCS(runtime.NumCPU())

        //初始化配置文件
        initConfig(cfg.AppYamlPath)

        //redis
        initRedis()

        //mysql
        initMysql()
    })

}
