/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanProductAnalyseListV10ResponseDataPageInfo
type QianchuanProductAnalyseListV10ResponseDataPageInfo struct {
	//
	Page int64 `json:"page"`
	//
	PageSize int64 `json:"page_size"`
	//
	TotalNumber *int64 `json:"total_number,omitempty"`
	//
	TotalPage *int64 `json:"total_page,omitempty"`
}
