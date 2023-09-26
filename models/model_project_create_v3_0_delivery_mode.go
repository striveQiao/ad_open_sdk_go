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

// ProjectCreateV30DeliveryMode
type ProjectCreateV30DeliveryMode string

// List of project_create_v3.0_delivery_mode
const (
	MANUAL_ProjectCreateV30DeliveryMode     ProjectCreateV30DeliveryMode = "MANUAL"
	PROCEDURAL_ProjectCreateV30DeliveryMode ProjectCreateV30DeliveryMode = "PROCEDURAL"
)

// All allowed values of ProjectCreateV30DeliveryMode enum
var AllowedProjectCreateV30DeliveryModeEnumValues = []ProjectCreateV30DeliveryMode{
	"MANUAL",
	"PROCEDURAL",
}

// NewProjectCreateV30DeliveryModeFromValue returns a pointer to a valid ProjectCreateV30DeliveryMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30DeliveryModeFromValue(v string) (*ProjectCreateV30DeliveryMode, error) {
	ev := ProjectCreateV30DeliveryMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30DeliveryMode: valid values are %v", v, AllowedProjectCreateV30DeliveryModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30DeliveryMode) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30DeliveryModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_delivery_mode value
func (v ProjectCreateV30DeliveryMode) Ptr() *ProjectCreateV30DeliveryMode {
	return &v
}
