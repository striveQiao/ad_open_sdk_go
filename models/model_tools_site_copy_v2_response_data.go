/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteCopyV2ResponseData
type ToolsSiteCopyV2ResponseData struct {
	// 复制失败列表，整体成功不返回该列表
	ErrorList []*ToolsSiteCopyV2ResponseDataErrorListInner `json:"error_list,omitempty"`
	// 复制成功列表，整体失败不返回该列表
	SuccessList []*ToolsSiteCopyV2ResponseDataSuccessListInner `json:"success_list,omitempty"`
}
