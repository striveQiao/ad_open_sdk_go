/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaAssetV2ListV2ResponseDataPageInfo 分页信息
type DpaAssetV2ListV2ResponseDataPageInfo struct {
	// 页码
	Page *int64 `json:"page,omitempty"`
	// 页面大小
	PageSize *int64 `json:"page_size,omitempty"`
	// 总数
	TotalNumber *int64 `json:"total_number,omitempty"`
}
