package sts

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
    "github.com/caijw-go/library/log"
    "github.com/caijw-go/library/types/convert"
    "github.com/gin-gonic/gin"
    "time"
)

type Config struct {
    RegionId        string
    AccessKeyId     string
    AccessKeySecret string
    RoleArn         string
}

// GetAssumeRole RAM用户调用AssumeRole接口获取一个扮演RAM角色的临时身份凭证（STS Token）
// 即RAM里的用户，申请成为某个角色，成功后返回临时凭证，用这个凭证可以操作角色权限内的功能
// doc https://help.aliyun.com/document_detail/28763.htm?spm=a2c4g.11186623.0.0.70a464a811R415#reference-clc-3sv-xdb
func GetAssumeRole(config Config) (*sts.Credentials, error) {
    client, err := sts.NewClientWithAccessKey(config.RegionId, config.AccessKeyId, config.AccessKeySecret)
    if err != nil {
        log.Error("ali sts getAssumeRole.NewClientWithAccessKey error", err.Error())
        return nil, err
    }
    //构建请求对象。
    request := sts.CreateAssumeRoleRequest()
    request.Scheme = "https"

    //设置参数。关于参数含义和设置方法，请参见API参考。
    request.RoleArn = config.RoleArn
    request.RoleSessionName = convert.MustString(time.Now().Unix())

    //发起请求，并得到响应。
    response, err := client.AssumeRole(request)
    if err != nil {
        log.Error("ali getAssumeRole.AssumeRole error", gin.H{
            "errMsg":  err.Error(),
            "request": response,
        })
        return nil, err
    }
    return &response.Credentials, nil
}
