/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupListV30DataGroupsCampaignInfoMarketingPurpose
type AdlabGroupListV30DataGroupsCampaignInfoMarketingPurpose string

// List of adlab_group_list_v3.0_data_groups_campaign_info_marketing_purpose
const (
	ACKNOWLEDGE_AdlabGroupListV30DataGroupsCampaignInfoMarketingPurpose AdlabGroupListV30DataGroupsCampaignInfoMarketingPurpose = "ACKNOWLEDGE"
	CONVERSION_AdlabGroupListV30DataGroupsCampaignInfoMarketingPurpose  AdlabGroupListV30DataGroupsCampaignInfoMarketingPurpose = "CONVERSION"
	INTENTION_AdlabGroupListV30DataGroupsCampaignInfoMarketingPurpose   AdlabGroupListV30DataGroupsCampaignInfoMarketingPurpose = "INTENTION"
	UNLIMITED_AdlabGroupListV30DataGroupsCampaignInfoMarketingPurpose   AdlabGroupListV30DataGroupsCampaignInfoMarketingPurpose = "UNLIMITED"
)

// All allowed values of AdlabGroupListV30DataGroupsCampaignInfoMarketingPurpose enum
var AllowedAdlabGroupListV30DataGroupsCampaignInfoMarketingPurposeEnumValues = []AdlabGroupListV30DataGroupsCampaignInfoMarketingPurpose{
	"ACKNOWLEDGE",
	"CONVERSION",
	"INTENTION",
	"UNLIMITED",
}

// NewAdlabGroupListV30DataGroupsCampaignInfoMarketingPurposeFromValue returns a pointer to a valid AdlabGroupListV30DataGroupsCampaignInfoMarketingPurpose
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupListV30DataGroupsCampaignInfoMarketingPurposeFromValue(v string) (*AdlabGroupListV30DataGroupsCampaignInfoMarketingPurpose, error) {
	ev := AdlabGroupListV30DataGroupsCampaignInfoMarketingPurpose(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupListV30DataGroupsCampaignInfoMarketingPurpose: valid values are %v", v, AllowedAdlabGroupListV30DataGroupsCampaignInfoMarketingPurposeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupListV30DataGroupsCampaignInfoMarketingPurpose) IsValid() bool {
	for _, existing := range AllowedAdlabGroupListV30DataGroupsCampaignInfoMarketingPurposeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_list_v3.0_data_groups_campaign_info_marketing_purpose value
func (v AdlabGroupListV30DataGroupsCampaignInfoMarketingPurpose) Ptr() *AdlabGroupListV30DataGroupsCampaignInfoMarketingPurpose {
	return &v
}
