/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DmpCustomAudienceReadV2ResponseDataCustomAudienceListInner struct for DmpCustomAudienceReadV2ResponseDataCustomAudienceListInner
type DmpCustomAudienceReadV2ResponseDataCustomAudienceListInner struct {
	//
	CoverNum *int64 `json:"cover_num,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	CustomAudienceId *int64 `json:"custom_audience_id,omitempty"`
	//
	DataSourceId   *string                                                      `json:"data_source_id,omitempty"`
	DeliveryStatus *DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus `json:"delivery_status,omitempty"`
	//
	ExistPullOffTag *int64 `json:"exist_pull_off_tag,omitempty"`
	//
	Isdel *int64 `json:"isdel,omitempty"`
	//
	ModifyTime *string `json:"modify_time,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	PushStatus *int64                                               `json:"push_status,omitempty"`
	Source     *DmpCustomAudienceReadV2DataCustomAudienceListSource `json:"source,omitempty"`
	//
	Status *int64 `json:"status,omitempty"`
	//
	Tag *string `json:"tag,omitempty"`
	//
	ThirdPartyInfo *string `json:"third_party_info,omitempty"`
	//
	UploadNum *int64 `json:"upload_num,omitempty"`
}
