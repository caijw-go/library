package robot

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/imroc/req"
)

type Robot struct {
    Url       string
    isAtAll   bool
    atMobiles []string
    atUserIds []string
    remark    string
}

func Init(accessToken string) *Robot {
    return &Robot{
        Url: fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", accessToken),
    }
}

func InitBySessionWebhook(sessionWebhook string) *Robot {
    return &Robot{
        Url: sessionWebhook,
    }
}

func (t *Robot) SetAtMobiles(mobiles ...string) *Robot {
    if len(mobiles) > 0 {
        t.atMobiles = append(t.atMobiles, mobiles...)
    }
    return t
}

func (t *Robot) SetAtUserIds(UserIds ...string) *Robot {
    if len(UserIds) > 0 {
        t.atUserIds = append(t.atUserIds, UserIds...)
    }
    return t
}

func (t *Robot) SetAtAll() *Robot {
    t.isAtAll = true
    return t
}

func (t *Robot) SetRemark(remark string) *Robot {
    t.remark = remark
    return t
}

func (t *Robot) SendText(content string) {
    t.send("text", gin.H{"content": content})
}

func (t *Robot) send(msgType string, data gin.H) {
    var requestData = gin.H{
        "msgtype": msgType,
        "at": gin.H{
            "atMobiles": t.atMobiles,
            "atUserIds": t.atUserIds,
            "isAtAll":   t.isAtAll,
        },
    }
    requestData[msgType] = data
    req.Post(t.Url, req.BodyJSON(requestData)) //忽略报错
}
