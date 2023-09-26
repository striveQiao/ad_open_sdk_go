/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListStatusSecond
type ProjectListV30DataListStatusSecond string

// List of project_list_v3.0_data_list_status_second
const (
	PROJECT_STATUS_BUDGET_EXCEED_ProjectListV30DataListStatusSecond ProjectListV30DataListStatusSecond = "PROJECT_STATUS_BUDGET_EXCEED"
	PROJECT_STATUS_NOT_START_ProjectListV30DataListStatusSecond     ProjectListV30DataListStatusSecond = "PROJECT_STATUS_NOT_START"
	PROJECT_STATUS_NO_SCHEDULE_ProjectListV30DataListStatusSecond   ProjectListV30DataListStatusSecond = "PROJECT_STATUS_NO_SCHEDULE"
	PROJECT_STATUS_STOP_ProjectListV30DataListStatusSecond          ProjectListV30DataListStatusSecond = "PROJECT_STATUS_STOP"
)

// All allowed values of ProjectListV30DataListStatusSecond enum
var AllowedProjectListV30DataListStatusSecondEnumValues = []ProjectListV30DataListStatusSecond{
	"PROJECT_STATUS_BUDGET_EXCEED",
	"PROJECT_STATUS_NOT_START",
	"PROJECT_STATUS_NO_SCHEDULE",
	"PROJECT_STATUS_STOP",
}

// NewProjectListV30DataListStatusSecondFromValue returns a pointer to a valid ProjectListV30DataListStatusSecond
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListStatusSecondFromValue(v string) (*ProjectListV30DataListStatusSecond, error) {
	ev := ProjectListV30DataListStatusSecond(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListStatusSecond: valid values are %v", v, AllowedProjectListV30DataListStatusSecondEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListStatusSecond) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListStatusSecondEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_status_second value
func (v ProjectListV30DataListStatusSecond) Ptr() *ProjectListV30DataListStatusSecond {
	return &v
}
