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

// ProjectCreateV30AudienceFilterAwemeFansCount
type ProjectCreateV30AudienceFilterAwemeFansCount int64

// List of project_create_v3.0_audience_filter_aweme_fans_count
const (
	Enum_1000_ProjectCreateV30AudienceFilterAwemeFansCount ProjectCreateV30AudienceFilterAwemeFansCount = 1000
	Enum_200_ProjectCreateV30AudienceFilterAwemeFansCount  ProjectCreateV30AudienceFilterAwemeFansCount = 200
	Enum_500_ProjectCreateV30AudienceFilterAwemeFansCount  ProjectCreateV30AudienceFilterAwemeFansCount = 500
)

// All allowed values of ProjectCreateV30AudienceFilterAwemeFansCount enum
var AllowedProjectCreateV30AudienceFilterAwemeFansCountEnumValues = []ProjectCreateV30AudienceFilterAwemeFansCount{
	1000,
	200,
	500,
}

// NewProjectCreateV30AudienceFilterAwemeFansCountFromValue returns a pointer to a valid ProjectCreateV30AudienceFilterAwemeFansCount
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AudienceFilterAwemeFansCountFromValue(v int64) (*ProjectCreateV30AudienceFilterAwemeFansCount, error) {
	ev := ProjectCreateV30AudienceFilterAwemeFansCount(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AudienceFilterAwemeFansCount: valid values are %v", v, AllowedProjectCreateV30AudienceFilterAwemeFansCountEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AudienceFilterAwemeFansCount) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AudienceFilterAwemeFansCountEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_audience_filter_aweme_fans_count value
func (v ProjectCreateV30AudienceFilterAwemeFansCount) Ptr() *ProjectCreateV30AudienceFilterAwemeFansCount {
	return &v
}
