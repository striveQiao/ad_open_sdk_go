/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserFundTransactionGetV2ResponseData
type AdvertiserFundTransactionGetV2ResponseData struct {
	//
	List     []*AdvertiserFundTransactionGetV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *AdvertiserFundTransactionGetV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
