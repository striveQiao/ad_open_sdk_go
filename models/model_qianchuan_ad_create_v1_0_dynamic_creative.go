/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdCreateV10DynamicCreative
type QianchuanAdCreateV10DynamicCreative int64

// List of qianchuan_ad_create_v1.0_dynamic_creative
const (
	Enum_0_QianchuanAdCreateV10DynamicCreative QianchuanAdCreateV10DynamicCreative = 0
	Enum_1_QianchuanAdCreateV10DynamicCreative QianchuanAdCreateV10DynamicCreative = 1
)

// All allowed values of QianchuanAdCreateV10DynamicCreative enum
var AllowedQianchuanAdCreateV10DynamicCreativeEnumValues = []QianchuanAdCreateV10DynamicCreative{
	0,
	1,
}

// NewQianchuanAdCreateV10DynamicCreativeFromValue returns a pointer to a valid QianchuanAdCreateV10DynamicCreative
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10DynamicCreativeFromValue(v int64) (*QianchuanAdCreateV10DynamicCreative, error) {
	ev := QianchuanAdCreateV10DynamicCreative(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10DynamicCreative: valid values are %v", v, AllowedQianchuanAdCreateV10DynamicCreativeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10DynamicCreative) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10DynamicCreativeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_dynamic_creative value
func (v QianchuanAdCreateV10DynamicCreative) Ptr() *QianchuanAdCreateV10DynamicCreative {
	return &v
}
