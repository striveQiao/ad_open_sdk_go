/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListAudienceSuperiorPopularityType
type ProjectListV30DataListAudienceSuperiorPopularityType string

// List of project_list_v3.0_data_list_audience_superior_popularity_type
const (
	GAME_ProjectListV30DataListAudienceSuperiorPopularityType ProjectListV30DataListAudienceSuperiorPopularityType = "GAME"
	NONE_ProjectListV30DataListAudienceSuperiorPopularityType ProjectListV30DataListAudienceSuperiorPopularityType = "NONE"
)

// All allowed values of ProjectListV30DataListAudienceSuperiorPopularityType enum
var AllowedProjectListV30DataListAudienceSuperiorPopularityTypeEnumValues = []ProjectListV30DataListAudienceSuperiorPopularityType{
	"GAME",
	"NONE",
}

// NewProjectListV30DataListAudienceSuperiorPopularityTypeFromValue returns a pointer to a valid ProjectListV30DataListAudienceSuperiorPopularityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAudienceSuperiorPopularityTypeFromValue(v string) (*ProjectListV30DataListAudienceSuperiorPopularityType, error) {
	ev := ProjectListV30DataListAudienceSuperiorPopularityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAudienceSuperiorPopularityType: valid values are %v", v, AllowedProjectListV30DataListAudienceSuperiorPopularityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAudienceSuperiorPopularityType) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAudienceSuperiorPopularityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_audience_superior_popularity_type value
func (v ProjectListV30DataListAudienceSuperiorPopularityType) Ptr() *ProjectListV30DataListAudienceSuperiorPopularityType {
	return &v
}
