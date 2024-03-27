/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListAudienceFilterAwemeFansCount
type ProjectListV30DataListAudienceFilterAwemeFansCount int64

// List of project_list_v3.0_data_list_audience_filter_aweme_fans_count
const (
	Enum_1000_ProjectListV30DataListAudienceFilterAwemeFansCount ProjectListV30DataListAudienceFilterAwemeFansCount = 1000
	Enum_200_ProjectListV30DataListAudienceFilterAwemeFansCount  ProjectListV30DataListAudienceFilterAwemeFansCount = 200
	Enum_500_ProjectListV30DataListAudienceFilterAwemeFansCount  ProjectListV30DataListAudienceFilterAwemeFansCount = 500
)

// All allowed values of ProjectListV30DataListAudienceFilterAwemeFansCount enum
var AllowedProjectListV30DataListAudienceFilterAwemeFansCountEnumValues = []ProjectListV30DataListAudienceFilterAwemeFansCount{
	1000,
	200,
	500,
}

// NewProjectListV30DataListAudienceFilterAwemeFansCountFromValue returns a pointer to a valid ProjectListV30DataListAudienceFilterAwemeFansCount
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAudienceFilterAwemeFansCountFromValue(v int64) (*ProjectListV30DataListAudienceFilterAwemeFansCount, error) {
	ev := ProjectListV30DataListAudienceFilterAwemeFansCount(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAudienceFilterAwemeFansCount: valid values are %v", v, AllowedProjectListV30DataListAudienceFilterAwemeFansCountEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAudienceFilterAwemeFansCount) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAudienceFilterAwemeFansCountEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_audience_filter_aweme_fans_count value
func (v ProjectListV30DataListAudienceFilterAwemeFansCount) Ptr() *ProjectListV30DataListAudienceFilterAwemeFansCount {
	return &v
}
