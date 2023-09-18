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

// AdvertiserUpdateBudgetV2BudgetMode
type AdvertiserUpdateBudgetV2BudgetMode string

// List of advertiser_update_budget_v2_budget_mode
const (
	BUDGET_MODE_DAY_AdvertiserUpdateBudgetV2BudgetMode      AdvertiserUpdateBudgetV2BudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_INFINITE_AdvertiserUpdateBudgetV2BudgetMode AdvertiserUpdateBudgetV2BudgetMode = "BUDGET_MODE_INFINITE"
)

// All allowed values of AdvertiserUpdateBudgetV2BudgetMode enum
var AllowedAdvertiserUpdateBudgetV2BudgetModeEnumValues = []AdvertiserUpdateBudgetV2BudgetMode{
	"BUDGET_MODE_DAY",
	"BUDGET_MODE_INFINITE",
}

// NewAdvertiserUpdateBudgetV2BudgetModeFromValue returns a pointer to a valid AdvertiserUpdateBudgetV2BudgetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserUpdateBudgetV2BudgetModeFromValue(v string) (*AdvertiserUpdateBudgetV2BudgetMode, error) {
	ev := AdvertiserUpdateBudgetV2BudgetMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserUpdateBudgetV2BudgetMode: valid values are %v", v, AllowedAdvertiserUpdateBudgetV2BudgetModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserUpdateBudgetV2BudgetMode) IsValid() bool {
	for _, existing := range AllowedAdvertiserUpdateBudgetV2BudgetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_update_budget_v2_budget_mode value
func (v AdvertiserUpdateBudgetV2BudgetMode) Ptr() *AdvertiserUpdateBudgetV2BudgetMode {
	return &v
}
