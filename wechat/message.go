package wechat

import (
    "encoding/json"
    "fmt"
    "github.com/caijw-go/library/wechat/util"
    "github.com/imroc/req"
)

const sendMessageUrl = domain + "/cgi-bin/message/subscribe/send?access_token=%s"

type SendMessageReq struct {
    TemplateId       string                            `json:"template_id"`       // 所需下发的订阅模板id
    ToUser           string                            `json:"touser"`            // 接收者（用户）的 openid
    Data             map[string]map[string]interface{} `json:"data"`              // 模板内容，格式形如 { "key1": { "value": any }, "key2": { "value": any } }的object
    MiniprogramState string                            `json:"miniprogram_state"` // 【选填】跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
    Page             string                            `json:"page"`              // 【选填】点击模板卡片后的跳转页面，支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转
    Lang             string                            `json:"lang"`              // 【选填】进入小程序查看”的语言类型
}

type sendMessageResp struct {
    *util.CommonError
}

func SendMessage(request *SendMessageReq) *sendMessageResp {
    accessToken, commonError := getAccessToken()
    if commonError != nil {
        return &sendMessageResp{
            CommonError: commonError,
        }
    }
    if request.MiniprogramState == "" {
        request.MiniprogramState = "formal"
    }
    if request.Lang == "" {
        request.Lang = "zh_CN"
    }
    resp, err := req.Post(fmt.Sprintf(sendMessageUrl, accessToken), req.BodyJSON(request))
    if err != nil {
        return &sendMessageResp{
            CommonError: util.NewError(500, "request wechat code2session error", err),
        }
    }
    result := &sendMessageResp{}
    if err = json.Unmarshal(resp.Bytes(), &result); err != nil {
        return &sendMessageResp{
            CommonError: util.NewError(500, "json unmarshal error", err),
        }
    }
    return result
}
