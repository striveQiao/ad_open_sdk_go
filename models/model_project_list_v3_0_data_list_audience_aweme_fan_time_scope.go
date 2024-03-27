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

// ProjectListV30DataListAudienceAwemeFanTimeScope
type ProjectListV30DataListAudienceAwemeFanTimeScope string

// List of project_list_v3.0_data_list_audience_aweme_fan_time_scope
const (
	FIFTEEN_DAYS_ProjectListV30DataListAudienceAwemeFanTimeScope ProjectListV30DataListAudienceAwemeFanTimeScope = "FIFTEEN_DAYS"
	SIXTY_DAYS_ProjectListV30DataListAudienceAwemeFanTimeScope   ProjectListV30DataListAudienceAwemeFanTimeScope = "SIXTY_DAYS"
	THIRTY_DAYS_ProjectListV30DataListAudienceAwemeFanTimeScope  ProjectListV30DataListAudienceAwemeFanTimeScope = "THIRTY_DAYS"
)

// All allowed values of ProjectListV30DataListAudienceAwemeFanTimeScope enum
var AllowedProjectListV30DataListAudienceAwemeFanTimeScopeEnumValues = []ProjectListV30DataListAudienceAwemeFanTimeScope{
	"FIFTEEN_DAYS",
	"SIXTY_DAYS",
	"THIRTY_DAYS",
}

// NewProjectListV30DataListAudienceAwemeFanTimeScopeFromValue returns a pointer to a valid ProjectListV30DataListAudienceAwemeFanTimeScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAudienceAwemeFanTimeScopeFromValue(v string) (*ProjectListV30DataListAudienceAwemeFanTimeScope, error) {
	ev := ProjectListV30DataListAudienceAwemeFanTimeScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAudienceAwemeFanTimeScope: valid values are %v", v, AllowedProjectListV30DataListAudienceAwemeFanTimeScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAudienceAwemeFanTimeScope) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAudienceAwemeFanTimeScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_audience_aweme_fan_time_scope value
func (v ProjectListV30DataListAudienceAwemeFanTimeScope) Ptr() *ProjectListV30DataListAudienceAwemeFanTimeScope {
	return &v
}
