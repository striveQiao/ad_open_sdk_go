/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueGetV2ResponseData
type ToolsClueGetV2ResponseData struct {
	//
	List     []*ToolsClueGetV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *ToolsClueGetV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
