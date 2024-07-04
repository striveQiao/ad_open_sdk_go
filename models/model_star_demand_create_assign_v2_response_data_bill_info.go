/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandCreateAssignV2ResponseDataBillInfo 账单明细
type StarDemandCreateAssignV2ResponseDataBillInfo struct {
	// 可用余额（单位元，下同）
	Balance *int64 `json:"balance,omitempty"`
	// 每个任务金额明细
	OrderBillList []*StarDemandCreateAssignV2ResponseDataBillInfoOrderBillListInner `json:"order_bill_list,omitempty"`
	// 精确可用余额（单位元，下同）
	PreciseBalance *float64 `json:"precise_balance,omitempty"`
	// 精确总金额
	PreciseTotalAmount *float64 `json:"precise_total_amount,omitempty"`
	// 精确总服务费
	PreciseTotalPlatformFee *float64 `json:"precise_total_platform_fee,omitempty"`
	// 精确待付金额
	PreciseTotalRemaining *float64 `json:"precise_total_remaining,omitempty"`
	// 精确任务总金额
	PreciseTotalTaskCost *float64 `json:"precise_total_task_cost,omitempty"`
	// 总金额
	TotalAmount *int64 `json:"total_amount,omitempty"`
	// 总服务费
	TotalPlatformFee *int64 `json:"total_platform_fee,omitempty"`
	// 待付金额
	TotalRemaining *int64 `json:"total_remaining,omitempty"`
	// 任务总金额
	TotalTaskCost *int64 `json:"total_task_cost,omitempty"`
}
