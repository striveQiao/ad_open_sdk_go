/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30DeliveryType
type ProjectCreateV30DeliveryType string

// List of project_create_v3.0_delivery_type
const (
	DURATION_ProjectCreateV30DeliveryType ProjectCreateV30DeliveryType = "DURATION"
	NORMAL_ProjectCreateV30DeliveryType   ProjectCreateV30DeliveryType = "NORMAL"
)

// All allowed values of ProjectCreateV30DeliveryType enum
var AllowedProjectCreateV30DeliveryTypeEnumValues = []ProjectCreateV30DeliveryType{
	"DURATION",
	"NORMAL",
}

// NewProjectCreateV30DeliveryTypeFromValue returns a pointer to a valid ProjectCreateV30DeliveryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30DeliveryTypeFromValue(v string) (*ProjectCreateV30DeliveryType, error) {
	ev := ProjectCreateV30DeliveryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30DeliveryType: valid values are %v", v, AllowedProjectCreateV30DeliveryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30DeliveryType) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30DeliveryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_delivery_type value
func (v ProjectCreateV30DeliveryType) Ptr() *ProjectCreateV30DeliveryType {
	return &v
}
