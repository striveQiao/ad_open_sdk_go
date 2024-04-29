/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandCreateChallengeV2ResponseDataBillInfo 账单明细
type StarDemandCreateChallengeV2ResponseDataBillInfo struct {
	// （任务发布后）可用余额（单位元，下同）
	Balance *int64 `json:"balance,omitempty"`
	// 总金额（预算+服务费）
	TotalAmount *int64 `json:"total_amount,omitempty"`
	// 服务费
	TotalPlatformFee *int64 `json:"total_platform_fee,omitempty"`
	// 任务金额（预算）
	TotalTaskCost *int64 `json:"total_task_cost,omitempty"`
}
