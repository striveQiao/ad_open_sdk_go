/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BusinessPlatformCompanyAccountGetV30ResponseData
type BusinessPlatformCompanyAccountGetV30ResponseData struct {
	//
	AccountList []*BusinessPlatformCompanyAccountGetV30ResponseDataAccountListInner `json:"account_list,omitempty"`
	PageInfo    *BusinessPlatformCompanyAccountGetV30ResponseDataPageInfo           `json:"page_info,omitempty"`
}
