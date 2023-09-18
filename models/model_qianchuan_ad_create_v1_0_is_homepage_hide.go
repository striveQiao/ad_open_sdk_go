/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdCreateV10IsHomepageHide
type QianchuanAdCreateV10IsHomepageHide int64

// List of qianchuan_ad_create_v1.0_is_homepage_hide
const (
	Enum_0_QianchuanAdCreateV10IsHomepageHide QianchuanAdCreateV10IsHomepageHide = 0
	Enum_1_QianchuanAdCreateV10IsHomepageHide QianchuanAdCreateV10IsHomepageHide = 1
)

// All allowed values of QianchuanAdCreateV10IsHomepageHide enum
var AllowedQianchuanAdCreateV10IsHomepageHideEnumValues = []QianchuanAdCreateV10IsHomepageHide{
	0,
	1,
}

// NewQianchuanAdCreateV10IsHomepageHideFromValue returns a pointer to a valid QianchuanAdCreateV10IsHomepageHide
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10IsHomepageHideFromValue(v int64) (*QianchuanAdCreateV10IsHomepageHide, error) {
	ev := QianchuanAdCreateV10IsHomepageHide(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10IsHomepageHide: valid values are %v", v, AllowedQianchuanAdCreateV10IsHomepageHideEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10IsHomepageHide) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10IsHomepageHideEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_is_homepage_hide value
func (v QianchuanAdCreateV10IsHomepageHide) Ptr() *QianchuanAdCreateV10IsHomepageHide {
	return &v
}
