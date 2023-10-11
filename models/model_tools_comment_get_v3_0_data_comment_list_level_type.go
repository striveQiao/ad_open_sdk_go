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

// ToolsCommentGetV30DataCommentListLevelType
type ToolsCommentGetV30DataCommentListLevelType string

// List of tools_comment_get_v3.0_data_comment_list_level_type
const (
	LEVEL_ALL_ToolsCommentGetV30DataCommentListLevelType ToolsCommentGetV30DataCommentListLevelType = "LEVEL_ALL"
	LEVEL_ONE_ToolsCommentGetV30DataCommentListLevelType ToolsCommentGetV30DataCommentListLevelType = "LEVEL_ONE"
	LEVEL_TWO_ToolsCommentGetV30DataCommentListLevelType ToolsCommentGetV30DataCommentListLevelType = "LEVEL_TWO"
)

// All allowed values of ToolsCommentGetV30DataCommentListLevelType enum
var AllowedToolsCommentGetV30DataCommentListLevelTypeEnumValues = []ToolsCommentGetV30DataCommentListLevelType{
	"LEVEL_ALL",
	"LEVEL_ONE",
	"LEVEL_TWO",
}

// NewToolsCommentGetV30DataCommentListLevelTypeFromValue returns a pointer to a valid ToolsCommentGetV30DataCommentListLevelType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsCommentGetV30DataCommentListLevelTypeFromValue(v string) (*ToolsCommentGetV30DataCommentListLevelType, error) {
	ev := ToolsCommentGetV30DataCommentListLevelType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsCommentGetV30DataCommentListLevelType: valid values are %v", v, AllowedToolsCommentGetV30DataCommentListLevelTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsCommentGetV30DataCommentListLevelType) IsValid() bool {
	for _, existing := range AllowedToolsCommentGetV30DataCommentListLevelTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_comment_get_v3.0_data_comment_list_level_type value
func (v ToolsCommentGetV30DataCommentListLevelType) Ptr() *ToolsCommentGetV30DataCommentListLevelType {
	return &v
}
