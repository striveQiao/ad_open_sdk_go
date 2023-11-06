/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdQuotaGetV10ResponseDataQuotaFeedDeliveryInfo 当前所在投放阶段信息
type QianchuanAdQuotaGetV10ResponseDataQuotaFeedDeliveryInfo struct {
	// 托管在投计划数
	AdlabNum *int64 `json:"adlab_num,omitempty"`
	// 非托管在投计划数
	NoAdlabNum *int64 `json:"no_adlab_num,omitempty"`
	// 总在投计划数
	TotalNum *int64 `json:"total_num,omitempty"`
}
