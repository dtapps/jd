package jd

import (
	"encoding/json"
	"go.dtapp.net/gorequest"
)

type UnionOpenPromotionBySubUnionIdGetResultResponse struct {
	JdUnionOpenPromotionBySubUnionIdGetResponce struct {
		Code      string `json:"code"`
		GetResult string `json:"getResult"`
	} `json:"jd_union_open_promotion_common_get_responce"`
}

type UnionOpenPromotionBySubUnionIdGetGetResult struct {
	Code int `json:"code"`
	Data struct {
		ClickURL string `json:"clickURL"`
		JCommand string `json:"jCommand"`
	} `json:"data"`
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
}

type UnionOpenPromotionBySubUnionIdGetResult struct {
	Responce UnionOpenPromotionBySubUnionIdGetResultResponse // 结果
	Result   UnionOpenPromotionBySubUnionIdGetGetResult      // 结果
	Body     []byte                                          // 内容
	Http     gorequest.Response                              // 请求
	Err      error                                           // 错误
}

func NewUnionOpenPromotionBySubUnionIdGetResult(responce UnionOpenPromotionBySubUnionIdGetResultResponse, result UnionOpenPromotionBySubUnionIdGetGetResult, body []byte, http gorequest.Response, err error) *UnionOpenPromotionBySubUnionIdGetResult {
	return &UnionOpenPromotionBySubUnionIdGetResult{Responce: responce, Result: result, Body: body, Http: http, Err: err}
}

// UnionOpenPromotionBySubUnionIdGet 社交媒体获取推广链接接口【申请】
// https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.promotion.bysubunionid.get
func (app *App) UnionOpenPromotionBySubUnionIdGet(notMustParams ...Params) *UnionOpenPromotionBySubUnionIdGetResult {
	// 参数
	params := NewParamsWithType("jd.union.open.promotion.bysubunionid.get", notMustParams...)
	// 请求
	request, err := app.request(params)
	// 定义
	var responce UnionOpenPromotionBySubUnionIdGetResultResponse
	var result UnionOpenPromotionBySubUnionIdGetGetResult
	err = json.Unmarshal(request.ResponseBody, &responce)
	err = json.Unmarshal([]byte(responce.JdUnionOpenPromotionBySubUnionIdGetResponce.GetResult), &result)
	return NewUnionOpenPromotionBySubUnionIdGetResult(responce, result, request.ResponseBody, request, err)
}
