/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeAuthListGetV10FilteringAuthRange
type QianchuanAwemeAuthListGetV10FilteringAuthRange string

// List of qianchuan_aweme_auth_list_get_v1.0_filtering_auth_range
const (
	ALL_QianchuanAwemeAuthListGetV10FilteringAuthRange      QianchuanAwemeAuthListGetV10FilteringAuthRange = "ALL"
	BY_AWEME_QianchuanAwemeAuthListGetV10FilteringAuthRange QianchuanAwemeAuthListGetV10FilteringAuthRange = "BY_AWEME"
	BY_VIDEO_QianchuanAwemeAuthListGetV10FilteringAuthRange QianchuanAwemeAuthListGetV10FilteringAuthRange = "BY_VIDEO"
)

// All allowed values of QianchuanAwemeAuthListGetV10FilteringAuthRange enum
var AllowedQianchuanAwemeAuthListGetV10FilteringAuthRangeEnumValues = []QianchuanAwemeAuthListGetV10FilteringAuthRange{
	"ALL",
	"BY_AWEME",
	"BY_VIDEO",
}

// NewQianchuanAwemeAuthListGetV10FilteringAuthRangeFromValue returns a pointer to a valid QianchuanAwemeAuthListGetV10FilteringAuthRange
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeAuthListGetV10FilteringAuthRangeFromValue(v string) (*QianchuanAwemeAuthListGetV10FilteringAuthRange, error) {
	ev := QianchuanAwemeAuthListGetV10FilteringAuthRange(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeAuthListGetV10FilteringAuthRange: valid values are %v", v, AllowedQianchuanAwemeAuthListGetV10FilteringAuthRangeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeAuthListGetV10FilteringAuthRange) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeAuthListGetV10FilteringAuthRangeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_auth_list_get_v1.0_filtering_auth_range value
func (v QianchuanAwemeAuthListGetV10FilteringAuthRange) Ptr() *QianchuanAwemeAuthListGetV10FilteringAuthRange {
	return &v
}
