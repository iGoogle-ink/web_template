package pay

import (
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
	Redis  *redis.ClusterClient
	WeChat *wechat.Client
	AliPay *alipay.Client
	QQ     *qq.Client
}

func New(c *conf.Config, db *xorm.Engine, rds *redis.ClusterClient) (d *Dao) {
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
	// todo: 设置证书
	//err := client.AddCertFilePath(c.CertFilePath, c.KeyFilePath, c.Pkcs12FilePath)
	//if err != nil {
	//	panic(err)
	//}
	return client
}

func initAliPay(c *conf.AliPay) (client *alipay.Client) {
	client = alipay.NewClient(c.Appid, c.PrivateKey, c.IsProd)
	client.SetNotifyUrl("https://www.igoogle.ink")
	// todo: 设置证书
	//err := client.SetCertSnByPath(c.AppCertPath, c.RootCertPath, c.PublicCertPath)
	//if err != nil {
	//	panic(err)
	//}
	return client
}

func initQQ(c *conf.QQPay) (client *qq.Client) {
	client = qq.NewClient(c.MchId, c.ApiKey)
	// todo: 设置证书
	//err := client.AddCertFilePath(c.CertFilePath, c.KeyFilePath, c.Pkcs12FilePath)
	//if err != nil {
	//	panic(err)
	//}
	return client
}
