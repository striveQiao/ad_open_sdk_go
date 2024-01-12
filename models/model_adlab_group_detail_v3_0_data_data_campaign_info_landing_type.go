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

// AdlabGroupDetailV30DataDataCampaignInfoLandingType
type AdlabGroupDetailV30DataDataCampaignInfoLandingType string

// List of adlab_group_detail_v3.0_data_data_campaign_info_landing_type
const (
	APP_AdlabGroupDetailV30DataDataCampaignInfoLandingType  AdlabGroupDetailV30DataDataCampaignInfoLandingType = "APP"
	DPA_AdlabGroupDetailV30DataDataCampaignInfoLandingType  AdlabGroupDetailV30DataDataCampaignInfoLandingType = "DPA"
	LINK_AdlabGroupDetailV30DataDataCampaignInfoLandingType AdlabGroupDetailV30DataDataCampaignInfoLandingType = "LINK"
	LIVE_AdlabGroupDetailV30DataDataCampaignInfoLandingType AdlabGroupDetailV30DataDataCampaignInfoLandingType = "LIVE"
	SHOP_AdlabGroupDetailV30DataDataCampaignInfoLandingType AdlabGroupDetailV30DataDataCampaignInfoLandingType = "SHOP"
)

// All allowed values of AdlabGroupDetailV30DataDataCampaignInfoLandingType enum
var AllowedAdlabGroupDetailV30DataDataCampaignInfoLandingTypeEnumValues = []AdlabGroupDetailV30DataDataCampaignInfoLandingType{
	"APP",
	"DPA",
	"LINK",
	"LIVE",
	"SHOP",
}

// NewAdlabGroupDetailV30DataDataCampaignInfoLandingTypeFromValue returns a pointer to a valid AdlabGroupDetailV30DataDataCampaignInfoLandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupDetailV30DataDataCampaignInfoLandingTypeFromValue(v string) (*AdlabGroupDetailV30DataDataCampaignInfoLandingType, error) {
	ev := AdlabGroupDetailV30DataDataCampaignInfoLandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupDetailV30DataDataCampaignInfoLandingType: valid values are %v", v, AllowedAdlabGroupDetailV30DataDataCampaignInfoLandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupDetailV30DataDataCampaignInfoLandingType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupDetailV30DataDataCampaignInfoLandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_detail_v3.0_data_data_campaign_info_landing_type value
func (v AdlabGroupDetailV30DataDataCampaignInfoLandingType) Ptr() *AdlabGroupDetailV30DataDataCampaignInfoLandingType {
	return &v
}
