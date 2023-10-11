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

// ToolsAppManagementBookingGetV2SearchType
type ToolsAppManagementBookingGetV2SearchType string

// List of tools_app_management_booking_get_v2_search_type
const (
	ALL_ToolsAppManagementBookingGetV2SearchType         ToolsAppManagementBookingGetV2SearchType = "ALL"
	CREATE_ONLY_ToolsAppManagementBookingGetV2SearchType ToolsAppManagementBookingGetV2SearchType = "CREATE_ONLY"
	SHARED_ONLY_ToolsAppManagementBookingGetV2SearchType ToolsAppManagementBookingGetV2SearchType = "SHARED_ONLY"
)

// All allowed values of ToolsAppManagementBookingGetV2SearchType enum
var AllowedToolsAppManagementBookingGetV2SearchTypeEnumValues = []ToolsAppManagementBookingGetV2SearchType{
	"ALL",
	"CREATE_ONLY",
	"SHARED_ONLY",
}

// NewToolsAppManagementBookingGetV2SearchTypeFromValue returns a pointer to a valid ToolsAppManagementBookingGetV2SearchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementBookingGetV2SearchTypeFromValue(v string) (*ToolsAppManagementBookingGetV2SearchType, error) {
	ev := ToolsAppManagementBookingGetV2SearchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementBookingGetV2SearchType: valid values are %v", v, AllowedToolsAppManagementBookingGetV2SearchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementBookingGetV2SearchType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementBookingGetV2SearchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_booking_get_v2_search_type value
func (v ToolsAppManagementBookingGetV2SearchType) Ptr() *ToolsAppManagementBookingGetV2SearchType {
	return &v
}
