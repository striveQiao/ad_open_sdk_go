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

// BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecond
type BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecond string

// List of budget_group_list_v3.0_data_budget_groups_budget_group_status_second
const (
	ACCOUNT_EXCEEDED_BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecond BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecond = "ACCOUNT_EXCEEDED"
	GROUP_EXCEEDED_BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecond   BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecond = "GROUP_EXCEEDED"
)

// All allowed values of BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecond enum
var AllowedBudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecondEnumValues = []BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecond{
	"ACCOUNT_EXCEEDED",
	"GROUP_EXCEEDED",
}

// NewBudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecondFromValue returns a pointer to a valid BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecond
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecondFromValue(v string) (*BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecond, error) {
	ev := BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecond(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecond: valid values are %v", v, AllowedBudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecondEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecond) IsValid() bool {
	for _, existing := range AllowedBudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecondEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to budget_group_list_v3.0_data_budget_groups_budget_group_status_second value
func (v BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecond) Ptr() *BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecond {
	return &v
}
