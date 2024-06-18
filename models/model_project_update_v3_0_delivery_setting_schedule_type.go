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

// ProjectUpdateV30DeliverySettingScheduleType
type ProjectUpdateV30DeliverySettingScheduleType string

// List of project_update_v3.0_delivery_setting_schedule_type
const (
	SCHEDULE_FROM_NOW_ProjectUpdateV30DeliverySettingScheduleType  ProjectUpdateV30DeliverySettingScheduleType = "SCHEDULE_FROM_NOW"
	SCHEDULE_START_END_ProjectUpdateV30DeliverySettingScheduleType ProjectUpdateV30DeliverySettingScheduleType = "SCHEDULE_START_END"
)

// All allowed values of ProjectUpdateV30DeliverySettingScheduleType enum
var AllowedProjectUpdateV30DeliverySettingScheduleTypeEnumValues = []ProjectUpdateV30DeliverySettingScheduleType{
	"SCHEDULE_FROM_NOW",
	"SCHEDULE_START_END",
}

// NewProjectUpdateV30DeliverySettingScheduleTypeFromValue returns a pointer to a valid ProjectUpdateV30DeliverySettingScheduleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30DeliverySettingScheduleTypeFromValue(v string) (*ProjectUpdateV30DeliverySettingScheduleType, error) {
	ev := ProjectUpdateV30DeliverySettingScheduleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30DeliverySettingScheduleType: valid values are %v", v, AllowedProjectUpdateV30DeliverySettingScheduleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30DeliverySettingScheduleType) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30DeliverySettingScheduleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_delivery_setting_schedule_type value
func (v ProjectUpdateV30DeliverySettingScheduleType) Ptr() *ProjectUpdateV30DeliverySettingScheduleType {
	return &v
}
