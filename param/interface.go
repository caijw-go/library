package param

import (
    "github.com/caijw-go/library/log"
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
)

func Validate(c *gin.Context, parameter interface{}) error {
    var err error
    if c.Request.Method == "GET" {
        err = c.ShouldBindQuery(parameter)
    } else if c.ContentType() == gin.MIMEJSON {
        err = c.ShouldBindBodyWith(parameter, binding.JSON)
    } else {
        err = c.ShouldBind(parameter)
    }
    if err != nil {
        log.Info("参数校验失败", err.Error())
    }
    return err
}
