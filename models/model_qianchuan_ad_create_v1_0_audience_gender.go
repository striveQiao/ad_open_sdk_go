/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdCreateV10AudienceGender
type QianchuanAdCreateV10AudienceGender string

// List of qianchuan_ad_create_v1.0_audience_gender
const (
	GENDER_FEMALE_QianchuanAdCreateV10AudienceGender QianchuanAdCreateV10AudienceGender = "GENDER_FEMALE"
	GENDER_MALE_QianchuanAdCreateV10AudienceGender   QianchuanAdCreateV10AudienceGender = "GENDER_MALE"
	NONE_QianchuanAdCreateV10AudienceGender          QianchuanAdCreateV10AudienceGender = "NONE"
)

// All allowed values of QianchuanAdCreateV10AudienceGender enum
var AllowedQianchuanAdCreateV10AudienceGenderEnumValues = []QianchuanAdCreateV10AudienceGender{
	"GENDER_FEMALE",
	"GENDER_MALE",
	"NONE",
}

// NewQianchuanAdCreateV10AudienceGenderFromValue returns a pointer to a valid QianchuanAdCreateV10AudienceGender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10AudienceGenderFromValue(v string) (*QianchuanAdCreateV10AudienceGender, error) {
	ev := QianchuanAdCreateV10AudienceGender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10AudienceGender: valid values are %v", v, AllowedQianchuanAdCreateV10AudienceGenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10AudienceGender) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10AudienceGenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_audience_gender value
func (v QianchuanAdCreateV10AudienceGender) Ptr() *QianchuanAdCreateV10AudienceGender {
	return &v
}
