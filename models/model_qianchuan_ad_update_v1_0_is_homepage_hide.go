/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdUpdateV10IsHomepageHide
type QianchuanAdUpdateV10IsHomepageHide int64

// List of qianchuan_ad_update_v1.0_is_homepage_hide
const (
	Enum_0_QianchuanAdUpdateV10IsHomepageHide QianchuanAdUpdateV10IsHomepageHide = 0
	Enum_1_QianchuanAdUpdateV10IsHomepageHide QianchuanAdUpdateV10IsHomepageHide = 1
)

// All allowed values of QianchuanAdUpdateV10IsHomepageHide enum
var AllowedQianchuanAdUpdateV10IsHomepageHideEnumValues = []QianchuanAdUpdateV10IsHomepageHide{
	0,
	1,
}

// NewQianchuanAdUpdateV10IsHomepageHideFromValue returns a pointer to a valid QianchuanAdUpdateV10IsHomepageHide
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10IsHomepageHideFromValue(v int64) (*QianchuanAdUpdateV10IsHomepageHide, error) {
	ev := QianchuanAdUpdateV10IsHomepageHide(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10IsHomepageHide: valid values are %v", v, AllowedQianchuanAdUpdateV10IsHomepageHideEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10IsHomepageHide) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10IsHomepageHideEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_is_homepage_hide value
func (v QianchuanAdUpdateV10IsHomepageHide) Ptr() *QianchuanAdUpdateV10IsHomepageHide {
	return &v
}
