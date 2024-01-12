/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAudienceListGetV10ResponseDataAudiencesListInner struct for QianchuanAudienceListGetV10ResponseDataAudiencesListInner
type QianchuanAudienceListGetV10ResponseDataAudiencesListInner struct {
	//
	AudienceGroup *string `json:"audience_group,omitempty"`
	//
	AudienceId *int64 `json:"audience_id,omitempty"`
	//
	AudienceName   *string                                                     `json:"audience_name,omitempty"`
	AudienceSource *QianchuanAudienceListGetV10DataAudiencesListAudienceSource `json:"audience_source,omitempty"`
	AudienceType   *QianchuanAudienceListGetV10DataAudiencesListAudienceType   `json:"audience_type,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	PushProduct *string `json:"push_product,omitempty"`
	//
	Status *int32 `json:"status,omitempty"`
	//
	StockNum *int64 `json:"stock_num,omitempty"`
	//
	StockStatus *int32 `json:"stock_status,omitempty"`
}
