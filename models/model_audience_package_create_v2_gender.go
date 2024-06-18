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

// AudiencePackageCreateV2Gender
type AudiencePackageCreateV2Gender string

// List of audience_package_create_v2_gender
const (
	GENDER_FEMALE_AudiencePackageCreateV2Gender    AudiencePackageCreateV2Gender = "GENDER_FEMALE"
	GENDER_UNLIMITED_AudiencePackageCreateV2Gender AudiencePackageCreateV2Gender = "GENDER_UNLIMITED"
	NONE_AudiencePackageCreateV2Gender             AudiencePackageCreateV2Gender = "NONE"
	GENDER_MALE_AudiencePackageCreateV2Gender      AudiencePackageCreateV2Gender = "GENDER_MALE"
)

// All allowed values of AudiencePackageCreateV2Gender enum
var AllowedAudiencePackageCreateV2GenderEnumValues = []AudiencePackageCreateV2Gender{
	"GENDER_FEMALE",
	"GENDER_UNLIMITED",
	"NONE",
	"GENDER_MALE",
}

// NewAudiencePackageCreateV2GenderFromValue returns a pointer to a valid AudiencePackageCreateV2Gender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageCreateV2GenderFromValue(v string) (*AudiencePackageCreateV2Gender, error) {
	ev := AudiencePackageCreateV2Gender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageCreateV2Gender: valid values are %v", v, AllowedAudiencePackageCreateV2GenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageCreateV2Gender) IsValid() bool {
	for _, existing := range AllowedAudiencePackageCreateV2GenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_create_v2_gender value
func (v AudiencePackageCreateV2Gender) Ptr() *AudiencePackageCreateV2Gender {
	return &v
}
