/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementAppGetV2ResponseData
type ToolsAppManagementAppGetV2ResponseData struct {
	//
	List     []*ToolsAppManagementAppGetV2ResponseDataListInner `json:"list"`
	PageInfo *ToolsAppManagementAppGetV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
