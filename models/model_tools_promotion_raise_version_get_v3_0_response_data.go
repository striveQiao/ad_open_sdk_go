/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPromotionRaiseVersionGetV30ResponseData
type ToolsPromotionRaiseVersionGetV30ResponseData struct {
	PageInfo *ToolsPromotionRaiseVersionGetV30ResponseDataPageInfo `json:"page_info,omitempty"`
	// 起量版本列表
	RaiseVersionList []*ToolsPromotionRaiseVersionGetV30ResponseDataRaiseVersionListInner `json:"raise_version_list,omitempty"`
}
