/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupUpdateBudgetV30ResponseDataDataInner struct for AdlabGroupUpdateBudgetV30ResponseDataDataInner
type AdlabGroupUpdateBudgetV30ResponseDataDataInner struct {
	// 更新失败的项目id信息
	Errors []*AdlabGroupUpdateBudgetV30ResponseDataDataInnerErrorsInner `json:"errors,omitempty"`
	// 更新成功的项目id列表
	Ids []int64 `json:"ids,omitempty"`
}
