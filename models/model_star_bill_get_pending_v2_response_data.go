/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarBillGetPendingV2ResponseData
type StarBillGetPendingV2ResponseData struct {
	// 可用余额
	Balance *int64 `json:"balance,omitempty"`
	// 每个任务金额明细
	OrderBillList []*StarBillGetPendingV2ResponseDataOrderBillListInner `json:"order_bill_list,omitempty"`
	// 精确可用余额
	PreciseBalance *float64 `json:"precise_balance,omitempty"`
	// 精确总金额（单位元，下同）
	PreciseTotalPay *float64 `json:"precise_total_pay,omitempty"`
	// 精确总服务费
	PreciseTotalPlatformFee *float64 `json:"precise_total_platform_fee,omitempty"`
	// 精确待付金额
	PreciseTotalRemaining *float64 `json:"precise_total_remaining,omitempty"`
	// 精确总任务金额
	PreciseTotalTaskCost *float64 `json:"precise_total_task_cost,omitempty"`
	// 总金额（单位元，下同）
	TotalPay *int64 `json:"total_pay,omitempty"`
	//
	TotalPlatformFee *int64 `json:"total_platform_fee,omitempty"`
	// 待付金额
	TotalRemaining *int64 `json:"total_remaining,omitempty"`
	//
	TotalTaskCost *int64 `json:"total_task_cost,omitempty"`
}
