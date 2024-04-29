/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsCommentGetV30FilteringLevelType
type ToolsCommentGetV30FilteringLevelType string

// List of tools_comment_get_v3.0_filtering_level_type
const (
	LEVEL_ALL_ToolsCommentGetV30FilteringLevelType ToolsCommentGetV30FilteringLevelType = "LEVEL_ALL"
	LEVEL_ONE_ToolsCommentGetV30FilteringLevelType ToolsCommentGetV30FilteringLevelType = "LEVEL_ONE"
	LEVEL_TWO_ToolsCommentGetV30FilteringLevelType ToolsCommentGetV30FilteringLevelType = "LEVEL_TWO"
)

// All allowed values of ToolsCommentGetV30FilteringLevelType enum
var AllowedToolsCommentGetV30FilteringLevelTypeEnumValues = []ToolsCommentGetV30FilteringLevelType{
	"LEVEL_ALL",
	"LEVEL_ONE",
	"LEVEL_TWO",
}

// NewToolsCommentGetV30FilteringLevelTypeFromValue returns a pointer to a valid ToolsCommentGetV30FilteringLevelType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsCommentGetV30FilteringLevelTypeFromValue(v string) (*ToolsCommentGetV30FilteringLevelType, error) {
	ev := ToolsCommentGetV30FilteringLevelType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsCommentGetV30FilteringLevelType: valid values are %v", v, AllowedToolsCommentGetV30FilteringLevelTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsCommentGetV30FilteringLevelType) IsValid() bool {
	for _, existing := range AllowedToolsCommentGetV30FilteringLevelTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_comment_get_v3.0_filtering_level_type value
func (v ToolsCommentGetV30FilteringLevelType) Ptr() *ToolsCommentGetV30FilteringLevelType {
	return &v
}
