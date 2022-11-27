package middlewares

import (
    "fmt"
    "github.com/getsentry/sentry-go"
    "github.com/gin-gonic/gin"
    "github.com/pkg/errors"
    "log"
    "time"
)

func Recovery(appName, env, sentryDsn string) gin.HandlerFunc {
    log.Println("MiddleWare Recovery Init")
    if sentryDsn != "" {
        log.Println("Sentry", sentry.Init(sentry.ClientOptions{
            Environment:  "APP_MODE=" + env,
            Release:      appName,
            IgnoreErrors: []string{"write tcp"}, //过滤掉你不感兴趣的
            Dsn:          sentryDsn,
        }))
    }

    return func(c *gin.Context) {
        defer func() {
            if e := recover(); e != nil {
                if sentryDsn != "" {
                    sentry.CaptureException(errors.New(fmt.Sprintf("%v", e)))
                    sentry.Flush(time.Second)
                }
                log.Println("Recovery.Gin", e)
            }
        }()

        c.Next()
    }
}
