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

// AdlabGroupDetailV30DataDataAdInfoAdTarget
type AdlabGroupDetailV30DataDataAdInfoAdTarget string

// List of adlab_group_detail_v3.0_data_data_ad_info_ad_target
const (
	SMART_BID_CONSERVATIVE_AdlabGroupDetailV30DataDataAdInfoAdTarget AdlabGroupDetailV30DataDataAdInfoAdTarget = "SMART_BID_CONSERVATIVE"
	SMART_BID_CUSTOM_AdlabGroupDetailV30DataDataAdInfoAdTarget       AdlabGroupDetailV30DataDataAdInfoAdTarget = "SMART_BID_CUSTOM"
)

// All allowed values of AdlabGroupDetailV30DataDataAdInfoAdTarget enum
var AllowedAdlabGroupDetailV30DataDataAdInfoAdTargetEnumValues = []AdlabGroupDetailV30DataDataAdInfoAdTarget{
	"SMART_BID_CONSERVATIVE",
	"SMART_BID_CUSTOM",
}

// NewAdlabGroupDetailV30DataDataAdInfoAdTargetFromValue returns a pointer to a valid AdlabGroupDetailV30DataDataAdInfoAdTarget
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupDetailV30DataDataAdInfoAdTargetFromValue(v string) (*AdlabGroupDetailV30DataDataAdInfoAdTarget, error) {
	ev := AdlabGroupDetailV30DataDataAdInfoAdTarget(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupDetailV30DataDataAdInfoAdTarget: valid values are %v", v, AllowedAdlabGroupDetailV30DataDataAdInfoAdTargetEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupDetailV30DataDataAdInfoAdTarget) IsValid() bool {
	for _, existing := range AllowedAdlabGroupDetailV30DataDataAdInfoAdTargetEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_detail_v3.0_data_data_ad_info_ad_target value
func (v AdlabGroupDetailV30DataDataAdInfoAdTarget) Ptr() *AdlabGroupDetailV30DataDataAdInfoAdTarget {
	return &v
}
