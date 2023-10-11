/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupDetailV30DataDataAdInfoAudiencePlatform
type AdlabGroupDetailV30DataDataAdInfoAudiencePlatform string

// List of adlab_group_detail_v3.0_data_data_ad_info_audience_platform
const (
	ANDROID_AdlabGroupDetailV30DataDataAdInfoAudiencePlatform AdlabGroupDetailV30DataDataAdInfoAudiencePlatform = "ANDROID"
	IOS_AdlabGroupDetailV30DataDataAdInfoAudiencePlatform     AdlabGroupDetailV30DataDataAdInfoAudiencePlatform = "IOS"
)

// All allowed values of AdlabGroupDetailV30DataDataAdInfoAudiencePlatform enum
var AllowedAdlabGroupDetailV30DataDataAdInfoAudiencePlatformEnumValues = []AdlabGroupDetailV30DataDataAdInfoAudiencePlatform{
	"ANDROID",
	"IOS",
}

// NewAdlabGroupDetailV30DataDataAdInfoAudiencePlatformFromValue returns a pointer to a valid AdlabGroupDetailV30DataDataAdInfoAudiencePlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupDetailV30DataDataAdInfoAudiencePlatformFromValue(v string) (*AdlabGroupDetailV30DataDataAdInfoAudiencePlatform, error) {
	ev := AdlabGroupDetailV30DataDataAdInfoAudiencePlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupDetailV30DataDataAdInfoAudiencePlatform: valid values are %v", v, AllowedAdlabGroupDetailV30DataDataAdInfoAudiencePlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupDetailV30DataDataAdInfoAudiencePlatform) IsValid() bool {
	for _, existing := range AllowedAdlabGroupDetailV30DataDataAdInfoAudiencePlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_detail_v3.0_data_data_ad_info_audience_platform value
func (v AdlabGroupDetailV30DataDataAdInfoAudiencePlatform) Ptr() *AdlabGroupDetailV30DataDataAdInfoAudiencePlatform {
	return &v
}
