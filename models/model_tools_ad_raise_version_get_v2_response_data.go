/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAdRaiseVersionGetV2ResponseData
type ToolsAdRaiseVersionGetV2ResponseData struct {
	//
	AdRaiseVersionList []*ToolsAdRaiseVersionGetV2ResponseDataAdRaiseVersionListInner `json:"ad_raise_version_list,omitempty"`
	PageInfo           *ToolsAdRaiseVersionGetV2ResponseDataPageInfo                  `json:"page_info,omitempty"`
}
