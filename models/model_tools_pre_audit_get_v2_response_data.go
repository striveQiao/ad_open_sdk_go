/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPreAuditGetV2ResponseData
type ToolsPreAuditGetV2ResponseData struct {
	// 预审结果列表
	List     []*ToolsPreAuditGetV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *ToolsPreAuditGetV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
