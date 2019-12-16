package pay

import (
	"fmt"

	"web_template/conf"

	"github.com/go-redis/redis/v7"
	"github.com/iGoogle-ink/gopay/v2/alipay"
	"github.com/iGoogle-ink/gopay/v2/qq"
	"github.com/iGoogle-ink/gopay/v2/wechat"
	"xorm.io/xorm"
)

type Dao struct {
	c      *conf.Config
	DB     *xorm.Engine
	Redis  *redis.Client
	WeChat *wechat.Client
	AliPay *alipay.Client
	QQ     *qq.Client
}

func New(c *conf.Config, db *xorm.Engine, rds *redis.Client) (d *Dao) {
	fmt.Println("c.WeChatPay:", c.WeChatPay)
	d = &Dao{
		c:      c,
		DB:     db,
		Redis:  rds,
		WeChat: initWeChatPay(c.WeChatPay),
		AliPay: initAliPay(c.AliPay),
		QQ:     initQQ(c.QQPay),
	}
	return d
}

func initWeChatPay(c *conf.WeChatPay) (client *wechat.Client) {
	client = wechat.NewClient(c.Appid, c.MchId, c.ApiKey, c.IsProd)
	err := client.AddCertFilePath(c.CertFilePath, c.KeyFilePath, c.Pkcs12FilePath)
	if err != nil {
		// todo: gopay 改掉此地方，client不存 []byte 数组，直接存 tls.Certificate，在 AddCertFilePath() 里直接解析好
		panic(err)
	}
	return client
}

func initAliPay(c *conf.AliPay) (client *alipay.Client) {
	client = alipay.NewClient(c.Appid, c.PrivateKey, c.IsProd)
	client.SetNotifyUrl("https://www.igoogle.ink")
	err := client.SetCertSnByPath(c.AppCertPath, c.RootCertPath, c.PublicCertPath)
	if err != nil {
		panic(err)
	}
	return client
}

func initQQ(c *conf.QQPay) (client *qq.Client) {
	client = qq.NewClient(c.MchId, c.ApiKey)
	err := client.AddCertFilePath(c.CertFilePath, c.KeyFilePath, c.Pkcs12FilePath)
	if err != nil {
		// todo: gopay 改掉此地方，client不存 []byte 数组，直接存 tls.Certificate，在 AddCertFilePath() 里直接解析好
		panic(err)
	}
	return client
}
