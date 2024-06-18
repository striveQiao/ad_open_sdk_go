/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListAudienceDeviceType
type ProjectListV30DataListAudienceDeviceType string

// List of project_list_v3.0_data_list_audience_device_type
const (
	MOBILE_ProjectListV30DataListAudienceDeviceType ProjectListV30DataListAudienceDeviceType = "MOBILE"
	PAD_ProjectListV30DataListAudienceDeviceType    ProjectListV30DataListAudienceDeviceType = "PAD"
)

// All allowed values of ProjectListV30DataListAudienceDeviceType enum
var AllowedProjectListV30DataListAudienceDeviceTypeEnumValues = []ProjectListV30DataListAudienceDeviceType{
	"MOBILE",
	"PAD",
}

// NewProjectListV30DataListAudienceDeviceTypeFromValue returns a pointer to a valid ProjectListV30DataListAudienceDeviceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAudienceDeviceTypeFromValue(v string) (*ProjectListV30DataListAudienceDeviceType, error) {
	ev := ProjectListV30DataListAudienceDeviceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAudienceDeviceType: valid values are %v", v, AllowedProjectListV30DataListAudienceDeviceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAudienceDeviceType) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAudienceDeviceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_audience_device_type value
func (v ProjectListV30DataListAudienceDeviceType) Ptr() *ProjectListV30DataListAudienceDeviceType {
	return &v
}
