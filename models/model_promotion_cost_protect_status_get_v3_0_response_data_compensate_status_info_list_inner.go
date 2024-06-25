/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionCostProtectStatusGetV30ResponseDataCompensateStatusInfoListInner struct for PromotionCostProtectStatusGetV30ResponseDataCompensateStatusInfoListInner
type PromotionCostProtectStatusGetV30ResponseDataCompensateStatusInfoListInner struct {
	// 赔付金额
	CompensateAmount *float64                                                                      `json:"compensate_amount,omitempty"`
	CompensateStatus *PromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus `json:"compensate_status,omitempty"`
	// 成本保障结束原因
	EndReasons *map[string]string `json:"end_reasons,omitempty"`
	// 成本保障失效原因
	InvalidReasons *map[string]string `json:"invalid_reasons,omitempty"`
	// 请求计划id
	QueryId *int64                                                              `json:"query_id,omitempty"`
	Status  *PromotionCostProtectStatusGetV30DataCompensateStatusInfoListStatus `json:"status,omitempty"`
	// 赔付规则链接
	Url *string `json:"url,omitempty"`
}
