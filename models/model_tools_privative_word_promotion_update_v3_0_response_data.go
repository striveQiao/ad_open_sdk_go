/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordPromotionUpdateV30ResponseData
type ToolsPrivativeWordPromotionUpdateV30ResponseData struct {
	//
	ErrorList []*ToolsPrivativeWordPromotionAddV30ResponseDataErrorListInner `json:"error_list,omitempty"`
	//
	PromotionErrorList []int64 `json:"promotion_error_list,omitempty"`
	//
	PromotionList []map[string]interface{} `json:"promotion_list,omitempty"`
}
