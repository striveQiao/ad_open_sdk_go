/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueSmartPhoneGetV2ResponseData
type ToolsClueSmartPhoneGetV2ResponseData struct {
	//
	AdvKey *string `json:"adv_key,omitempty"`
	//
	List     []*ToolsClueSmartPhoneGetV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *ToolsClueSmartPhoneGetV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
