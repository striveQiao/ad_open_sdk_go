/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsMicroAppListV30FilteringSearchType
type ToolsMicroAppListV30FilteringSearchType string

// List of tools_micro_app_list_v3.0_filtering_search_type
const (
	CREATE_ONLY_ToolsMicroAppListV30FilteringSearchType ToolsMicroAppListV30FilteringSearchType = "CREATE_ONLY"
	SHARE_ONLY_ToolsMicroAppListV30FilteringSearchType  ToolsMicroAppListV30FilteringSearchType = "SHARE_ONLY"
)

// All allowed values of ToolsMicroAppListV30FilteringSearchType enum
var AllowedToolsMicroAppListV30FilteringSearchTypeEnumValues = []ToolsMicroAppListV30FilteringSearchType{
	"CREATE_ONLY",
	"SHARE_ONLY",
}

// NewToolsMicroAppListV30FilteringSearchTypeFromValue returns a pointer to a valid ToolsMicroAppListV30FilteringSearchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsMicroAppListV30FilteringSearchTypeFromValue(v string) (*ToolsMicroAppListV30FilteringSearchType, error) {
	ev := ToolsMicroAppListV30FilteringSearchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsMicroAppListV30FilteringSearchType: valid values are %v", v, AllowedToolsMicroAppListV30FilteringSearchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsMicroAppListV30FilteringSearchType) IsValid() bool {
	for _, existing := range AllowedToolsMicroAppListV30FilteringSearchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_micro_app_list_v3.0_filtering_search_type value
func (v ToolsMicroAppListV30FilteringSearchType) Ptr() *ToolsMicroAppListV30FilteringSearchType {
	return &v
}
