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

// BudgetGroupCreateV30DeliveryMode
type BudgetGroupCreateV30DeliveryMode string

// List of budget_group_create_v3.0_delivery_mode
const (
	MANUAL_BudgetGroupCreateV30DeliveryMode     BudgetGroupCreateV30DeliveryMode = "MANUAL"
	PROCEDURAL_BudgetGroupCreateV30DeliveryMode BudgetGroupCreateV30DeliveryMode = "PROCEDURAL"
)

// All allowed values of BudgetGroupCreateV30DeliveryMode enum
var AllowedBudgetGroupCreateV30DeliveryModeEnumValues = []BudgetGroupCreateV30DeliveryMode{
	"MANUAL",
	"PROCEDURAL",
}

// NewBudgetGroupCreateV30DeliveryModeFromValue returns a pointer to a valid BudgetGroupCreateV30DeliveryMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBudgetGroupCreateV30DeliveryModeFromValue(v string) (*BudgetGroupCreateV30DeliveryMode, error) {
	ev := BudgetGroupCreateV30DeliveryMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BudgetGroupCreateV30DeliveryMode: valid values are %v", v, AllowedBudgetGroupCreateV30DeliveryModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BudgetGroupCreateV30DeliveryMode) IsValid() bool {
	for _, existing := range AllowedBudgetGroupCreateV30DeliveryModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to budget_group_create_v3.0_delivery_mode value
func (v BudgetGroupCreateV30DeliveryMode) Ptr() *BudgetGroupCreateV30DeliveryMode {
	return &v
}
