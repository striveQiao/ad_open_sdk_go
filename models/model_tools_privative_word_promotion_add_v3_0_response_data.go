/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordPromotionAddV30ResponseData
type ToolsPrivativeWordPromotionAddV30ResponseData struct {
	//
	ErrorList []*ToolsPrivativeWordPromotionAddV30ResponseDataErrorListInner `json:"error_list,omitempty"`
	// 添加失败的广告id列表
	PromotionErrorList []int64 `json:"promotion_error_list,omitempty"`
	// 添加成功的广告列表
	PromotionList []map[string]interface{} `json:"promotion_list,omitempty"`
}
