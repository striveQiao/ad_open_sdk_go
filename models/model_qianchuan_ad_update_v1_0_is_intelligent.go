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

// QianchuanAdUpdateV10IsIntelligent
type QianchuanAdUpdateV10IsIntelligent int64

// List of qianchuan_ad_update_v1.0_is_intelligent
const (
	Enum_0_QianchuanAdUpdateV10IsIntelligent QianchuanAdUpdateV10IsIntelligent = 0
	Enum_1_QianchuanAdUpdateV10IsIntelligent QianchuanAdUpdateV10IsIntelligent = 1
)

// All allowed values of QianchuanAdUpdateV10IsIntelligent enum
var AllowedQianchuanAdUpdateV10IsIntelligentEnumValues = []QianchuanAdUpdateV10IsIntelligent{
	0,
	1,
}

// NewQianchuanAdUpdateV10IsIntelligentFromValue returns a pointer to a valid QianchuanAdUpdateV10IsIntelligent
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10IsIntelligentFromValue(v int64) (*QianchuanAdUpdateV10IsIntelligent, error) {
	ev := QianchuanAdUpdateV10IsIntelligent(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10IsIntelligent: valid values are %v", v, AllowedQianchuanAdUpdateV10IsIntelligentEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10IsIntelligent) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10IsIntelligentEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_is_intelligent value
func (v QianchuanAdUpdateV10IsIntelligent) Ptr() *QianchuanAdUpdateV10IsIntelligent {
	return &v
}
