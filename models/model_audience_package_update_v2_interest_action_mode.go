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

// AudiencePackageUpdateV2InterestActionMode
type AudiencePackageUpdateV2InterestActionMode string

// List of audience_package_update_v2_interest_action_mode
const (
	CUSTOM_AudiencePackageUpdateV2InterestActionMode    AudiencePackageUpdateV2InterestActionMode = "CUSTOM"
	RECOMMEND_AudiencePackageUpdateV2InterestActionMode AudiencePackageUpdateV2InterestActionMode = "RECOMMEND"
	UNLIMITED_AudiencePackageUpdateV2InterestActionMode AudiencePackageUpdateV2InterestActionMode = "UNLIMITED"
)

// All allowed values of AudiencePackageUpdateV2InterestActionMode enum
var AllowedAudiencePackageUpdateV2InterestActionModeEnumValues = []AudiencePackageUpdateV2InterestActionMode{
	"CUSTOM",
	"RECOMMEND",
	"UNLIMITED",
}

// NewAudiencePackageUpdateV2InterestActionModeFromValue returns a pointer to a valid AudiencePackageUpdateV2InterestActionMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageUpdateV2InterestActionModeFromValue(v string) (*AudiencePackageUpdateV2InterestActionMode, error) {
	ev := AudiencePackageUpdateV2InterestActionMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageUpdateV2InterestActionMode: valid values are %v", v, AllowedAudiencePackageUpdateV2InterestActionModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageUpdateV2InterestActionMode) IsValid() bool {
	for _, existing := range AllowedAudiencePackageUpdateV2InterestActionModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_update_v2_interest_action_mode value
func (v AudiencePackageUpdateV2InterestActionMode) Ptr() *AudiencePackageUpdateV2InterestActionMode {
	return &v
}
