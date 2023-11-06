/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdUpdateV10CreativeAutoGenerate
type QianchuanAdUpdateV10CreativeAutoGenerate int64

// List of qianchuan_ad_update_v1.0_creative_auto_generate
const (
	Enum_0_QianchuanAdUpdateV10CreativeAutoGenerate QianchuanAdUpdateV10CreativeAutoGenerate = 0
	Enum_1_QianchuanAdUpdateV10CreativeAutoGenerate QianchuanAdUpdateV10CreativeAutoGenerate = 1
)

// All allowed values of QianchuanAdUpdateV10CreativeAutoGenerate enum
var AllowedQianchuanAdUpdateV10CreativeAutoGenerateEnumValues = []QianchuanAdUpdateV10CreativeAutoGenerate{
	0,
	1,
}

// NewQianchuanAdUpdateV10CreativeAutoGenerateFromValue returns a pointer to a valid QianchuanAdUpdateV10CreativeAutoGenerate
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10CreativeAutoGenerateFromValue(v int64) (*QianchuanAdUpdateV10CreativeAutoGenerate, error) {
	ev := QianchuanAdUpdateV10CreativeAutoGenerate(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10CreativeAutoGenerate: valid values are %v", v, AllowedQianchuanAdUpdateV10CreativeAutoGenerateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10CreativeAutoGenerate) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10CreativeAutoGenerateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_creative_auto_generate value
func (v QianchuanAdUpdateV10CreativeAutoGenerate) Ptr() *QianchuanAdUpdateV10CreativeAutoGenerate {
	return &v
}
