/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAbTestListGetV2ResponseData
type ToolsAbTestListGetV2ResponseData struct {
	PageInfo *ToolsAbTestListGetV2ResponseDataPageInfo `json:"page_info,omitempty"`
	// 实验列表
	TestList []*ToolsAbTestListGetV2ResponseDataTestListInner `json:"test_list"`
}
