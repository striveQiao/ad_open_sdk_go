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

// QianchuanAwemeOrderCreateV10AudienceAge
type QianchuanAwemeOrderCreateV10AudienceAge string

// List of qianchuan_aweme_order_create_v1.0_audience_age
const (
	AGE_BETWEEN_18_23_QianchuanAwemeOrderCreateV10AudienceAge QianchuanAwemeOrderCreateV10AudienceAge = "AGE_BETWEEN_18_23"
	AGE_BETWEEN_24_30_QianchuanAwemeOrderCreateV10AudienceAge QianchuanAwemeOrderCreateV10AudienceAge = "AGE_BETWEEN_24_30"
	AGE_BETWEEN_31_40_QianchuanAwemeOrderCreateV10AudienceAge QianchuanAwemeOrderCreateV10AudienceAge = "AGE_BETWEEN_31_40"
	AGE_BETWEEN_41_50_QianchuanAwemeOrderCreateV10AudienceAge QianchuanAwemeOrderCreateV10AudienceAge = "AGE_BETWEEN_41_50"
)

// All allowed values of QianchuanAwemeOrderCreateV10AudienceAge enum
var AllowedQianchuanAwemeOrderCreateV10AudienceAgeEnumValues = []QianchuanAwemeOrderCreateV10AudienceAge{
	"AGE_BETWEEN_18_23",
	"AGE_BETWEEN_24_30",
	"AGE_BETWEEN_31_40",
	"AGE_BETWEEN_41_50",
}

// NewQianchuanAwemeOrderCreateV10AudienceAgeFromValue returns a pointer to a valid QianchuanAwemeOrderCreateV10AudienceAge
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderCreateV10AudienceAgeFromValue(v string) (*QianchuanAwemeOrderCreateV10AudienceAge, error) {
	ev := QianchuanAwemeOrderCreateV10AudienceAge(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderCreateV10AudienceAge: valid values are %v", v, AllowedQianchuanAwemeOrderCreateV10AudienceAgeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderCreateV10AudienceAge) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderCreateV10AudienceAgeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_create_v1.0_audience_age value
func (v QianchuanAwemeOrderCreateV10AudienceAge) Ptr() *QianchuanAwemeOrderCreateV10AudienceAge {
	return &v
}
