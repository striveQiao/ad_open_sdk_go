/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueFormGetV2ResponseDataListInner struct for ToolsClueFormGetV2ResponseDataListInner
type ToolsClueFormGetV2ResponseDataListInner struct {
	//
	AdvertiserId *int64                                  `json:"advertiser_id,omitempty"`
	ContainPhone *ToolsClueFormGetV2DataListContainPhone `json:"contain_phone,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	EnableLayer *bool                               `json:"enable_layer,omitempty"`
	FormType    *ToolsClueFormGetV2DataListFormType `json:"form_type,omitempty"`
	//
	InstanceId *int64 `json:"instance_id,omitempty"`
	//
	IsDel *int64 `json:"is_del,omitempty"`
	//
	LightingPageUrl *string `json:"lighting_page_url,omitempty"`
	//
	Name    *string                            `json:"name,omitempty"`
	SubType *ToolsClueFormGetV2DataListSubType `json:"sub_type,omitempty"`
	//
	Version *int64 `json:"version,omitempty"`
}
