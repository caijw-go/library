package server

import (
    "github.com/caijw-go/library/base"
    "github.com/caijw-go/library/middlewares"
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
)

type FuncController func(engine *gin.Engine)

type Config struct {
    AppName            string            //app name
    Address            string            //服务地址，如：0.0.0.0:8080
    FuncController     FuncController    //初始化路由函数
    StaticRelativePath string            //静态文件请求相对路径
    StaticRoot         string            //静态文件基础目录
    HtmlPattern        string            //html模板路径解析规则
    BeforeMiddleware   []gin.HandlerFunc //中间件之前运行的所有中间件
}

func IsProd() bool {
    return base.Config().GetBool("server.isProd")
}

func AppName() string {
    return base.Config().GetString("server.name")
}

// Create 创建配置
func Create(cfg Config) {
    engine := gin.New()
    engine.Use(cfg.BeforeMiddleware...)
    if IsProd() {
        gin.SetMode(gin.ReleaseMode)
    } else {
        engine.Use(gin.Logger())
    }

    if sentryDsn := base.Config().GetString("sentryDsn"); sentryDsn != "" {
        env := "prod"
        if !IsProd() {
            env = "local"
        }
        engine.Use(middlewares.Recovery(AppName(), env, sentryDsn))
    }

    //设置静态文件
    if cfg.StaticRelativePath != "" && cfg.StaticRoot != "" {
        engine.Static(cfg.StaticRelativePath, cfg.StaticRoot)
    }
    //加载视图模板
    if cfg.HtmlPattern != "" {
        engine.LoadHTMLGlob(cfg.HtmlPattern)
    }

    //执行路由函数
    if cfg.FuncController != nil {
        cfg.FuncController(engine)
    }

    if cfg.Address == "" {
        cfg.Address = "0.0.0.0:8080"
    }
    log.Printf("server start at %v", cfg.Address)
    s := &http.Server{Addr: cfg.Address, Handler: engine}
    if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Printf("ding service start:%v", err)
    }
}
