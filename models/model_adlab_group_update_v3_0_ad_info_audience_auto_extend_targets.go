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

// AdlabGroupUpdateV30AdInfoAudienceAutoExtendTargets
type AdlabGroupUpdateV30AdInfoAudienceAutoExtendTargets string

// List of adlab_group_update_v3.0_ad_info_audience_auto_extend_targets
const (
	AD_TAG_AdlabGroupUpdateV30AdInfoAudienceAutoExtendTargets          AdlabGroupUpdateV30AdInfoAudienceAutoExtendTargets = "AD_TAG"
	AGE_AdlabGroupUpdateV30AdInfoAudienceAutoExtendTargets             AdlabGroupUpdateV30AdInfoAudienceAutoExtendTargets = "AGE"
	CUSTOM_AUDIENCE_AdlabGroupUpdateV30AdInfoAudienceAutoExtendTargets AdlabGroupUpdateV30AdInfoAudienceAutoExtendTargets = "CUSTOM_AUDIENCE"
	GENDER_AdlabGroupUpdateV30AdInfoAudienceAutoExtendTargets          AdlabGroupUpdateV30AdInfoAudienceAutoExtendTargets = "GENDER"
	INTEREST_TAG_AdlabGroupUpdateV30AdInfoAudienceAutoExtendTargets    AdlabGroupUpdateV30AdInfoAudienceAutoExtendTargets = "INTEREST_TAG"
	REGION_AdlabGroupUpdateV30AdInfoAudienceAutoExtendTargets          AdlabGroupUpdateV30AdInfoAudienceAutoExtendTargets = "REGION"
)

// All allowed values of AdlabGroupUpdateV30AdInfoAudienceAutoExtendTargets enum
var AllowedAdlabGroupUpdateV30AdInfoAudienceAutoExtendTargetsEnumValues = []AdlabGroupUpdateV30AdInfoAudienceAutoExtendTargets{
	"AD_TAG",
	"AGE",
	"CUSTOM_AUDIENCE",
	"GENDER",
	"INTEREST_TAG",
	"REGION",
}

// NewAdlabGroupUpdateV30AdInfoAudienceAutoExtendTargetsFromValue returns a pointer to a valid AdlabGroupUpdateV30AdInfoAudienceAutoExtendTargets
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupUpdateV30AdInfoAudienceAutoExtendTargetsFromValue(v string) (*AdlabGroupUpdateV30AdInfoAudienceAutoExtendTargets, error) {
	ev := AdlabGroupUpdateV30AdInfoAudienceAutoExtendTargets(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupUpdateV30AdInfoAudienceAutoExtendTargets: valid values are %v", v, AllowedAdlabGroupUpdateV30AdInfoAudienceAutoExtendTargetsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupUpdateV30AdInfoAudienceAutoExtendTargets) IsValid() bool {
	for _, existing := range AllowedAdlabGroupUpdateV30AdInfoAudienceAutoExtendTargetsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_update_v3.0_ad_info_audience_auto_extend_targets value
func (v AdlabGroupUpdateV30AdInfoAudienceAutoExtendTargets) Ptr() *AdlabGroupUpdateV30AdInfoAudienceAutoExtendTargets {
	return &v
}
