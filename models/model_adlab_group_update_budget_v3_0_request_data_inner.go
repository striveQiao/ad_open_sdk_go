/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupUpdateBudgetV30RequestDataInner struct for AdlabGroupUpdateBudgetV30RequestDataInner
type AdlabGroupUpdateBudgetV30RequestDataInner struct {
	// 项目日预算，范围: 300 <= budget <= 9999999.99
	Budget float64 `json:"budget"`
	// 项目id
	Id int64 `json:"id"`
}
