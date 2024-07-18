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

// ToolsClueLifeGetV2DataListActionType
type ToolsClueLifeGetV2DataListActionType string

// List of tools_clue_life_get_v2_data_list_action_type
const (
	SHORT_VIDEO_ToolsClueLifeGetV2DataListActionType   ToolsClueLifeGetV2DataListActionType = "SHORT_VIDEO"
	LIVE_VIDEO_ToolsClueLifeGetV2DataListActionType    ToolsClueLifeGetV2DataListActionType = "LIVE_VIDEO"
	HOME_PAGE_ToolsClueLifeGetV2DataListActionType     ToolsClueLifeGetV2DataListActionType = "HOME_PAGE"
	IM_MESSAGE_ToolsClueLifeGetV2DataListActionType    ToolsClueLifeGetV2DataListActionType = "IM_MESSAGE"
	GROUPON_ORDER_ToolsClueLifeGetV2DataListActionType ToolsClueLifeGetV2DataListActionType = "GROUPON_ORDER"
	ALIEN_CARD_ToolsClueLifeGetV2DataListActionType    ToolsClueLifeGetV2DataListActionType = "ALIEN_CARD"
	OTHERS_ToolsClueLifeGetV2DataListActionType        ToolsClueLifeGetV2DataListActionType = "OTHERS"
)

// All allowed values of ToolsClueLifeGetV2DataListActionType enum
var AllowedToolsClueLifeGetV2DataListActionTypeEnumValues = []ToolsClueLifeGetV2DataListActionType{
	"SHORT_VIDEO",
	"LIVE_VIDEO",
	"HOME_PAGE",
	"IM_MESSAGE",
	"GROUPON_ORDER",
	"ALIEN_CARD",
	"OTHERS",
}

// NewToolsClueLifeGetV2DataListActionTypeFromValue returns a pointer to a valid ToolsClueLifeGetV2DataListActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsClueLifeGetV2DataListActionTypeFromValue(v string) (*ToolsClueLifeGetV2DataListActionType, error) {
	ev := ToolsClueLifeGetV2DataListActionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsClueLifeGetV2DataListActionType: valid values are %v", v, AllowedToolsClueLifeGetV2DataListActionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsClueLifeGetV2DataListActionType) IsValid() bool {
	for _, existing := range AllowedToolsClueLifeGetV2DataListActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_clue_life_get_v2_data_list_action_type value
func (v ToolsClueLifeGetV2DataListActionType) Ptr() *ToolsClueLifeGetV2DataListActionType {
	return &v
}
