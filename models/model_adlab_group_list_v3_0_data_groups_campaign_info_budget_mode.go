/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupListV30DataGroupsCampaignInfoBudgetMode
type AdlabGroupListV30DataGroupsCampaignInfoBudgetMode string

// List of adlab_group_list_v3.0_data_groups_campaign_info_budget_mode
const (
	INFINITE_AdlabGroupListV30DataGroupsCampaignInfoBudgetMode AdlabGroupListV30DataGroupsCampaignInfoBudgetMode = "INFINITE"
	PERDAY_AdlabGroupListV30DataGroupsCampaignInfoBudgetMode   AdlabGroupListV30DataGroupsCampaignInfoBudgetMode = "PERDAY"
	TOTAL_AdlabGroupListV30DataGroupsCampaignInfoBudgetMode    AdlabGroupListV30DataGroupsCampaignInfoBudgetMode = "TOTAL"
)

// All allowed values of AdlabGroupListV30DataGroupsCampaignInfoBudgetMode enum
var AllowedAdlabGroupListV30DataGroupsCampaignInfoBudgetModeEnumValues = []AdlabGroupListV30DataGroupsCampaignInfoBudgetMode{
	"INFINITE",
	"PERDAY",
	"TOTAL",
}

// NewAdlabGroupListV30DataGroupsCampaignInfoBudgetModeFromValue returns a pointer to a valid AdlabGroupListV30DataGroupsCampaignInfoBudgetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupListV30DataGroupsCampaignInfoBudgetModeFromValue(v string) (*AdlabGroupListV30DataGroupsCampaignInfoBudgetMode, error) {
	ev := AdlabGroupListV30DataGroupsCampaignInfoBudgetMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupListV30DataGroupsCampaignInfoBudgetMode: valid values are %v", v, AllowedAdlabGroupListV30DataGroupsCampaignInfoBudgetModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupListV30DataGroupsCampaignInfoBudgetMode) IsValid() bool {
	for _, existing := range AllowedAdlabGroupListV30DataGroupsCampaignInfoBudgetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_list_v3.0_data_groups_campaign_info_budget_mode value
func (v AdlabGroupListV30DataGroupsCampaignInfoBudgetMode) Ptr() *AdlabGroupListV30DataGroupsCampaignInfoBudgetMode {
	return &v
}
