/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAudienceListGetV10ResponseData
type QianchuanAudienceListGetV10ResponseData struct {
	//
	AudiencesList []*QianchuanAudienceListGetV10ResponseDataAudiencesListInner `json:"audiences_list,omitempty"`
	PageInfo      *QianchuanAudienceListGetV10ResponseDataPageInfo             `json:"page_info,omitempty"`
}
