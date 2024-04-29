/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanToolsSmartBoostAdBoostReportGetV10ResponseDataRaiseResultsInner struct for QianchuanToolsSmartBoostAdBoostReportGetV10ResponseDataRaiseResultsInner
type QianchuanToolsSmartBoostAdBoostReportGetV10ResponseDataRaiseResultsInner struct {
	// 广告计划id
	AdId *int64 `json:"ad_id,omitempty"`
	// 千川广告主账户ID
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	// 起量点击次数
	ClickCnt *float64 `json:"click_cnt,omitempty"`
	// 起量转化数
	ConvertCnt *float64 `json:"convert_cnt,omitempty"`
	// 起量转化率
	ConvertRate *float64 `json:"convert_rate,omitempty"`
	// 起量点击率
	Ctr *float64 `json:"ctr,omitempty"`
	// 起量直接成交金额，单位元，小数点后两位
	PayOrderAmountGmv *float64 `json:"pay_order_amount_gmv,omitempty"`
	// 起量直接支付ROI
	PrepayAndPayOrderRoi *float64 `json:"prepay_and_pay_order_roi,omitempty"`
	// 起量展示次数
	ShowCnt *float64 `json:"show_cnt,omitempty"`
	// 起量消耗，单位元，小数点后两位
	StatCost *float64 `json:"stat_cost,omitempty"`
	// 数据起始时间，格式为：yyyy-MM-dd HH:mm:ss
	StatDatetime *string `json:"stat_datetime,omitempty"`
}
