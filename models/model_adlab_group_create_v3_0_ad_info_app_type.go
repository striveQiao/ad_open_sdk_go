/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupCreateV30AdInfoAppType
type AdlabGroupCreateV30AdInfoAppType string

// List of adlab_group_create_v3.0_ad_info_app_type
const (
	APP_ANDROID_AdlabGroupCreateV30AdInfoAppType AdlabGroupCreateV30AdInfoAppType = "APP_ANDROID"
	APP_IOS_AdlabGroupCreateV30AdInfoAppType     AdlabGroupCreateV30AdInfoAppType = "APP_IOS"
)

// All allowed values of AdlabGroupCreateV30AdInfoAppType enum
var AllowedAdlabGroupCreateV30AdInfoAppTypeEnumValues = []AdlabGroupCreateV30AdInfoAppType{
	"APP_ANDROID",
	"APP_IOS",
}

// NewAdlabGroupCreateV30AdInfoAppTypeFromValue returns a pointer to a valid AdlabGroupCreateV30AdInfoAppType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupCreateV30AdInfoAppTypeFromValue(v string) (*AdlabGroupCreateV30AdInfoAppType, error) {
	ev := AdlabGroupCreateV30AdInfoAppType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupCreateV30AdInfoAppType: valid values are %v", v, AllowedAdlabGroupCreateV30AdInfoAppTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupCreateV30AdInfoAppType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupCreateV30AdInfoAppTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_create_v3.0_ad_info_app_type value
func (v AdlabGroupCreateV30AdInfoAppType) Ptr() *AdlabGroupCreateV30AdInfoAppType {
	return &v
}
