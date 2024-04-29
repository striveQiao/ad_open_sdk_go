/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueRobotScriptQueryV2ResponseDataPage 分页信息
type ToolsClueRobotScriptQueryV2ResponseDataPage struct {
	// 页面编号
	PageNumber *int64 `json:"page_number,omitempty"`
	// 页面大小
	PageSize *int64 `json:"page_size,omitempty"`
	// 总页数
	PageTotal *int64 `json:"page_total,omitempty"`
	// 数据总数
	Total *int64 `json:"total,omitempty"`
}
