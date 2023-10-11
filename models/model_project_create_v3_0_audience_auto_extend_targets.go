/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30AudienceAutoExtendTargets
type ProjectCreateV30AudienceAutoExtendTargets string

// List of project_create_v3.0_audience_auto_extend_targets
const (
	AD_TAG_ProjectCreateV30AudienceAutoExtendTargets          ProjectCreateV30AudienceAutoExtendTargets = "AD_TAG"
	AGE_ProjectCreateV30AudienceAutoExtendTargets             ProjectCreateV30AudienceAutoExtendTargets = "AGE"
	CUSTOM_AUDIENCE_ProjectCreateV30AudienceAutoExtendTargets ProjectCreateV30AudienceAutoExtendTargets = "CUSTOM_AUDIENCE"
	GENDER_ProjectCreateV30AudienceAutoExtendTargets          ProjectCreateV30AudienceAutoExtendTargets = "GENDER"
	INTEREST_ACTION_ProjectCreateV30AudienceAutoExtendTargets ProjectCreateV30AudienceAutoExtendTargets = "INTEREST_ACTION"
	INTEREST_TAG_ProjectCreateV30AudienceAutoExtendTargets    ProjectCreateV30AudienceAutoExtendTargets = "INTEREST_TAG"
	REGION_ProjectCreateV30AudienceAutoExtendTargets          ProjectCreateV30AudienceAutoExtendTargets = "REGION"
)

// All allowed values of ProjectCreateV30AudienceAutoExtendTargets enum
var AllowedProjectCreateV30AudienceAutoExtendTargetsEnumValues = []ProjectCreateV30AudienceAutoExtendTargets{
	"AD_TAG",
	"AGE",
	"CUSTOM_AUDIENCE",
	"GENDER",
	"INTEREST_ACTION",
	"INTEREST_TAG",
	"REGION",
}

// NewProjectCreateV30AudienceAutoExtendTargetsFromValue returns a pointer to a valid ProjectCreateV30AudienceAutoExtendTargets
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AudienceAutoExtendTargetsFromValue(v string) (*ProjectCreateV30AudienceAutoExtendTargets, error) {
	ev := ProjectCreateV30AudienceAutoExtendTargets(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AudienceAutoExtendTargets: valid values are %v", v, AllowedProjectCreateV30AudienceAutoExtendTargetsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AudienceAutoExtendTargets) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AudienceAutoExtendTargetsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_audience_auto_extend_targets value
func (v ProjectCreateV30AudienceAutoExtendTargets) Ptr() *ProjectCreateV30AudienceAutoExtendTargets {
	return &v
}
