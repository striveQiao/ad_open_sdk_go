/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageUpdateV2ConvertedTimeDuration
type AudiencePackageUpdateV2ConvertedTimeDuration string

// List of audience_package_update_v2_converted_time_duration
const (
	THREE_MONTH_AudiencePackageUpdateV2ConvertedTimeDuration  AudiencePackageUpdateV2ConvertedTimeDuration = "THREE_MONTH"
	NONE_AudiencePackageUpdateV2ConvertedTimeDuration         AudiencePackageUpdateV2ConvertedTimeDuration = "NONE"
	SIX_MONTH_AudiencePackageUpdateV2ConvertedTimeDuration    AudiencePackageUpdateV2ConvertedTimeDuration = "SIX_MONTH"
	SEVEN_DAY_AudiencePackageUpdateV2ConvertedTimeDuration    AudiencePackageUpdateV2ConvertedTimeDuration = "SEVEN_DAY"
	ONE_MONTH_AudiencePackageUpdateV2ConvertedTimeDuration    AudiencePackageUpdateV2ConvertedTimeDuration = "ONE_MONTH"
	TWELVE_MONTH_AudiencePackageUpdateV2ConvertedTimeDuration AudiencePackageUpdateV2ConvertedTimeDuration = "TWELVE_MONTH"
	TODAY_AudiencePackageUpdateV2ConvertedTimeDuration        AudiencePackageUpdateV2ConvertedTimeDuration = "TODAY"
)

// All allowed values of AudiencePackageUpdateV2ConvertedTimeDuration enum
var AllowedAudiencePackageUpdateV2ConvertedTimeDurationEnumValues = []AudiencePackageUpdateV2ConvertedTimeDuration{
	"THREE_MONTH",
	"NONE",
	"SIX_MONTH",
	"SEVEN_DAY",
	"ONE_MONTH",
	"TWELVE_MONTH",
	"TODAY",
}

// NewAudiencePackageUpdateV2ConvertedTimeDurationFromValue returns a pointer to a valid AudiencePackageUpdateV2ConvertedTimeDuration
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageUpdateV2ConvertedTimeDurationFromValue(v string) (*AudiencePackageUpdateV2ConvertedTimeDuration, error) {
	ev := AudiencePackageUpdateV2ConvertedTimeDuration(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageUpdateV2ConvertedTimeDuration: valid values are %v", v, AllowedAudiencePackageUpdateV2ConvertedTimeDurationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageUpdateV2ConvertedTimeDuration) IsValid() bool {
	for _, existing := range AllowedAudiencePackageUpdateV2ConvertedTimeDurationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_update_v2_converted_time_duration value
func (v AudiencePackageUpdateV2ConvertedTimeDuration) Ptr() *AudiencePackageUpdateV2ConvertedTimeDuration {
	return &v
}
