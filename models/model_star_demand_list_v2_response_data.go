/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandListV2ResponseData
type StarDemandListV2ResponseData struct {
	//
	List     []*StarDemandListV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *StarDemandListV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
