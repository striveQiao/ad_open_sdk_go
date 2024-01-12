/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectUpdateV30AudienceConvertedTimeDuration
type ProjectUpdateV30AudienceConvertedTimeDuration string

// List of project_update_v3.0_audience_converted_time_duration
const (
	NONE_ProjectUpdateV30AudienceConvertedTimeDuration         ProjectUpdateV30AudienceConvertedTimeDuration = "NONE"
	ONE_MONTH_ProjectUpdateV30AudienceConvertedTimeDuration    ProjectUpdateV30AudienceConvertedTimeDuration = "ONE_MONTH"
	SEVEN_DAY_ProjectUpdateV30AudienceConvertedTimeDuration    ProjectUpdateV30AudienceConvertedTimeDuration = "SEVEN_DAY"
	SIX_MONTH_ProjectUpdateV30AudienceConvertedTimeDuration    ProjectUpdateV30AudienceConvertedTimeDuration = "SIX_MONTH"
	THREE_MONTH_ProjectUpdateV30AudienceConvertedTimeDuration  ProjectUpdateV30AudienceConvertedTimeDuration = "THREE_MONTH"
	TODAY_ProjectUpdateV30AudienceConvertedTimeDuration        ProjectUpdateV30AudienceConvertedTimeDuration = "TODAY"
	TWELVE_MONTH_ProjectUpdateV30AudienceConvertedTimeDuration ProjectUpdateV30AudienceConvertedTimeDuration = "TWELVE_MONTH"
)

// All allowed values of ProjectUpdateV30AudienceConvertedTimeDuration enum
var AllowedProjectUpdateV30AudienceConvertedTimeDurationEnumValues = []ProjectUpdateV30AudienceConvertedTimeDuration{
	"NONE",
	"ONE_MONTH",
	"SEVEN_DAY",
	"SIX_MONTH",
	"THREE_MONTH",
	"TODAY",
	"TWELVE_MONTH",
}

// NewProjectUpdateV30AudienceConvertedTimeDurationFromValue returns a pointer to a valid ProjectUpdateV30AudienceConvertedTimeDuration
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30AudienceConvertedTimeDurationFromValue(v string) (*ProjectUpdateV30AudienceConvertedTimeDuration, error) {
	ev := ProjectUpdateV30AudienceConvertedTimeDuration(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30AudienceConvertedTimeDuration: valid values are %v", v, AllowedProjectUpdateV30AudienceConvertedTimeDurationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30AudienceConvertedTimeDuration) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30AudienceConvertedTimeDurationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_audience_converted_time_duration value
func (v ProjectUpdateV30AudienceConvertedTimeDuration) Ptr() *ProjectUpdateV30AudienceConvertedTimeDuration {
	return &v
}
