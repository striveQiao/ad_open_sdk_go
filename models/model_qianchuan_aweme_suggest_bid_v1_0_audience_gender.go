/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeSuggestBidV10AudienceGender
type QianchuanAwemeSuggestBidV10AudienceGender string

// List of qianchuan_aweme_suggest_bid_v1.0_audience_gender
const (
	GENDER_FEMALE_QianchuanAwemeSuggestBidV10AudienceGender QianchuanAwemeSuggestBidV10AudienceGender = "GENDER_FEMALE"
	GENDER_MALE_QianchuanAwemeSuggestBidV10AudienceGender   QianchuanAwemeSuggestBidV10AudienceGender = "GENDER_MALE"
)

// All allowed values of QianchuanAwemeSuggestBidV10AudienceGender enum
var AllowedQianchuanAwemeSuggestBidV10AudienceGenderEnumValues = []QianchuanAwemeSuggestBidV10AudienceGender{
	"GENDER_FEMALE",
	"GENDER_MALE",
}

// NewQianchuanAwemeSuggestBidV10AudienceGenderFromValue returns a pointer to a valid QianchuanAwemeSuggestBidV10AudienceGender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeSuggestBidV10AudienceGenderFromValue(v string) (*QianchuanAwemeSuggestBidV10AudienceGender, error) {
	ev := QianchuanAwemeSuggestBidV10AudienceGender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeSuggestBidV10AudienceGender: valid values are %v", v, AllowedQianchuanAwemeSuggestBidV10AudienceGenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeSuggestBidV10AudienceGender) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeSuggestBidV10AudienceGenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_suggest_bid_v1.0_audience_gender value
func (v QianchuanAwemeSuggestBidV10AudienceGender) Ptr() *QianchuanAwemeSuggestBidV10AudienceGender {
	return &v
}
