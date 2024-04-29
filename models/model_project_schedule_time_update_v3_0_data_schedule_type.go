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

// ProjectScheduleTimeUpdateV30DataScheduleType
type ProjectScheduleTimeUpdateV30DataScheduleType string

// List of project_schedule_time_update_v3.0_data_schedule_type
const (
	SCHEDULE_FROM_NOW_ProjectScheduleTimeUpdateV30DataScheduleType  ProjectScheduleTimeUpdateV30DataScheduleType = "SCHEDULE_FROM_NOW"
	SCHEDULE_START_END_ProjectScheduleTimeUpdateV30DataScheduleType ProjectScheduleTimeUpdateV30DataScheduleType = "SCHEDULE_START_END"
)

// All allowed values of ProjectScheduleTimeUpdateV30DataScheduleType enum
var AllowedProjectScheduleTimeUpdateV30DataScheduleTypeEnumValues = []ProjectScheduleTimeUpdateV30DataScheduleType{
	"SCHEDULE_FROM_NOW",
	"SCHEDULE_START_END",
}

// NewProjectScheduleTimeUpdateV30DataScheduleTypeFromValue returns a pointer to a valid ProjectScheduleTimeUpdateV30DataScheduleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectScheduleTimeUpdateV30DataScheduleTypeFromValue(v string) (*ProjectScheduleTimeUpdateV30DataScheduleType, error) {
	ev := ProjectScheduleTimeUpdateV30DataScheduleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectScheduleTimeUpdateV30DataScheduleType: valid values are %v", v, AllowedProjectScheduleTimeUpdateV30DataScheduleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectScheduleTimeUpdateV30DataScheduleType) IsValid() bool {
	for _, existing := range AllowedProjectScheduleTimeUpdateV30DataScheduleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_schedule_time_update_v3.0_data_schedule_type value
func (v ProjectScheduleTimeUpdateV30DataScheduleType) Ptr() *ProjectScheduleTimeUpdateV30DataScheduleType {
	return &v
}
