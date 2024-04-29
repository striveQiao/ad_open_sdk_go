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

// AdlabGroupListV30DataGroupsAdInfoAudienceGender
type AdlabGroupListV30DataGroupsAdInfoAudienceGender string

// List of adlab_group_list_v3.0_data_groups_ad_info_audience_gender
const (
	GENDER_FEMALE_AdlabGroupListV30DataGroupsAdInfoAudienceGender AdlabGroupListV30DataGroupsAdInfoAudienceGender = "GENDER_FEMALE"
	GENDER_MALE_AdlabGroupListV30DataGroupsAdInfoAudienceGender   AdlabGroupListV30DataGroupsAdInfoAudienceGender = "GENDER_MALE"
	NONE_AdlabGroupListV30DataGroupsAdInfoAudienceGender          AdlabGroupListV30DataGroupsAdInfoAudienceGender = "NONE"
)

// All allowed values of AdlabGroupListV30DataGroupsAdInfoAudienceGender enum
var AllowedAdlabGroupListV30DataGroupsAdInfoAudienceGenderEnumValues = []AdlabGroupListV30DataGroupsAdInfoAudienceGender{
	"GENDER_FEMALE",
	"GENDER_MALE",
	"NONE",
}

// NewAdlabGroupListV30DataGroupsAdInfoAudienceGenderFromValue returns a pointer to a valid AdlabGroupListV30DataGroupsAdInfoAudienceGender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupListV30DataGroupsAdInfoAudienceGenderFromValue(v string) (*AdlabGroupListV30DataGroupsAdInfoAudienceGender, error) {
	ev := AdlabGroupListV30DataGroupsAdInfoAudienceGender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupListV30DataGroupsAdInfoAudienceGender: valid values are %v", v, AllowedAdlabGroupListV30DataGroupsAdInfoAudienceGenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupListV30DataGroupsAdInfoAudienceGender) IsValid() bool {
	for _, existing := range AllowedAdlabGroupListV30DataGroupsAdInfoAudienceGenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_list_v3.0_data_groups_ad_info_audience_gender value
func (v AdlabGroupListV30DataGroupsAdInfoAudienceGender) Ptr() *AdlabGroupListV30DataGroupsAdInfoAudienceGender {
	return &v
}
