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

// CampaignGetV2DataListCampaignBudgetOptimization
type CampaignGetV2DataListCampaignBudgetOptimization string

// List of campaign_get_v2_data_list_campaign_budget_optimization
const (
	OFF_CampaignGetV2DataListCampaignBudgetOptimization CampaignGetV2DataListCampaignBudgetOptimization = "OFF"
	ON_CampaignGetV2DataListCampaignBudgetOptimization  CampaignGetV2DataListCampaignBudgetOptimization = "ON"
)

// All allowed values of CampaignGetV2DataListCampaignBudgetOptimization enum
var AllowedCampaignGetV2DataListCampaignBudgetOptimizationEnumValues = []CampaignGetV2DataListCampaignBudgetOptimization{
	"OFF",
	"ON",
}

// NewCampaignGetV2DataListCampaignBudgetOptimizationFromValue returns a pointer to a valid CampaignGetV2DataListCampaignBudgetOptimization
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignGetV2DataListCampaignBudgetOptimizationFromValue(v string) (*CampaignGetV2DataListCampaignBudgetOptimization, error) {
	ev := CampaignGetV2DataListCampaignBudgetOptimization(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignGetV2DataListCampaignBudgetOptimization: valid values are %v", v, AllowedCampaignGetV2DataListCampaignBudgetOptimizationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignGetV2DataListCampaignBudgetOptimization) IsValid() bool {
	for _, existing := range AllowedCampaignGetV2DataListCampaignBudgetOptimizationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign_get_v2_data_list_campaign_budget_optimization value
func (v CampaignGetV2DataListCampaignBudgetOptimization) Ptr() *CampaignGetV2DataListCampaignBudgetOptimization {
	return &v
}
