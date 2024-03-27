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

// ProjectListV30DataListAudienceLocationType
type ProjectListV30DataListAudienceLocationType string

// List of project_list_v3.0_data_list_audience_location_type
const (
	ALL_ProjectListV30DataListAudienceLocationType     ProjectListV30DataListAudienceLocationType = "ALL"
	CURRENT_ProjectListV30DataListAudienceLocationType ProjectListV30DataListAudienceLocationType = "CURRENT"
	HOME_ProjectListV30DataListAudienceLocationType    ProjectListV30DataListAudienceLocationType = "HOME"
	TRAVEL_ProjectListV30DataListAudienceLocationType  ProjectListV30DataListAudienceLocationType = "TRAVEL"
)

// All allowed values of ProjectListV30DataListAudienceLocationType enum
var AllowedProjectListV30DataListAudienceLocationTypeEnumValues = []ProjectListV30DataListAudienceLocationType{
	"ALL",
	"CURRENT",
	"HOME",
	"TRAVEL",
}

// NewProjectListV30DataListAudienceLocationTypeFromValue returns a pointer to a valid ProjectListV30DataListAudienceLocationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAudienceLocationTypeFromValue(v string) (*ProjectListV30DataListAudienceLocationType, error) {
	ev := ProjectListV30DataListAudienceLocationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAudienceLocationType: valid values are %v", v, AllowedProjectListV30DataListAudienceLocationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAudienceLocationType) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAudienceLocationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_audience_location_type value
func (v ProjectListV30DataListAudienceLocationType) Ptr() *ProjectListV30DataListAudienceLocationType {
	return &v
}
