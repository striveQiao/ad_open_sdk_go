/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsRubeexPlayableListV2DataListAdStatus
type ToolsRubeexPlayableListV2DataListAdStatus string

// List of tools_rubeex_playable_list_v2_data_list_ad_status
const (
	INUSE_ToolsRubeexPlayableListV2DataListAdStatus ToolsRubeexPlayableListV2DataListAdStatus = "INUSE"
	UNUSE_ToolsRubeexPlayableListV2DataListAdStatus ToolsRubeexPlayableListV2DataListAdStatus = "UNUSE"
)

// All allowed values of ToolsRubeexPlayableListV2DataListAdStatus enum
var AllowedToolsRubeexPlayableListV2DataListAdStatusEnumValues = []ToolsRubeexPlayableListV2DataListAdStatus{
	"INUSE",
	"UNUSE",
}

// NewToolsRubeexPlayableListV2DataListAdStatusFromValue returns a pointer to a valid ToolsRubeexPlayableListV2DataListAdStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsRubeexPlayableListV2DataListAdStatusFromValue(v string) (*ToolsRubeexPlayableListV2DataListAdStatus, error) {
	ev := ToolsRubeexPlayableListV2DataListAdStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsRubeexPlayableListV2DataListAdStatus: valid values are %v", v, AllowedToolsRubeexPlayableListV2DataListAdStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsRubeexPlayableListV2DataListAdStatus) IsValid() bool {
	for _, existing := range AllowedToolsRubeexPlayableListV2DataListAdStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_rubeex_playable_list_v2_data_list_ad_status value
func (v ToolsRubeexPlayableListV2DataListAdStatus) Ptr() *ToolsRubeexPlayableListV2DataListAdStatus {
	return &v
}
