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

// BudgetGroupListV30DataBudgetGroupsAdType
type BudgetGroupListV30DataBudgetGroupsAdType string

// List of budget_group_list_v3.0_data_budget_groups_ad_type
const (
	ALL_BudgetGroupListV30DataBudgetGroupsAdType    BudgetGroupListV30DataBudgetGroupsAdType = "ALL"
	SEARCH_BudgetGroupListV30DataBudgetGroupsAdType BudgetGroupListV30DataBudgetGroupsAdType = "SEARCH"
)

// All allowed values of BudgetGroupListV30DataBudgetGroupsAdType enum
var AllowedBudgetGroupListV30DataBudgetGroupsAdTypeEnumValues = []BudgetGroupListV30DataBudgetGroupsAdType{
	"ALL",
	"SEARCH",
}

// NewBudgetGroupListV30DataBudgetGroupsAdTypeFromValue returns a pointer to a valid BudgetGroupListV30DataBudgetGroupsAdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBudgetGroupListV30DataBudgetGroupsAdTypeFromValue(v string) (*BudgetGroupListV30DataBudgetGroupsAdType, error) {
	ev := BudgetGroupListV30DataBudgetGroupsAdType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BudgetGroupListV30DataBudgetGroupsAdType: valid values are %v", v, AllowedBudgetGroupListV30DataBudgetGroupsAdTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BudgetGroupListV30DataBudgetGroupsAdType) IsValid() bool {
	for _, existing := range AllowedBudgetGroupListV30DataBudgetGroupsAdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to budget_group_list_v3.0_data_budget_groups_ad_type value
func (v BudgetGroupListV30DataBudgetGroupsAdType) Ptr() *BudgetGroupListV30DataBudgetGroupsAdType {
	return &v
}
