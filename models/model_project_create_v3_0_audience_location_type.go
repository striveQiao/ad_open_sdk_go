/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30AudienceLocationType
type ProjectCreateV30AudienceLocationType string

// List of project_create_v3.0_audience_location_type
const (
	ALL_ProjectCreateV30AudienceLocationType     ProjectCreateV30AudienceLocationType = "ALL"
	CURRENT_ProjectCreateV30AudienceLocationType ProjectCreateV30AudienceLocationType = "CURRENT"
	HOME_ProjectCreateV30AudienceLocationType    ProjectCreateV30AudienceLocationType = "HOME"
	TRAVEL_ProjectCreateV30AudienceLocationType  ProjectCreateV30AudienceLocationType = "TRAVEL"
)

// All allowed values of ProjectCreateV30AudienceLocationType enum
var AllowedProjectCreateV30AudienceLocationTypeEnumValues = []ProjectCreateV30AudienceLocationType{
	"ALL",
	"CURRENT",
	"HOME",
	"TRAVEL",
}

// NewProjectCreateV30AudienceLocationTypeFromValue returns a pointer to a valid ProjectCreateV30AudienceLocationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AudienceLocationTypeFromValue(v string) (*ProjectCreateV30AudienceLocationType, error) {
	ev := ProjectCreateV30AudienceLocationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AudienceLocationType: valid values are %v", v, AllowedProjectCreateV30AudienceLocationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AudienceLocationType) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AudienceLocationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_audience_location_type value
func (v ProjectCreateV30AudienceLocationType) Ptr() *ProjectCreateV30AudienceLocationType {
	return &v
}
