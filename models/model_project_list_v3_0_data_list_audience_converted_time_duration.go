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

// ProjectListV30DataListAudienceConvertedTimeDuration
type ProjectListV30DataListAudienceConvertedTimeDuration string

// List of project_list_v3.0_data_list_audience_converted_time_duration
const (
	NONE_ProjectListV30DataListAudienceConvertedTimeDuration         ProjectListV30DataListAudienceConvertedTimeDuration = "NONE"
	ONE_MONTH_ProjectListV30DataListAudienceConvertedTimeDuration    ProjectListV30DataListAudienceConvertedTimeDuration = "ONE_MONTH"
	SEVEN_DAY_ProjectListV30DataListAudienceConvertedTimeDuration    ProjectListV30DataListAudienceConvertedTimeDuration = "SEVEN_DAY"
	SIX_MONTH_ProjectListV30DataListAudienceConvertedTimeDuration    ProjectListV30DataListAudienceConvertedTimeDuration = "SIX_MONTH"
	THREE_MONTH_ProjectListV30DataListAudienceConvertedTimeDuration  ProjectListV30DataListAudienceConvertedTimeDuration = "THREE_MONTH"
	TODAY_ProjectListV30DataListAudienceConvertedTimeDuration        ProjectListV30DataListAudienceConvertedTimeDuration = "TODAY"
	TWELVE_MONTH_ProjectListV30DataListAudienceConvertedTimeDuration ProjectListV30DataListAudienceConvertedTimeDuration = "TWELVE_MONTH"
)

// All allowed values of ProjectListV30DataListAudienceConvertedTimeDuration enum
var AllowedProjectListV30DataListAudienceConvertedTimeDurationEnumValues = []ProjectListV30DataListAudienceConvertedTimeDuration{
	"NONE",
	"ONE_MONTH",
	"SEVEN_DAY",
	"SIX_MONTH",
	"THREE_MONTH",
	"TODAY",
	"TWELVE_MONTH",
}

// NewProjectListV30DataListAudienceConvertedTimeDurationFromValue returns a pointer to a valid ProjectListV30DataListAudienceConvertedTimeDuration
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAudienceConvertedTimeDurationFromValue(v string) (*ProjectListV30DataListAudienceConvertedTimeDuration, error) {
	ev := ProjectListV30DataListAudienceConvertedTimeDuration(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAudienceConvertedTimeDuration: valid values are %v", v, AllowedProjectListV30DataListAudienceConvertedTimeDurationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAudienceConvertedTimeDuration) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAudienceConvertedTimeDurationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_audience_converted_time_duration value
func (v ProjectListV30DataListAudienceConvertedTimeDuration) Ptr() *ProjectListV30DataListAudienceConvertedTimeDuration {
	return &v
}
