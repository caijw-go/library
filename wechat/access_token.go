package wechat

import (
    "github.com/caijw-go/library/base"
    "github.com/caijw-go/library/wechat/util"
    "github.com/imroc/req"
    jsonIter "github.com/json-iterator/go"
    "time"
)

const accessTokenUrl = domain + "/cgi-bin/token"

func getAccessToken() (string, *util.CommonError) {
    accessToken := base.Redis(config.AccessTokenRedisName).Get(config.AccessTokenRedisKey).Val()
    if accessToken != "" {
        return accessToken, nil
    }
    resp, err := req.Get(accessTokenUrl, req.QueryParam{
        "appid":      config.Appid,
        "secret":     config.Secret,
        "grant_type": "client_credential",
    })
    if err != nil {
        return "", util.NewError(500, "request wechat code2session error", err)
    }
    result := jsonIter.Get(resp.Bytes())
    if errCode := result.Get("errcode").ToInt64(); errCode != 0 {
        return "", &util.CommonError{
            ErrCode: errCode,
            ErrMsg:  result.Get("errmsg").ToString(),
        }
    }
    accessToken = result.Get("access_token").ToString()
    base.Redis(config.AccessTokenRedisName).SetNX(config.AccessTokenRedisKey, accessToken, time.Duration(result.Get("expires_in").ToInt64()-5*60)*time.Second)

    return accessToken, nil
}
