/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanToolsSmartBoostAdBoostSetV10Request struct for QianchuanToolsSmartBoostAdBoostSetV10Request
type QianchuanToolsSmartBoostAdBoostSetV10Request struct {
	// 广告计划ID
	AdId int64 `json:"ad_id"`
	// 千川广告主账户ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 起量预算，单位元，小数点后两位 注意：不得超过计划日预算/总预算 注意：当opt_type=START_RAISE、SUBSCRIBE_RAISE时必填
	Budget  *float64                                     `json:"budget,omitempty"`
	OptType QianchuanToolsSmartBoostAdBoostSetV10OptType `json:"opt_type"`
	// 起量时长，单位小时，小数点后1位，范围0.5-6小时 注意：当opt_type=START_RAISE、SUBSCRIBE_RAISE时必填
	RaiseDuration *float64 `json:"raise_duration,omitempty"`
	// 起量时间，仅支持当前时间往后的时间 注意：当opt_type=START_RAISE/SUBSCRIBE_RAISE必填
	RaiseTime *string `json:"raise_time,omitempty"`
}
