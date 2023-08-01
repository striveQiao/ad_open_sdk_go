/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30FilteringAppPromotionType
type ProjectListV30FilteringAppPromotionType string

// List of project_list_v3.0_filtering_app_promotion_type
const (
	DOWNLOAD_ProjectListV30FilteringAppPromotionType ProjectListV30FilteringAppPromotionType = "DOWNLOAD"
	LAUNCH_ProjectListV30FilteringAppPromotionType   ProjectListV30FilteringAppPromotionType = "LAUNCH"
	RESERVE_ProjectListV30FilteringAppPromotionType  ProjectListV30FilteringAppPromotionType = "RESERVE"
)

// All allowed values of ProjectListV30FilteringAppPromotionType enum
var AllowedProjectListV30FilteringAppPromotionTypeEnumValues = []ProjectListV30FilteringAppPromotionType{
	"DOWNLOAD",
	"LAUNCH",
	"RESERVE",
}

// NewProjectListV30FilteringAppPromotionTypeFromValue returns a pointer to a valid ProjectListV30FilteringAppPromotionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30FilteringAppPromotionTypeFromValue(v string) (*ProjectListV30FilteringAppPromotionType, error) {
	ev := ProjectListV30FilteringAppPromotionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30FilteringAppPromotionType: valid values are %v", v, AllowedProjectListV30FilteringAppPromotionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30FilteringAppPromotionType) IsValid() bool {
	for _, existing := range AllowedProjectListV30FilteringAppPromotionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_filtering_app_promotion_type value
func (v ProjectListV30FilteringAppPromotionType) Ptr() *ProjectListV30FilteringAppPromotionType {
	return &v
}
