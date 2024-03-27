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

// AdlabGroupCreateV30CampaignInfoLandingType
type AdlabGroupCreateV30CampaignInfoLandingType string

// List of adlab_group_create_v3.0_campaign_info_landing_type
const (
	APP_AdlabGroupCreateV30CampaignInfoLandingType  AdlabGroupCreateV30CampaignInfoLandingType = "APP"
	LINK_AdlabGroupCreateV30CampaignInfoLandingType AdlabGroupCreateV30CampaignInfoLandingType = "LINK"
)

// All allowed values of AdlabGroupCreateV30CampaignInfoLandingType enum
var AllowedAdlabGroupCreateV30CampaignInfoLandingTypeEnumValues = []AdlabGroupCreateV30CampaignInfoLandingType{
	"APP",
	"LINK",
}

// NewAdlabGroupCreateV30CampaignInfoLandingTypeFromValue returns a pointer to a valid AdlabGroupCreateV30CampaignInfoLandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupCreateV30CampaignInfoLandingTypeFromValue(v string) (*AdlabGroupCreateV30CampaignInfoLandingType, error) {
	ev := AdlabGroupCreateV30CampaignInfoLandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupCreateV30CampaignInfoLandingType: valid values are %v", v, AllowedAdlabGroupCreateV30CampaignInfoLandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupCreateV30CampaignInfoLandingType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupCreateV30CampaignInfoLandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_create_v3.0_campaign_info_landing_type value
func (v AdlabGroupCreateV30CampaignInfoLandingType) Ptr() *AdlabGroupCreateV30CampaignInfoLandingType {
	return &v
}
