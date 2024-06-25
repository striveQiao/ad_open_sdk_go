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

// CampaignGetV2DataListMarketingPurpose
type CampaignGetV2DataListMarketingPurpose string

// List of campaign_get_v2_data_list_marketing_purpose
const (
	CONVERSION_CampaignGetV2DataListMarketingPurpose  CampaignGetV2DataListMarketingPurpose = "CONVERSION"
	UNLIMITED_CampaignGetV2DataListMarketingPurpose   CampaignGetV2DataListMarketingPurpose = "UNLIMITED"
	ACKNOWLEDGE_CampaignGetV2DataListMarketingPurpose CampaignGetV2DataListMarketingPurpose = "ACKNOWLEDGE"
	INTENTION_CampaignGetV2DataListMarketingPurpose   CampaignGetV2DataListMarketingPurpose = "INTENTION"
)

// All allowed values of CampaignGetV2DataListMarketingPurpose enum
var AllowedCampaignGetV2DataListMarketingPurposeEnumValues = []CampaignGetV2DataListMarketingPurpose{
	"CONVERSION",
	"UNLIMITED",
	"ACKNOWLEDGE",
	"INTENTION",
}

// NewCampaignGetV2DataListMarketingPurposeFromValue returns a pointer to a valid CampaignGetV2DataListMarketingPurpose
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignGetV2DataListMarketingPurposeFromValue(v string) (*CampaignGetV2DataListMarketingPurpose, error) {
	ev := CampaignGetV2DataListMarketingPurpose(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignGetV2DataListMarketingPurpose: valid values are %v", v, AllowedCampaignGetV2DataListMarketingPurposeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignGetV2DataListMarketingPurpose) IsValid() bool {
	for _, existing := range AllowedCampaignGetV2DataListMarketingPurposeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign_get_v2_data_list_marketing_purpose value
func (v CampaignGetV2DataListMarketingPurpose) Ptr() *CampaignGetV2DataListMarketingPurpose {
	return &v
}
