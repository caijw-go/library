package wechat

import (
    "context"
    "github.com/wechatpay-apiv3/wechatpay-go/core"
    "github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
    "github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
    "github.com/wechatpay-apiv3/wechatpay-go/core/notify"
    "github.com/wechatpay-apiv3/wechatpay-go/core/option"
    "github.com/wechatpay-apiv3/wechatpay-go/utils"
    "log"
)

type PayConfig struct {
    ApiClientKeyPath           string
    MchId                      string
    MchCertificateSerialNumber string
    MchAPIv3Key                string
}

type pay struct {
    Config        *PayConfig
    Client        *core.Client
    NotifyHandler *notify.Handler
}

var Pay = new(pay)

func InitPay(conf *PayConfig) {
    once.Do(func() {
        mchPrivateKey, err := utils.LoadPrivateKeyWithPath(conf.ApiClientKeyPath)
        if err != nil {
            log.Fatal("load merchant private key error")
        }
        ctx := context.Background()
        opts := []core.ClientOption{
            option.WithWechatPayAutoAuthCipher(conf.MchId, conf.MchCertificateSerialNumber, mchPrivateKey, conf.MchAPIv3Key),
        }
        Pay.Client, err = core.NewClient(ctx, opts...)
        if err != nil {
            log.Fatalf("new wechat pay client err:%s", err)
        }

        //注册notifyHandler
        err = downloader.MgrInstance().RegisterDownloaderWithPrivateKey(ctx, mchPrivateKey, conf.MchCertificateSerialNumber, conf.MchId, conf.MchAPIv3Key)
        if err != nil {
            log.Fatalf("RegisterDownloaderWithPrivateKey err:%s", err)
        }
        // 2. 获取商户号对应的微信支付平台证书访问器
        certificateVisitor := downloader.MgrInstance().GetCertificateVisitor(conf.MchId)
        // 3. 使用证书访问器初始化 `notify.Handler`
        Pay.NotifyHandler, err = notify.NewRSANotifyHandler(conf.MchAPIv3Key, verifiers.NewSHA256WithRSAVerifier(certificateVisitor))
        if err != nil {
            log.Fatalf("notify.NewRSANotifyHandler err:%s", err)
        }
        Pay.Config = conf
    })
}
