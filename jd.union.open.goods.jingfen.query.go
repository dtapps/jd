package jd

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
)

type UnionOpenGoodsJIngFenQueryResultResponse struct {
	JdUnionOpenGoodsJingfenQueryResponce struct {
		Code        string `json:"code"`
		QueryResult string `json:"queryResult"`
	} `json:"jd_union_open_goods_jingfen_query_responce"`
}

type UnionOpenGoodsJIngFenQueryQueryResult struct {
	Code int `json:"code"`
	Data []struct {
		BrandCode    string `json:"brandCode"`
		BrandName    string `json:"brandName"`
		CategoryInfo struct {
			Cid1     int64  `json:"cid1"`
			Cid1Name string `json:"cid1Name"`
			Cid2     int    `json:"cid2"`
			Cid2Name string `json:"cid2Name"`
			Cid3     int    `json:"cid3"`
			Cid3Name string `json:"cid3Name"`
		} `json:"categoryInfo"`
		Comments       int `json:"comments"`
		CommissionInfo struct {
			Commission          float64 `json:"commission"`
			CommissionShare     float64 `json:"commissionShare"`
			CouponCommission    float64 `json:"couponCommission"`
			EndTime             int64   `json:"endTime"`
			IsLock              int     `json:"isLock"`
			PlusCommissionShare float64 `json:"plusCommissionShare"`
			StartTime           int64   `json:"startTime"`
		} `json:"commissionInfo"`
		CouponInfo struct {
			CouponList []struct {
				BindType     int     `json:"bindType"`
				Discount     float64 `json:"discount"`
				GetEndTime   int64   `json:"getEndTime"`
				GetStartTime int64   `json:"getStartTime"`
				HotValue     int     `json:"hotValue,omitempty"`
				IsBest       int     `json:"isBest"`
				Link         string  `json:"link"`
				PlatformType int     `json:"platformType"`
				Quota        float64 `json:"quota"`
				UseEndTime   int64   `json:"useEndTime"`
				UseStartTime int64   `json:"useStartTime"`
			} `json:"couponList"`
		} `json:"couponInfo"`
		DeliveryType      int     `json:"deliveryType"`
		ForbidTypes       []int   `json:"forbidTypes"`
		GoodCommentsShare float64 `json:"goodCommentsShare"`
		ImageInfo         struct {
			ImageList []struct {
				Url string `json:"url"`
			} `json:"imageList"`
			WhiteImage string `json:"whiteImage,omitempty"`
		} `json:"imageInfo"`
		InOrderCount30Days    int64  `json:"inOrderCount30Days"`
		InOrderCount30DaysSku int    `json:"inOrderCount30DaysSku"`
		IsHot                 int    `json:"isHot"`
		JxFlags               []int  `json:"jxFlags,omitempty"`
		MaterialUrl           string `json:"materialUrl"`
		Owner                 string `json:"owner"`
		PinGouInfo            struct {
			PingouEndTime   int64   `json:"pingouEndTime,omitempty"`
			PingouPrice     float64 `json:"pingouPrice,omitempty"`
			PingouStartTime int64   `json:"pingouStartTime,omitempty"`
			PingouTmCount   int     `json:"pingouTmCount,omitempty"`
			PingouUrl       string  `json:"pingouUrl,omitempty"`
		} `json:"pinGouInfo"`
		PriceInfo struct {
			HistoryPriceDay   int     `json:"historyPriceDay"`
			LowestCouponPrice float64 `json:"lowestCouponPrice"`
			LowestPrice       float64 `json:"lowestPrice"`
			LowestPriceType   int     `json:"lowestPriceType"`
			Price             float64 `json:"price"`
		} `json:"priceInfo"`
		ResourceInfo struct {
			EliteId   int    `json:"eliteId"`
			EliteName string `json:"eliteName"`
		} `json:"resourceInfo"`
		ShopInfo struct {
			ShopId                        int64   `json:"shopId"`
			ShopLabel                     string  `json:"shopLabel"`
			ShopLevel                     float64 `json:"shopLevel"`
			ShopName                      string  `json:"shopName"`
			AfsFactorScoreRankGrade       string  `json:"afsFactorScoreRankGrade,omitempty"`
			AfterServiceScore             string  `json:"afterServiceScore,omitempty"`
			CommentFactorScoreRankGrade   string  `json:"commentFactorScoreRankGrade,omitempty"`
			LogisticsFactorScoreRankGrade string  `json:"logisticsFactorScoreRankGrade,omitempty"`
			LogisticsLvyueScore           string  `json:"logisticsLvyueScore,omitempty"`
			ScoreRankRate                 string  `json:"scoreRankRate,omitempty"`
			UserEvaluateScore             string  `json:"userEvaluateScore,omitempty"`
		} `json:"shopInfo"`
		SkuId        int64 `json:"skuId"`
		SkuLabelInfo struct {
			FxgServiceList []interface{} `json:"fxgServiceList"`
			Is7ToReturn    int           `json:"is7ToReturn"`
		} `json:"skuLabelInfo"`
		SkuName string `json:"skuName"`
		Spuid   int64  `json:"spuid"`
	} `json:"data"`
	Message    string `json:"message"`
	TotalCount int64  `json:"totalCount"`
}

type UnionOpenGoodsJIngFenQueryResult struct {
	Responce UnionOpenGoodsJIngFenQueryResultResponse // 结果
	Result   UnionOpenGoodsJIngFenQueryQueryResult    // 结果
	Body     []byte                                   // 内容
	Http     gorequest.Response                       // 请求
}

func newUnionOpenGoodsJIngFenQueryResult(responce UnionOpenGoodsJIngFenQueryResultResponse, result UnionOpenGoodsJIngFenQueryQueryResult, body []byte, http gorequest.Response) *UnionOpenGoodsJIngFenQueryResult {
	return &UnionOpenGoodsJIngFenQueryResult{Responce: responce, Result: result, Body: body, Http: http}
}

// UnionOpenGoodsJIngFenQuery 京粉精选商品查询接口
// https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.goods.jingfen.query
func (c *Client) UnionOpenGoodsJIngFenQuery(ctx context.Context, notMustParams ...gorequest.Params) (*UnionOpenGoodsJIngFenQueryResult, error) {
	// 参数
	params := NewParamsWithType("jd.union.open.goods.jingfen.query", notMustParams...)
	// 请求
	request, err := c.request(ctx, params)
	if err != nil {
		return newUnionOpenGoodsJIngFenQueryResult(UnionOpenGoodsJIngFenQueryResultResponse{}, UnionOpenGoodsJIngFenQueryQueryResult{}, request.ResponseBody, request), err
	}
	// 定义
	var responce UnionOpenGoodsJIngFenQueryResultResponse
	var result UnionOpenGoodsJIngFenQueryQueryResult
	err = gojson.Unmarshal(request.ResponseBody, &responce)
	err = gojson.Unmarshal([]byte(responce.JdUnionOpenGoodsJingfenQueryResponce.QueryResult), &result)
	return newUnionOpenGoodsJIngFenQueryResult(responce, result, request.ResponseBody, request), err
}
