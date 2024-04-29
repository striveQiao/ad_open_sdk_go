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

// AdlabGroupDetailV30DataDataCampaignInfoBudgetMode
type AdlabGroupDetailV30DataDataCampaignInfoBudgetMode string

// List of adlab_group_detail_v3.0_data_data_campaign_info_budget_mode
const (
	INFINITE_AdlabGroupDetailV30DataDataCampaignInfoBudgetMode AdlabGroupDetailV30DataDataCampaignInfoBudgetMode = "INFINITE"
	PERDAY_AdlabGroupDetailV30DataDataCampaignInfoBudgetMode   AdlabGroupDetailV30DataDataCampaignInfoBudgetMode = "PERDAY"
	TOTAL_AdlabGroupDetailV30DataDataCampaignInfoBudgetMode    AdlabGroupDetailV30DataDataCampaignInfoBudgetMode = "TOTAL"
)

// All allowed values of AdlabGroupDetailV30DataDataCampaignInfoBudgetMode enum
var AllowedAdlabGroupDetailV30DataDataCampaignInfoBudgetModeEnumValues = []AdlabGroupDetailV30DataDataCampaignInfoBudgetMode{
	"INFINITE",
	"PERDAY",
	"TOTAL",
}

// NewAdlabGroupDetailV30DataDataCampaignInfoBudgetModeFromValue returns a pointer to a valid AdlabGroupDetailV30DataDataCampaignInfoBudgetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupDetailV30DataDataCampaignInfoBudgetModeFromValue(v string) (*AdlabGroupDetailV30DataDataCampaignInfoBudgetMode, error) {
	ev := AdlabGroupDetailV30DataDataCampaignInfoBudgetMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupDetailV30DataDataCampaignInfoBudgetMode: valid values are %v", v, AllowedAdlabGroupDetailV30DataDataCampaignInfoBudgetModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupDetailV30DataDataCampaignInfoBudgetMode) IsValid() bool {
	for _, existing := range AllowedAdlabGroupDetailV30DataDataCampaignInfoBudgetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_detail_v3.0_data_data_campaign_info_budget_mode value
func (v AdlabGroupDetailV30DataDataCampaignInfoBudgetMode) Ptr() *AdlabGroupDetailV30DataDataCampaignInfoBudgetMode {
	return &v
}
