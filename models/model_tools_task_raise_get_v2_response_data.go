/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsTaskRaiseGetV2ResponseData
type ToolsTaskRaiseGetV2ResponseData struct {
	PageInfo *ToolsTaskRaiseGetV2ResponseDataPageInfo `json:"page_info,omitempty"`
	//
	Reports []*ToolsTaskRaiseGetV2ResponseDataReportsInner `json:"reports,omitempty"`
}
