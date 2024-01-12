/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueSmartphoneGetV2ResponseDataListInner struct for ClueSmartphoneGetV2ResponseDataListInner
type ClueSmartphoneGetV2ResponseDataListInner struct {
	//
	CallDisplay *int64 `json:"call_display,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	InstanceId *int64 `json:"instance_id,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	NoDisturbTimes []*ClueSmartphoneCreateV2RequestNoDisturbTimesInner `json:"no_disturb_times,omitempty"`
	//
	PhoneNumber  *string                                  `json:"phone_number,omitempty"`
	ValidateType *ClueSmartphoneGetV2DataListValidateType `json:"validate_type,omitempty"`
}
