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

// QianchuanAdDetailGetV10DataDynamicCreative
type QianchuanAdDetailGetV10DataDynamicCreative int64

// List of qianchuan_ad_detail_get_v1.0_data_dynamic_creative
const (
	Enum_0_QianchuanAdDetailGetV10DataDynamicCreative QianchuanAdDetailGetV10DataDynamicCreative = 0
	Enum_1_QianchuanAdDetailGetV10DataDynamicCreative QianchuanAdDetailGetV10DataDynamicCreative = 1
)

// All allowed values of QianchuanAdDetailGetV10DataDynamicCreative enum
var AllowedQianchuanAdDetailGetV10DataDynamicCreativeEnumValues = []QianchuanAdDetailGetV10DataDynamicCreative{
	0,
	1,
}

// NewQianchuanAdDetailGetV10DataDynamicCreativeFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataDynamicCreative
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataDynamicCreativeFromValue(v int64) (*QianchuanAdDetailGetV10DataDynamicCreative, error) {
	ev := QianchuanAdDetailGetV10DataDynamicCreative(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataDynamicCreative: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataDynamicCreativeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataDynamicCreative) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataDynamicCreativeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_dynamic_creative value
func (v QianchuanAdDetailGetV10DataDynamicCreative) Ptr() *QianchuanAdDetailGetV10DataDynamicCreative {
	return &v
}
