/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectUpdateV30DeliverySettingFilterNightSwitch
type ProjectUpdateV30DeliverySettingFilterNightSwitch string

// List of project_update_v3.0_delivery_setting_filter_night_switch
const (
	OFF_ProjectUpdateV30DeliverySettingFilterNightSwitch ProjectUpdateV30DeliverySettingFilterNightSwitch = "OFF"
	ON_ProjectUpdateV30DeliverySettingFilterNightSwitch  ProjectUpdateV30DeliverySettingFilterNightSwitch = "ON"
)

// All allowed values of ProjectUpdateV30DeliverySettingFilterNightSwitch enum
var AllowedProjectUpdateV30DeliverySettingFilterNightSwitchEnumValues = []ProjectUpdateV30DeliverySettingFilterNightSwitch{
	"OFF",
	"ON",
}

// NewProjectUpdateV30DeliverySettingFilterNightSwitchFromValue returns a pointer to a valid ProjectUpdateV30DeliverySettingFilterNightSwitch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30DeliverySettingFilterNightSwitchFromValue(v string) (*ProjectUpdateV30DeliverySettingFilterNightSwitch, error) {
	ev := ProjectUpdateV30DeliverySettingFilterNightSwitch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30DeliverySettingFilterNightSwitch: valid values are %v", v, AllowedProjectUpdateV30DeliverySettingFilterNightSwitchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30DeliverySettingFilterNightSwitch) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30DeliverySettingFilterNightSwitchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_delivery_setting_filter_night_switch value
func (v ProjectUpdateV30DeliverySettingFilterNightSwitch) Ptr() *ProjectUpdateV30DeliverySettingFilterNightSwitch {
	return &v
}
