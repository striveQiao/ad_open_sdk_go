/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30LaunchType
type ProjectCreateV30LaunchType string

// List of project_create_v3.0_launch_type
const (
	DIRECT_OPEN_ProjectCreateV30LaunchType   ProjectCreateV30LaunchType = "DIRECT_OPEN"
	EXTERNAL_OPEN_ProjectCreateV30LaunchType ProjectCreateV30LaunchType = "EXTERNAL_OPEN"
)

// All allowed values of ProjectCreateV30LaunchType enum
var AllowedProjectCreateV30LaunchTypeEnumValues = []ProjectCreateV30LaunchType{
	"DIRECT_OPEN",
	"EXTERNAL_OPEN",
}

// NewProjectCreateV30LaunchTypeFromValue returns a pointer to a valid ProjectCreateV30LaunchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30LaunchTypeFromValue(v string) (*ProjectCreateV30LaunchType, error) {
	ev := ProjectCreateV30LaunchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30LaunchType: valid values are %v", v, AllowedProjectCreateV30LaunchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30LaunchType) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30LaunchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_launch_type value
func (v ProjectCreateV30LaunchType) Ptr() *ProjectCreateV30LaunchType {
	return &v
}
