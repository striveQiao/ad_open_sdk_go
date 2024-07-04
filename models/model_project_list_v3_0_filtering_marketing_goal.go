/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30FilteringMarketingGoal
type ProjectListV30FilteringMarketingGoal string

// List of project_list_v3.0_filtering_marketing_goal
const (
	LIVE_ProjectListV30FilteringMarketingGoal            ProjectListV30FilteringMarketingGoal = "LIVE"
	VIDEO_AND_IMAGE_ProjectListV30FilteringMarketingGoal ProjectListV30FilteringMarketingGoal = "VIDEO_AND_IMAGE"
)

// All allowed values of ProjectListV30FilteringMarketingGoal enum
var AllowedProjectListV30FilteringMarketingGoalEnumValues = []ProjectListV30FilteringMarketingGoal{
	"LIVE",
	"VIDEO_AND_IMAGE",
}

// NewProjectListV30FilteringMarketingGoalFromValue returns a pointer to a valid ProjectListV30FilteringMarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30FilteringMarketingGoalFromValue(v string) (*ProjectListV30FilteringMarketingGoal, error) {
	ev := ProjectListV30FilteringMarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30FilteringMarketingGoal: valid values are %v", v, AllowedProjectListV30FilteringMarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30FilteringMarketingGoal) IsValid() bool {
	for _, existing := range AllowedProjectListV30FilteringMarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_filtering_marketing_goal value
func (v ProjectListV30FilteringMarketingGoal) Ptr() *ProjectListV30FilteringMarketingGoal {
	return &v
}
