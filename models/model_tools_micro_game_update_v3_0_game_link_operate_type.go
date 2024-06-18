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

// ToolsMicroGameUpdateV30GameLinkOperateType
type ToolsMicroGameUpdateV30GameLinkOperateType string

// List of tools_micro_game_update_v3.0_game_link_operate_type
const (
	NEW_ToolsMicroGameUpdateV30GameLinkOperateType    ToolsMicroGameUpdateV30GameLinkOperateType = "NEW"
	MODIFY_ToolsMicroGameUpdateV30GameLinkOperateType ToolsMicroGameUpdateV30GameLinkOperateType = "MODIFY"
	DELETE_ToolsMicroGameUpdateV30GameLinkOperateType ToolsMicroGameUpdateV30GameLinkOperateType = "DELETE"
)

// All allowed values of ToolsMicroGameUpdateV30GameLinkOperateType enum
var AllowedToolsMicroGameUpdateV30GameLinkOperateTypeEnumValues = []ToolsMicroGameUpdateV30GameLinkOperateType{
	"NEW",
	"MODIFY",
	"DELETE",
}

// NewToolsMicroGameUpdateV30GameLinkOperateTypeFromValue returns a pointer to a valid ToolsMicroGameUpdateV30GameLinkOperateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsMicroGameUpdateV30GameLinkOperateTypeFromValue(v string) (*ToolsMicroGameUpdateV30GameLinkOperateType, error) {
	ev := ToolsMicroGameUpdateV30GameLinkOperateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsMicroGameUpdateV30GameLinkOperateType: valid values are %v", v, AllowedToolsMicroGameUpdateV30GameLinkOperateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsMicroGameUpdateV30GameLinkOperateType) IsValid() bool {
	for _, existing := range AllowedToolsMicroGameUpdateV30GameLinkOperateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_micro_game_update_v3.0_game_link_operate_type value
func (v ToolsMicroGameUpdateV30GameLinkOperateType) Ptr() *ToolsMicroGameUpdateV30GameLinkOperateType {
	return &v
}
