/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageUpdateV2Age
type AudiencePackageUpdateV2Age string

// List of audience_package_update_v2_age
const (
	AGE_BELOW_18_AudiencePackageUpdateV2Age      AudiencePackageUpdateV2Age = "AGE_BELOW_18"
	AGE_BETWEEN_18_19_AudiencePackageUpdateV2Age AudiencePackageUpdateV2Age = "AGE_BETWEEN_18_19"
	AGE_ABOVE_50_AudiencePackageUpdateV2Age      AudiencePackageUpdateV2Age = "AGE_ABOVE_50"
	AGE_BETWEEN_41_45_AudiencePackageUpdateV2Age AudiencePackageUpdateV2Age = "AGE_BETWEEN_41_45"
	AGE_BETWEEN_46_50_AudiencePackageUpdateV2Age AudiencePackageUpdateV2Age = "AGE_BETWEEN_46_50"
	AGE_BETWEEN_36_40_AudiencePackageUpdateV2Age AudiencePackageUpdateV2Age = "AGE_BETWEEN_36_40"
	AGE_BETWEEN_41_49_AudiencePackageUpdateV2Age AudiencePackageUpdateV2Age = "AGE_BETWEEN_41_49"
	AGE_BETWEEN_24_30_AudiencePackageUpdateV2Age AudiencePackageUpdateV2Age = "AGE_BETWEEN_24_30"
	AGE_BETWEEN_51_55_AudiencePackageUpdateV2Age AudiencePackageUpdateV2Age = "AGE_BETWEEN_51_55"
	AGE_BETWEEN_20_23_AudiencePackageUpdateV2Age AudiencePackageUpdateV2Age = "AGE_BETWEEN_20_23"
	AGE_BETWEEN_31_35_AudiencePackageUpdateV2Age AudiencePackageUpdateV2Age = "AGE_BETWEEN_31_35"
	AGE_BETWEEN_18_23_AudiencePackageUpdateV2Age AudiencePackageUpdateV2Age = "AGE_BETWEEN_18_23"
	AGE_BETWEEN_31_40_AudiencePackageUpdateV2Age AudiencePackageUpdateV2Age = "AGE_BETWEEN_31_40"
	AGE_BETWEEN_56_59_AudiencePackageUpdateV2Age AudiencePackageUpdateV2Age = "AGE_BETWEEN_56_59"
	AGE_ABOVE_60_AudiencePackageUpdateV2Age      AudiencePackageUpdateV2Age = "AGE_ABOVE_60"
)

// All allowed values of AudiencePackageUpdateV2Age enum
var AllowedAudiencePackageUpdateV2AgeEnumValues = []AudiencePackageUpdateV2Age{
	"AGE_BELOW_18",
	"AGE_BETWEEN_18_19",
	"AGE_ABOVE_50",
	"AGE_BETWEEN_41_45",
	"AGE_BETWEEN_46_50",
	"AGE_BETWEEN_36_40",
	"AGE_BETWEEN_41_49",
	"AGE_BETWEEN_24_30",
	"AGE_BETWEEN_51_55",
	"AGE_BETWEEN_20_23",
	"AGE_BETWEEN_31_35",
	"AGE_BETWEEN_18_23",
	"AGE_BETWEEN_31_40",
	"AGE_BETWEEN_56_59",
	"AGE_ABOVE_60",
}

// NewAudiencePackageUpdateV2AgeFromValue returns a pointer to a valid AudiencePackageUpdateV2Age
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageUpdateV2AgeFromValue(v string) (*AudiencePackageUpdateV2Age, error) {
	ev := AudiencePackageUpdateV2Age(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageUpdateV2Age: valid values are %v", v, AllowedAudiencePackageUpdateV2AgeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageUpdateV2Age) IsValid() bool {
	for _, existing := range AllowedAudiencePackageUpdateV2AgeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_update_v2_age value
func (v AudiencePackageUpdateV2Age) Ptr() *AudiencePackageUpdateV2Age {
	return &v
}
