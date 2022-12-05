package wechat

import (
    "encoding/json"
    "fmt"
    "github.com/caijw-go/library/wechat/util"
    "github.com/gin-gonic/gin"
    "github.com/imroc/req"
)

const jsCodeToSessionUrl = domain + "/sns/jscode2session"
const getPhoneNumberUrl = domain + "/wxa/business/getuserphonenumber?access_token=%s"

type jsCodeToSessionResp struct {
    *util.CommonError

    OpenID     string `json:"openid"`      // 用户唯一标识
    SessionKey string `json:"session_key"` // 会话密钥
    UnionID    string `json:"unionid"`     // 用户在开放平台的唯一标识符，在满足UnionID下发条件的情况下会返回
}

func JsCodeToSession(code string) *jsCodeToSessionResp {
    resp, err := req.Get(jsCodeToSessionUrl, req.QueryParam{
        "appid":      config.Appid,
        "secret":     config.Secret,
        "grant_type": "authorization_code",
        "js_code":    code,
    })
    if err != nil {
        return &jsCodeToSessionResp{
            CommonError: util.NewError(500, "request wechat code2session error", err),
        }
    }
    result := &jsCodeToSessionResp{}
    if err = json.Unmarshal(resp.Bytes(), &result); err != nil {
        return &jsCodeToSessionResp{
            CommonError: util.NewError(500, "json unmarshal error", err),
        }
    }
    return result
}

type phoneInfo struct {
    PhoneNumber     string `json:"phoneNumber"`
    PurePhoneNumber string `json:"purePhoneNumber"`
    CountryCode     uint   `json:"countryCode"`
}
type getPhoneNumberResp struct {
    *util.CommonError

    PhoneInfo phoneInfo `json:"phone_info"`
}

func GetPhoneNumber(code string) *getPhoneNumberResp {
    accessToken, commonError := getAccessToken()
    if commonError != nil {
        return &getPhoneNumberResp{
            CommonError: commonError,
        }
    }
    resp, err := req.Post(fmt.Sprintf(getPhoneNumberUrl, accessToken), req.BodyJSON(gin.H{"code": code}))
    result := &getPhoneNumberResp{}
    if err = json.Unmarshal(resp.Bytes(), &result); err != nil {
        return &getPhoneNumberResp{
            CommonError: util.NewError(500, "json unmarshal error", err),
        }
    }
    return result
}
