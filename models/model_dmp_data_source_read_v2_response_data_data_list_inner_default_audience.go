/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DmpDataSourceReadV2ResponseDataDataListInnerDefaultAudience
type DmpDataSourceReadV2ResponseDataDataListInnerDefaultAudience struct {
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	CalculateSubType *int64 `json:"calculate_sub_type,omitempty"`
	//
	CalculateType *int64 `json:"calculate_type,omitempty"`
	//
	CoverNum *int64 `json:"cover_num,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	CustomAudienceId *int64 `json:"custom_audience_id,omitempty"`
	//
	CustomType *int64 `json:"custom_type,omitempty"`
	//
	DataSourceId   *string                                                       `json:"data_source_id,omitempty"`
	DeliveryStatus *DmpDataSourceReadV2DataDataListDefaultAudienceDeliveryStatus `json:"delivery_status,omitempty"`
	//
	ExpiryDate *string `json:"expiry_date,omitempty"`
	//
	Isdel *int64 `json:"isdel,omitempty"`
	//
	ModifyTime *string `json:"modify_time,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	PushStatus *int64 `json:"push_status,omitempty"`
	//
	Source *int64 `json:"source,omitempty"`
	//
	Status *int64 `json:"status,omitempty"`
	//
	Tag *string `json:"tag,omitempty"`
	//
	ThirdPartyInfo *string `json:"third_party_info,omitempty"`
	//
	UploadNum *int64 `json:"upload_num,omitempty"`
}
