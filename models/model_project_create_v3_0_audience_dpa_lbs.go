/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30AudienceDpaLbs
type ProjectCreateV30AudienceDpaLbs string

// List of project_create_v3.0_audience_dpa_lbs
const (
	OFF_ProjectCreateV30AudienceDpaLbs ProjectCreateV30AudienceDpaLbs = "OFF"
	ON_ProjectCreateV30AudienceDpaLbs  ProjectCreateV30AudienceDpaLbs = "ON"
)

// All allowed values of ProjectCreateV30AudienceDpaLbs enum
var AllowedProjectCreateV30AudienceDpaLbsEnumValues = []ProjectCreateV30AudienceDpaLbs{
	"OFF",
	"ON",
}

// NewProjectCreateV30AudienceDpaLbsFromValue returns a pointer to a valid ProjectCreateV30AudienceDpaLbs
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AudienceDpaLbsFromValue(v string) (*ProjectCreateV30AudienceDpaLbs, error) {
	ev := ProjectCreateV30AudienceDpaLbs(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AudienceDpaLbs: valid values are %v", v, AllowedProjectCreateV30AudienceDpaLbsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AudienceDpaLbs) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AudienceDpaLbsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_audience_dpa_lbs value
func (v ProjectCreateV30AudienceDpaLbs) Ptr() *ProjectCreateV30AudienceDpaLbs {
	return &v
}
