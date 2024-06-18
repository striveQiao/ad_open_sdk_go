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

// AdlabGroupDetailV30DataDataAdInfoAudienceGender
type AdlabGroupDetailV30DataDataAdInfoAudienceGender string

// List of adlab_group_detail_v3.0_data_data_ad_info_audience_gender
const (
	GENDER_FEMALE_AdlabGroupDetailV30DataDataAdInfoAudienceGender AdlabGroupDetailV30DataDataAdInfoAudienceGender = "GENDER_FEMALE"
	GENDER_MALE_AdlabGroupDetailV30DataDataAdInfoAudienceGender   AdlabGroupDetailV30DataDataAdInfoAudienceGender = "GENDER_MALE"
	NONE_AdlabGroupDetailV30DataDataAdInfoAudienceGender          AdlabGroupDetailV30DataDataAdInfoAudienceGender = "NONE"
)

// All allowed values of AdlabGroupDetailV30DataDataAdInfoAudienceGender enum
var AllowedAdlabGroupDetailV30DataDataAdInfoAudienceGenderEnumValues = []AdlabGroupDetailV30DataDataAdInfoAudienceGender{
	"GENDER_FEMALE",
	"GENDER_MALE",
	"NONE",
}

// NewAdlabGroupDetailV30DataDataAdInfoAudienceGenderFromValue returns a pointer to a valid AdlabGroupDetailV30DataDataAdInfoAudienceGender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupDetailV30DataDataAdInfoAudienceGenderFromValue(v string) (*AdlabGroupDetailV30DataDataAdInfoAudienceGender, error) {
	ev := AdlabGroupDetailV30DataDataAdInfoAudienceGender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupDetailV30DataDataAdInfoAudienceGender: valid values are %v", v, AllowedAdlabGroupDetailV30DataDataAdInfoAudienceGenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupDetailV30DataDataAdInfoAudienceGender) IsValid() bool {
	for _, existing := range AllowedAdlabGroupDetailV30DataDataAdInfoAudienceGenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_detail_v3.0_data_data_ad_info_audience_gender value
func (v AdlabGroupDetailV30DataDataAdInfoAudienceGender) Ptr() *AdlabGroupDetailV30DataDataAdInfoAudienceGender {
	return &v
}
