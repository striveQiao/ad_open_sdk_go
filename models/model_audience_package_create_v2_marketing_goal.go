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

// AudiencePackageCreateV2MarketingGoal
type AudiencePackageCreateV2MarketingGoal string

// List of audience_package_create_v2_marketing_goal
const (
	LIVE_AudiencePackageCreateV2MarketingGoal            AudiencePackageCreateV2MarketingGoal = "LIVE"
	VIDEO_AND_IMAGE_AudiencePackageCreateV2MarketingGoal AudiencePackageCreateV2MarketingGoal = "VIDEO_AND_IMAGE"
)

// All allowed values of AudiencePackageCreateV2MarketingGoal enum
var AllowedAudiencePackageCreateV2MarketingGoalEnumValues = []AudiencePackageCreateV2MarketingGoal{
	"LIVE",
	"VIDEO_AND_IMAGE",
}

// NewAudiencePackageCreateV2MarketingGoalFromValue returns a pointer to a valid AudiencePackageCreateV2MarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageCreateV2MarketingGoalFromValue(v string) (*AudiencePackageCreateV2MarketingGoal, error) {
	ev := AudiencePackageCreateV2MarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageCreateV2MarketingGoal: valid values are %v", v, AllowedAudiencePackageCreateV2MarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageCreateV2MarketingGoal) IsValid() bool {
	for _, existing := range AllowedAudiencePackageCreateV2MarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_create_v2_marketing_goal value
func (v AudiencePackageCreateV2MarketingGoal) Ptr() *AudiencePackageCreateV2MarketingGoal {
	return &v
}
