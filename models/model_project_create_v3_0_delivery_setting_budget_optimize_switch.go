/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30DeliverySettingBudgetOptimizeSwitch
type ProjectCreateV30DeliverySettingBudgetOptimizeSwitch string

// List of project_create_v3.0_delivery_setting_budget_optimize_switch
const (
	OFF_ProjectCreateV30DeliverySettingBudgetOptimizeSwitch ProjectCreateV30DeliverySettingBudgetOptimizeSwitch = "OFF"
	ON_ProjectCreateV30DeliverySettingBudgetOptimizeSwitch  ProjectCreateV30DeliverySettingBudgetOptimizeSwitch = "ON"
)

// All allowed values of ProjectCreateV30DeliverySettingBudgetOptimizeSwitch enum
var AllowedProjectCreateV30DeliverySettingBudgetOptimizeSwitchEnumValues = []ProjectCreateV30DeliverySettingBudgetOptimizeSwitch{
	"OFF",
	"ON",
}

// NewProjectCreateV30DeliverySettingBudgetOptimizeSwitchFromValue returns a pointer to a valid ProjectCreateV30DeliverySettingBudgetOptimizeSwitch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30DeliverySettingBudgetOptimizeSwitchFromValue(v string) (*ProjectCreateV30DeliverySettingBudgetOptimizeSwitch, error) {
	ev := ProjectCreateV30DeliverySettingBudgetOptimizeSwitch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30DeliverySettingBudgetOptimizeSwitch: valid values are %v", v, AllowedProjectCreateV30DeliverySettingBudgetOptimizeSwitchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30DeliverySettingBudgetOptimizeSwitch) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30DeliverySettingBudgetOptimizeSwitchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_delivery_setting_budget_optimize_switch value
func (v ProjectCreateV30DeliverySettingBudgetOptimizeSwitch) Ptr() *ProjectCreateV30DeliverySettingBudgetOptimizeSwitch {
	return &v
}
