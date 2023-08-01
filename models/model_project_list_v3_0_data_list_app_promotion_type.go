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

// ProjectListV30DataListAppPromotionType
type ProjectListV30DataListAppPromotionType string

// List of project_list_v3.0_data_list_app_promotion_type
const (
	DOWNLOAD_ProjectListV30DataListAppPromotionType ProjectListV30DataListAppPromotionType = "DOWNLOAD"
	LAUNCH_ProjectListV30DataListAppPromotionType   ProjectListV30DataListAppPromotionType = "LAUNCH"
	RESERVE_ProjectListV30DataListAppPromotionType  ProjectListV30DataListAppPromotionType = "RESERVE"
)

// All allowed values of ProjectListV30DataListAppPromotionType enum
var AllowedProjectListV30DataListAppPromotionTypeEnumValues = []ProjectListV30DataListAppPromotionType{
	"DOWNLOAD",
	"LAUNCH",
	"RESERVE",
}

// NewProjectListV30DataListAppPromotionTypeFromValue returns a pointer to a valid ProjectListV30DataListAppPromotionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAppPromotionTypeFromValue(v string) (*ProjectListV30DataListAppPromotionType, error) {
	ev := ProjectListV30DataListAppPromotionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAppPromotionType: valid values are %v", v, AllowedProjectListV30DataListAppPromotionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAppPromotionType) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAppPromotionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_app_promotion_type value
func (v ProjectListV30DataListAppPromotionType) Ptr() *ProjectListV30DataListAppPromotionType {
	return &v
}
