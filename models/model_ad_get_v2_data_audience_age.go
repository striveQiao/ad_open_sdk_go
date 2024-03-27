/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdGetV2DataAudienceAge
type AdGetV2DataAudienceAge string

// List of ad_get_v2_data_audience_age
const (
	AGE_BELOW_18_AdGetV2DataAudienceAge      AdGetV2DataAudienceAge = "AGE_BELOW_18"
	AGE_BETWEEN_41_49_AdGetV2DataAudienceAge AdGetV2DataAudienceAge = "AGE_BETWEEN_41_49"
	AGE_BETWEEN_18_23_AdGetV2DataAudienceAge AdGetV2DataAudienceAge = "AGE_BETWEEN_18_23"
	AGE_BETWEEN_36_40_AdGetV2DataAudienceAge AdGetV2DataAudienceAge = "AGE_BETWEEN_36_40"
	AGE_BETWEEN_31_40_AdGetV2DataAudienceAge AdGetV2DataAudienceAge = "AGE_BETWEEN_31_40"
	AGE_BETWEEN_24_30_AdGetV2DataAudienceAge AdGetV2DataAudienceAge = "AGE_BETWEEN_24_30"
	AGE_BETWEEN_41_45_AdGetV2DataAudienceAge AdGetV2DataAudienceAge = "AGE_BETWEEN_41_45"
	AGE_BETWEEN_56_59_AdGetV2DataAudienceAge AdGetV2DataAudienceAge = "AGE_BETWEEN_56_59"
	AGE_ABOVE_60_AdGetV2DataAudienceAge      AdGetV2DataAudienceAge = "AGE_ABOVE_60"
	AGE_BETWEEN_51_55_AdGetV2DataAudienceAge AdGetV2DataAudienceAge = "AGE_BETWEEN_51_55"
	AGE_BETWEEN_18_19_AdGetV2DataAudienceAge AdGetV2DataAudienceAge = "AGE_BETWEEN_18_19"
	AGE_ABOVE_50_AdGetV2DataAudienceAge      AdGetV2DataAudienceAge = "AGE_ABOVE_50"
	AGE_BETWEEN_31_35_AdGetV2DataAudienceAge AdGetV2DataAudienceAge = "AGE_BETWEEN_31_35"
	AGE_BETWEEN_20_23_AdGetV2DataAudienceAge AdGetV2DataAudienceAge = "AGE_BETWEEN_20_23"
	AGE_BETWEEN_46_50_AdGetV2DataAudienceAge AdGetV2DataAudienceAge = "AGE_BETWEEN_46_50"
)

// All allowed values of AdGetV2DataAudienceAge enum
var AllowedAdGetV2DataAudienceAgeEnumValues = []AdGetV2DataAudienceAge{
	"AGE_BELOW_18",
	"AGE_BETWEEN_41_49",
	"AGE_BETWEEN_18_23",
	"AGE_BETWEEN_36_40",
	"AGE_BETWEEN_31_40",
	"AGE_BETWEEN_24_30",
	"AGE_BETWEEN_41_45",
	"AGE_BETWEEN_56_59",
	"AGE_ABOVE_60",
	"AGE_BETWEEN_51_55",
	"AGE_BETWEEN_18_19",
	"AGE_ABOVE_50",
	"AGE_BETWEEN_31_35",
	"AGE_BETWEEN_20_23",
	"AGE_BETWEEN_46_50",
}

// NewAdGetV2DataAudienceAgeFromValue returns a pointer to a valid AdGetV2DataAudienceAge
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAudienceAgeFromValue(v string) (*AdGetV2DataAudienceAge, error) {
	ev := AdGetV2DataAudienceAge(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAudienceAge: valid values are %v", v, AllowedAdGetV2DataAudienceAgeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAudienceAge) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAudienceAgeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_audience_age value
func (v AdGetV2DataAudienceAge) Ptr() *AdGetV2DataAudienceAge {
	return &v
}
