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

// ProjectListV30DataListDeliveryMode
type ProjectListV30DataListDeliveryMode string

// List of project_list_v3.0_data_list_delivery_mode
const (
	MANUAL_ProjectListV30DataListDeliveryMode     ProjectListV30DataListDeliveryMode = "MANUAL"
	PROCEDURAL_ProjectListV30DataListDeliveryMode ProjectListV30DataListDeliveryMode = "PROCEDURAL"
)

// All allowed values of ProjectListV30DataListDeliveryMode enum
var AllowedProjectListV30DataListDeliveryModeEnumValues = []ProjectListV30DataListDeliveryMode{
	"MANUAL",
	"PROCEDURAL",
}

// NewProjectListV30DataListDeliveryModeFromValue returns a pointer to a valid ProjectListV30DataListDeliveryMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListDeliveryModeFromValue(v string) (*ProjectListV30DataListDeliveryMode, error) {
	ev := ProjectListV30DataListDeliveryMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListDeliveryMode: valid values are %v", v, AllowedProjectListV30DataListDeliveryModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListDeliveryMode) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListDeliveryModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_delivery_mode value
func (v ProjectListV30DataListDeliveryMode) Ptr() *ProjectListV30DataListDeliveryMode {
	return &v
}
