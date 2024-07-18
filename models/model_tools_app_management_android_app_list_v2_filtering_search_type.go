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

// ToolsAppManagementAndroidAppListV2FilteringSearchType
type ToolsAppManagementAndroidAppListV2FilteringSearchType string

// List of tools_app_management_android_app_list_v2_filtering_search_type
const (
	ALL_ToolsAppManagementAndroidAppListV2FilteringSearchType         ToolsAppManagementAndroidAppListV2FilteringSearchType = "ALL"
	CREATE_ONLY_ToolsAppManagementAndroidAppListV2FilteringSearchType ToolsAppManagementAndroidAppListV2FilteringSearchType = "CREATE_ONLY"
	SHARE_ONLY_ToolsAppManagementAndroidAppListV2FilteringSearchType  ToolsAppManagementAndroidAppListV2FilteringSearchType = "SHARE_ONLY"
)

// All allowed values of ToolsAppManagementAndroidAppListV2FilteringSearchType enum
var AllowedToolsAppManagementAndroidAppListV2FilteringSearchTypeEnumValues = []ToolsAppManagementAndroidAppListV2FilteringSearchType{
	"ALL",
	"CREATE_ONLY",
	"SHARE_ONLY",
}

// NewToolsAppManagementAndroidAppListV2FilteringSearchTypeFromValue returns a pointer to a valid ToolsAppManagementAndroidAppListV2FilteringSearchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementAndroidAppListV2FilteringSearchTypeFromValue(v string) (*ToolsAppManagementAndroidAppListV2FilteringSearchType, error) {
	ev := ToolsAppManagementAndroidAppListV2FilteringSearchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementAndroidAppListV2FilteringSearchType: valid values are %v", v, AllowedToolsAppManagementAndroidAppListV2FilteringSearchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementAndroidAppListV2FilteringSearchType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementAndroidAppListV2FilteringSearchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_android_app_list_v2_filtering_search_type value
func (v ToolsAppManagementAndroidAppListV2FilteringSearchType) Ptr() *ToolsAppManagementAndroidAppListV2FilteringSearchType {
	return &v
}
