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

// QianchuanAccountBudgetUpdateV10BudgetMode
type QianchuanAccountBudgetUpdateV10BudgetMode string

// List of qianchuan_account_budget_update_v1.0_budget_mode
const (
	INFINITE_QianchuanAccountBudgetUpdateV10BudgetMode  QianchuanAccountBudgetUpdateV10BudgetMode = "INFINITE"
	SPECIFIED_QianchuanAccountBudgetUpdateV10BudgetMode QianchuanAccountBudgetUpdateV10BudgetMode = "SPECIFIED"
)

// All allowed values of QianchuanAccountBudgetUpdateV10BudgetMode enum
var AllowedQianchuanAccountBudgetUpdateV10BudgetModeEnumValues = []QianchuanAccountBudgetUpdateV10BudgetMode{
	"INFINITE",
	"SPECIFIED",
}

// NewQianchuanAccountBudgetUpdateV10BudgetModeFromValue returns a pointer to a valid QianchuanAccountBudgetUpdateV10BudgetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAccountBudgetUpdateV10BudgetModeFromValue(v string) (*QianchuanAccountBudgetUpdateV10BudgetMode, error) {
	ev := QianchuanAccountBudgetUpdateV10BudgetMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAccountBudgetUpdateV10BudgetMode: valid values are %v", v, AllowedQianchuanAccountBudgetUpdateV10BudgetModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAccountBudgetUpdateV10BudgetMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAccountBudgetUpdateV10BudgetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_account_budget_update_v1.0_budget_mode value
func (v QianchuanAccountBudgetUpdateV10BudgetMode) Ptr() *QianchuanAccountBudgetUpdateV10BudgetMode {
	return &v
}
