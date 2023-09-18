/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsQuotaGetV2ResponseData
type ToolsQuotaGetV2ResponseData struct {
	// 是否在学习期，学习期即初始期，详细规则请进入巨量广告投放平台-工具-在投计划配额界面查看
	InLearning *bool `json:"in_learning,omitempty"`
	// 最高月消耗
	MaxCost *float64 `json:"max_cost,omitempty"`
	// 在投计划配额
	SumQuota *int64 `json:"sum_quota,omitempty"`
	// 在投计划数
	UsedQuota *int64 `json:"used_quota,omitempty"`
}
