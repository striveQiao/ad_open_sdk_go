/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BusinessPlatformPartnerOrganizationListV2ResponseData
type BusinessPlatformPartnerOrganizationListV2ResponseData struct {
	// 应用列表
	List     []*BusinessPlatformPartnerOrganizationListV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *BusinessPlatformPartnerOrganizationListV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
