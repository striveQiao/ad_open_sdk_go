/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueSmartphoneRecordV2ResponseData
type ClueSmartphoneRecordV2ResponseData struct {
	//
	List     []*ClueSmartphoneRecordV2ResponseDataListInner `json:"list[],omitempty"`
	PageInfo *ClueSmartphoneRecordV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
