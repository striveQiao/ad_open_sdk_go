/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeOrderCreateV10AudienceGender
type QianchuanAwemeOrderCreateV10AudienceGender string

// List of qianchuan_aweme_order_create_v1.0_audience_gender
const (
	GENDER_FEMALE_QianchuanAwemeOrderCreateV10AudienceGender QianchuanAwemeOrderCreateV10AudienceGender = "GENDER_FEMALE"
	GENDER_MALE_QianchuanAwemeOrderCreateV10AudienceGender   QianchuanAwemeOrderCreateV10AudienceGender = "GENDER_MALE"
)

// All allowed values of QianchuanAwemeOrderCreateV10AudienceGender enum
var AllowedQianchuanAwemeOrderCreateV10AudienceGenderEnumValues = []QianchuanAwemeOrderCreateV10AudienceGender{
	"GENDER_FEMALE",
	"GENDER_MALE",
}

// NewQianchuanAwemeOrderCreateV10AudienceGenderFromValue returns a pointer to a valid QianchuanAwemeOrderCreateV10AudienceGender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderCreateV10AudienceGenderFromValue(v string) (*QianchuanAwemeOrderCreateV10AudienceGender, error) {
	ev := QianchuanAwemeOrderCreateV10AudienceGender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderCreateV10AudienceGender: valid values are %v", v, AllowedQianchuanAwemeOrderCreateV10AudienceGenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderCreateV10AudienceGender) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderCreateV10AudienceGenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_create_v1.0_audience_gender value
func (v QianchuanAwemeOrderCreateV10AudienceGender) Ptr() *QianchuanAwemeOrderCreateV10AudienceGender {
	return &v
}
