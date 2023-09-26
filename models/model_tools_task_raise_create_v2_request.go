/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsTaskRaiseCreateV2Request struct for ToolsTaskRaiseCreateV2Request
type ToolsTaskRaiseCreateV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	// 日预算金额，当budget_mode为LIMIT时必填，范围：1000-9999999
	AllocatedBudget *int64                           `json:"allocated_budget,omitempty"`
	BudgetMode      ToolsTaskRaiseCreateV2BudgetMode `json:"budget_mode"`
	// 起量任务结束时间yyyy-mm-dd，传空为不限时长，起量将在指定日期0点结束
	EndTime         *string                                `json:"end_time,omitempty"`
	PlatformVersion *ToolsTaskRaiseCreateV2PlatformVersion `json:"platform_version,omitempty"`
	RaiseMode       *ToolsTaskRaiseCreateV2RaiseMode       `json:"raise_mode,omitempty"`
}
