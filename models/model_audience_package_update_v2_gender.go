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

// AudiencePackageUpdateV2Gender
type AudiencePackageUpdateV2Gender string

// List of audience_package_update_v2_gender
const (
	GENDER_FEMALE_AudiencePackageUpdateV2Gender    AudiencePackageUpdateV2Gender = "GENDER_FEMALE"
	GENDER_UNLIMITED_AudiencePackageUpdateV2Gender AudiencePackageUpdateV2Gender = "GENDER_UNLIMITED"
	NONE_AudiencePackageUpdateV2Gender             AudiencePackageUpdateV2Gender = "NONE"
	GENDER_MALE_AudiencePackageUpdateV2Gender      AudiencePackageUpdateV2Gender = "GENDER_MALE"
)

// All allowed values of AudiencePackageUpdateV2Gender enum
var AllowedAudiencePackageUpdateV2GenderEnumValues = []AudiencePackageUpdateV2Gender{
	"GENDER_FEMALE",
	"GENDER_UNLIMITED",
	"NONE",
	"GENDER_MALE",
}

// NewAudiencePackageUpdateV2GenderFromValue returns a pointer to a valid AudiencePackageUpdateV2Gender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageUpdateV2GenderFromValue(v string) (*AudiencePackageUpdateV2Gender, error) {
	ev := AudiencePackageUpdateV2Gender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageUpdateV2Gender: valid values are %v", v, AllowedAudiencePackageUpdateV2GenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageUpdateV2Gender) IsValid() bool {
	for _, existing := range AllowedAudiencePackageUpdateV2GenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_update_v2_gender value
func (v AudiencePackageUpdateV2Gender) Ptr() *AudiencePackageUpdateV2Gender {
	return &v
}
