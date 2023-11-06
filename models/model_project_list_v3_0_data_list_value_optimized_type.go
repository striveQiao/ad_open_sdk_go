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

// ProjectListV30DataListValueOptimizedType
type ProjectListV30DataListValueOptimizedType string

// List of project_list_v3.0_data_list_value_optimized_type
const (
	ACTION_ProjectListV30DataListValueOptimizedType ProjectListV30DataListValueOptimizedType = "ACTION"
	OFF_ProjectListV30DataListValueOptimizedType    ProjectListV30DataListValueOptimizedType = "OFF"
	VALUE_ProjectListV30DataListValueOptimizedType  ProjectListV30DataListValueOptimizedType = "VALUE"
)

// All allowed values of ProjectListV30DataListValueOptimizedType enum
var AllowedProjectListV30DataListValueOptimizedTypeEnumValues = []ProjectListV30DataListValueOptimizedType{
	"ACTION",
	"OFF",
	"VALUE",
}

// NewProjectListV30DataListValueOptimizedTypeFromValue returns a pointer to a valid ProjectListV30DataListValueOptimizedType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListValueOptimizedTypeFromValue(v string) (*ProjectListV30DataListValueOptimizedType, error) {
	ev := ProjectListV30DataListValueOptimizedType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListValueOptimizedType: valid values are %v", v, AllowedProjectListV30DataListValueOptimizedTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListValueOptimizedType) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListValueOptimizedTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_value_optimized_type value
func (v ProjectListV30DataListValueOptimizedType) Ptr() *ProjectListV30DataListValueOptimizedType {
	return &v
}
