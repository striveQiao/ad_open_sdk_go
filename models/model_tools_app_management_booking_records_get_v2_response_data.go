/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementBookingRecordsGetV2ResponseData
type ToolsAppManagementBookingRecordsGetV2ResponseData struct {
	//
	List     []*ToolsAppManagementBookingRecordsGetV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo ToolsAppManagementAndroidAppListV2ResponseDataPageInfo        `json:"page_info"`
}
