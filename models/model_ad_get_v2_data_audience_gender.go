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

// AdGetV2DataAudienceGender
type AdGetV2DataAudienceGender string

// List of ad_get_v2_data_audience_gender
const (
	GENDER_FEMALE_AdGetV2DataAudienceGender    AdGetV2DataAudienceGender = "GENDER_FEMALE"
	GENDER_UNLIMITED_AdGetV2DataAudienceGender AdGetV2DataAudienceGender = "GENDER_UNLIMITED"
	NONE_AdGetV2DataAudienceGender             AdGetV2DataAudienceGender = "NONE"
	GENDER_MALE_AdGetV2DataAudienceGender      AdGetV2DataAudienceGender = "GENDER_MALE"
)

// All allowed values of AdGetV2DataAudienceGender enum
var AllowedAdGetV2DataAudienceGenderEnumValues = []AdGetV2DataAudienceGender{
	"GENDER_FEMALE",
	"GENDER_UNLIMITED",
	"NONE",
	"GENDER_MALE",
}

// NewAdGetV2DataAudienceGenderFromValue returns a pointer to a valid AdGetV2DataAudienceGender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAudienceGenderFromValue(v string) (*AdGetV2DataAudienceGender, error) {
	ev := AdGetV2DataAudienceGender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAudienceGender: valid values are %v", v, AllowedAdGetV2DataAudienceGenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAudienceGender) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAudienceGenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_audience_gender value
func (v AdGetV2DataAudienceGender) Ptr() *AdGetV2DataAudienceGender {
	return &v
}
