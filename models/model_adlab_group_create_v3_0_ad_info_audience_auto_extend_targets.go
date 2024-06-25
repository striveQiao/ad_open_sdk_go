/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupCreateV30AdInfoAudienceAutoExtendTargets
type AdlabGroupCreateV30AdInfoAudienceAutoExtendTargets string

// List of adlab_group_create_v3.0_ad_info_audience_auto_extend_targets
const (
	AD_TAG_AdlabGroupCreateV30AdInfoAudienceAutoExtendTargets          AdlabGroupCreateV30AdInfoAudienceAutoExtendTargets = "AD_TAG"
	AGE_AdlabGroupCreateV30AdInfoAudienceAutoExtendTargets             AdlabGroupCreateV30AdInfoAudienceAutoExtendTargets = "AGE"
	CUSTOM_AUDIENCE_AdlabGroupCreateV30AdInfoAudienceAutoExtendTargets AdlabGroupCreateV30AdInfoAudienceAutoExtendTargets = "CUSTOM_AUDIENCE"
	GENDER_AdlabGroupCreateV30AdInfoAudienceAutoExtendTargets          AdlabGroupCreateV30AdInfoAudienceAutoExtendTargets = "GENDER"
	INTEREST_TAG_AdlabGroupCreateV30AdInfoAudienceAutoExtendTargets    AdlabGroupCreateV30AdInfoAudienceAutoExtendTargets = "INTEREST_TAG"
	REGION_AdlabGroupCreateV30AdInfoAudienceAutoExtendTargets          AdlabGroupCreateV30AdInfoAudienceAutoExtendTargets = "REGION"
)

// All allowed values of AdlabGroupCreateV30AdInfoAudienceAutoExtendTargets enum
var AllowedAdlabGroupCreateV30AdInfoAudienceAutoExtendTargetsEnumValues = []AdlabGroupCreateV30AdInfoAudienceAutoExtendTargets{
	"AD_TAG",
	"AGE",
	"CUSTOM_AUDIENCE",
	"GENDER",
	"INTEREST_TAG",
	"REGION",
}

// NewAdlabGroupCreateV30AdInfoAudienceAutoExtendTargetsFromValue returns a pointer to a valid AdlabGroupCreateV30AdInfoAudienceAutoExtendTargets
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupCreateV30AdInfoAudienceAutoExtendTargetsFromValue(v string) (*AdlabGroupCreateV30AdInfoAudienceAutoExtendTargets, error) {
	ev := AdlabGroupCreateV30AdInfoAudienceAutoExtendTargets(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupCreateV30AdInfoAudienceAutoExtendTargets: valid values are %v", v, AllowedAdlabGroupCreateV30AdInfoAudienceAutoExtendTargetsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupCreateV30AdInfoAudienceAutoExtendTargets) IsValid() bool {
	for _, existing := range AllowedAdlabGroupCreateV30AdInfoAudienceAutoExtendTargetsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_create_v3.0_ad_info_audience_auto_extend_targets value
func (v AdlabGroupCreateV30AdInfoAudienceAutoExtendTargets) Ptr() *AdlabGroupCreateV30AdInfoAudienceAutoExtendTargets {
	return &v
}
