/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandListV2ResponseData
type StarDemandListV2ResponseData struct {
	// 任务列表
	List     []*StarDemandListV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *StarDemandListV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
