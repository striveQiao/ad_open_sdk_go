/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueSmartphoneGetV2ResponseData
type ClueSmartphoneGetV2ResponseData struct {
	//
	Key *string `json:"key,omitempty"`
	//
	List     []*ClueSmartphoneGetV2ResponseDataListInner  `json:"list,omitempty"`
	PageInfo *AgentAdvertiserSelectV2ResponseDataPageInfo `json:"page_info,omitempty"`
}
