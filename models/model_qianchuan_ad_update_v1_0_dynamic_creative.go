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

// QianchuanAdUpdateV10DynamicCreative
type QianchuanAdUpdateV10DynamicCreative int64

// List of qianchuan_ad_update_v1.0_dynamic_creative
const (
	Enum_0_QianchuanAdUpdateV10DynamicCreative QianchuanAdUpdateV10DynamicCreative = 0
	Enum_1_QianchuanAdUpdateV10DynamicCreative QianchuanAdUpdateV10DynamicCreative = 1
)

// All allowed values of QianchuanAdUpdateV10DynamicCreative enum
var AllowedQianchuanAdUpdateV10DynamicCreativeEnumValues = []QianchuanAdUpdateV10DynamicCreative{
	0,
	1,
}

// NewQianchuanAdUpdateV10DynamicCreativeFromValue returns a pointer to a valid QianchuanAdUpdateV10DynamicCreative
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10DynamicCreativeFromValue(v int64) (*QianchuanAdUpdateV10DynamicCreative, error) {
	ev := QianchuanAdUpdateV10DynamicCreative(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10DynamicCreative: valid values are %v", v, AllowedQianchuanAdUpdateV10DynamicCreativeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10DynamicCreative) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10DynamicCreativeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_dynamic_creative value
func (v QianchuanAdUpdateV10DynamicCreative) Ptr() *QianchuanAdUpdateV10DynamicCreative {
	return &v
}
