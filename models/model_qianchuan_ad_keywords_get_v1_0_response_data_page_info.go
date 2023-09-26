/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdKeywordsGetV10ResponseDataPageInfo
type QianchuanAdKeywordsGetV10ResponseDataPageInfo struct {
	// 页码
	Page *int32 `json:"page,omitempty"`
	// 每页数目
	PageSize *int32 `json:"page_size,omitempty"`
	// 关键词总数
	TotalNum *int32 `json:"total_num,omitempty"`
	// 页面总数
	TotalPage *int32 `json:"total_page,omitempty"`
}
