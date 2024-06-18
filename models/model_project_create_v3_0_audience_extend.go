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

// ProjectCreateV30AudienceExtend
type ProjectCreateV30AudienceExtend string

// List of project_create_v3.0_audience_extend
const (
	OFF_ProjectCreateV30AudienceExtend ProjectCreateV30AudienceExtend = "OFF"
	ON_ProjectCreateV30AudienceExtend  ProjectCreateV30AudienceExtend = "ON"
)

// All allowed values of ProjectCreateV30AudienceExtend enum
var AllowedProjectCreateV30AudienceExtendEnumValues = []ProjectCreateV30AudienceExtend{
	"OFF",
	"ON",
}

// NewProjectCreateV30AudienceExtendFromValue returns a pointer to a valid ProjectCreateV30AudienceExtend
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AudienceExtendFromValue(v string) (*ProjectCreateV30AudienceExtend, error) {
	ev := ProjectCreateV30AudienceExtend(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AudienceExtend: valid values are %v", v, AllowedProjectCreateV30AudienceExtendEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AudienceExtend) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AudienceExtendEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_audience_extend value
func (v ProjectCreateV30AudienceExtend) Ptr() *ProjectCreateV30AudienceExtend {
	return &v
}
