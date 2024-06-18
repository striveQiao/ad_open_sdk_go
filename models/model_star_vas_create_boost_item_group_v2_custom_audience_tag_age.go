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

// StarVasCreateBoostItemGroupV2CustomAudienceTagAge
type StarVasCreateBoostItemGroupV2CustomAudienceTagAge string

// List of star_vas_create_boost_item_group_v2_custom_audience_tag_age
const (
	Enum_18_TO_23_StarVasCreateBoostItemGroupV2CustomAudienceTagAge StarVasCreateBoostItemGroupV2CustomAudienceTagAge = "18_TO_23"
	Enum_24_TO_30_StarVasCreateBoostItemGroupV2CustomAudienceTagAge StarVasCreateBoostItemGroupV2CustomAudienceTagAge = "24_TO_30"
	Enum_31_TO_40_StarVasCreateBoostItemGroupV2CustomAudienceTagAge StarVasCreateBoostItemGroupV2CustomAudienceTagAge = "31_TO_40"
	Enum_41_TO_50_StarVasCreateBoostItemGroupV2CustomAudienceTagAge StarVasCreateBoostItemGroupV2CustomAudienceTagAge = "41_TO_50"
	INFINITE_StarVasCreateBoostItemGroupV2CustomAudienceTagAge      StarVasCreateBoostItemGroupV2CustomAudienceTagAge = "INFINITE"
)

// All allowed values of StarVasCreateBoostItemGroupV2CustomAudienceTagAge enum
var AllowedStarVasCreateBoostItemGroupV2CustomAudienceTagAgeEnumValues = []StarVasCreateBoostItemGroupV2CustomAudienceTagAge{
	"18_TO_23",
	"24_TO_30",
	"31_TO_40",
	"41_TO_50",
	"INFINITE",
}

// NewStarVasCreateBoostItemGroupV2CustomAudienceTagAgeFromValue returns a pointer to a valid StarVasCreateBoostItemGroupV2CustomAudienceTagAge
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarVasCreateBoostItemGroupV2CustomAudienceTagAgeFromValue(v string) (*StarVasCreateBoostItemGroupV2CustomAudienceTagAge, error) {
	ev := StarVasCreateBoostItemGroupV2CustomAudienceTagAge(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarVasCreateBoostItemGroupV2CustomAudienceTagAge: valid values are %v", v, AllowedStarVasCreateBoostItemGroupV2CustomAudienceTagAgeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarVasCreateBoostItemGroupV2CustomAudienceTagAge) IsValid() bool {
	for _, existing := range AllowedStarVasCreateBoostItemGroupV2CustomAudienceTagAgeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_vas_create_boost_item_group_v2_custom_audience_tag_age value
func (v StarVasCreateBoostItemGroupV2CustomAudienceTagAge) Ptr() *StarVasCreateBoostItemGroupV2CustomAudienceTagAge {
	return &v
}
