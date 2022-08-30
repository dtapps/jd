package jd

import (
	"go.dtapp.net/dorm"
	"go.dtapp.net/golog"
	"go.dtapp.net/gorequest"
)

// client *dorm.GormClient
type gormClientFun func() *dorm.GormClient

// client *dorm.MongoClient
// databaseName string
type mongoClientFun func() (*dorm.MongoClient, string)

// ClientConfig 实例配置
type ClientConfig struct {
	AppKey         string         // 应用Key
	SecretKey      string         // 密钥
	SiteId         string         // 网站ID/APP ID
	PositionId     string         // 推广位id
	GormClientFun  gormClientFun  // 日志配置
	MongoClientFun mongoClientFun // 日志配置
	Debug          bool           // 日志开关
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	config        struct {
		appKey     string // 应用Key
		secretKey  string // 密钥
		siteId     string // 网站ID/APP ID
		positionId string // 推广位id
	}
	log struct {
		gorm           bool              // 日志开关
		gormClient     *dorm.GormClient  // 日志数据库
		logGormClient  *golog.ApiClient  // 日志服务
		mongo          bool              // 日志开关
		mongoClient    *dorm.MongoClient // 日志数据库
		logMongoClient *golog.ApiClient  // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	var err error
	c := &Client{}

	c.config.appKey = config.AppKey
	c.config.secretKey = config.SecretKey
	c.config.siteId = config.SiteId
	c.config.positionId = config.PositionId

	c.requestClient = gorequest.NewHttp()
	c.requestClient.Uri = apiUrl

	gormClient := config.GormClientFun()
	if gormClient != nil && gormClient.Db != nil {
		c.log.logGormClient, err = golog.NewApiGormClient(func() (*dorm.GormClient, string) {
			return gormClient, logTable
		}, config.Debug)
		if err != nil {
			return nil, err
		}
		c.log.gorm = true
		c.log.gormClient = gormClient
	}

	mongoClient, databaseName := config.MongoClientFun()
	if mongoClient != nil && mongoClient.Db != nil {
		c.log.logMongoClient, err = golog.NewApiMongoClient(func() (*dorm.MongoClient, string, string) {
			return mongoClient, databaseName, logTable
		}, config.Debug)
		if err != nil {
			return nil, err
		}
		c.log.mongo = true
		c.log.mongoClient = mongoClient
	}

	return c, nil
}

// GoodsPriceToInt64 商品券后价
func (c *Client) GoodsPriceToInt64(LowestCouponPrice float64) int64 {
	return int64(LowestCouponPrice * 100)
}

// GoodsOriginalPriceToInt64 商品原价
func (c *Client) GoodsOriginalPriceToInt64(Price float64) int64 {
	return int64(Price * 100)
}

// CouponProportionToInt64 佣金比率
func (c *Client) CouponProportionToInt64(CommissionShare float64) int64 {
	return int64(CommissionShare * 10)
}

// CouponAmountToInt64 优惠券金额
func (c *Client) CouponAmountToInt64(Commission float64) int64 {
	return int64(Commission * 100)
}

// CommissionIntegralToInt64 佣金积分
func (c *Client) CommissionIntegralToInt64(GoodsPrice, CouponProportion int64) int64 {
	return (GoodsPrice * CouponProportion) / 1000
}
