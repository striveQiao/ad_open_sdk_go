/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BudgetGroupDeleteV30Request struct for BudgetGroupDeleteV30Request
type BudgetGroupDeleteV30Request struct {
	// 广告主id
	AdvertiserId int64 `json:"advertiser_id"`
	// 要删除的预算组id列表
	BudgetGroupIds []int64 `json:"budget_group_ids"`
}
