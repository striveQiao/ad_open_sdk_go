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

// ToolsCommentGetV30FilteringCommentType
type ToolsCommentGetV30FilteringCommentType string

// List of tools_comment_get_v3.0_filtering_comment_type
const (
	IMAGE_COMMENT_ToolsCommentGetV30FilteringCommentType      ToolsCommentGetV30FilteringCommentType = "IMAGE_COMMENT"
	IMAGE_TEXT_COMMENT_ToolsCommentGetV30FilteringCommentType ToolsCommentGetV30FilteringCommentType = "IMAGE_TEXT_COMMENT"
	TEXT_COMMENT_ToolsCommentGetV30FilteringCommentType       ToolsCommentGetV30FilteringCommentType = "TEXT_COMMENT"
)

// All allowed values of ToolsCommentGetV30FilteringCommentType enum
var AllowedToolsCommentGetV30FilteringCommentTypeEnumValues = []ToolsCommentGetV30FilteringCommentType{
	"IMAGE_COMMENT",
	"IMAGE_TEXT_COMMENT",
	"TEXT_COMMENT",
}

// NewToolsCommentGetV30FilteringCommentTypeFromValue returns a pointer to a valid ToolsCommentGetV30FilteringCommentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsCommentGetV30FilteringCommentTypeFromValue(v string) (*ToolsCommentGetV30FilteringCommentType, error) {
	ev := ToolsCommentGetV30FilteringCommentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsCommentGetV30FilteringCommentType: valid values are %v", v, AllowedToolsCommentGetV30FilteringCommentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsCommentGetV30FilteringCommentType) IsValid() bool {
	for _, existing := range AllowedToolsCommentGetV30FilteringCommentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_comment_get_v3.0_filtering_comment_type value
func (v ToolsCommentGetV30FilteringCommentType) Ptr() *ToolsCommentGetV30FilteringCommentType {
	return &v
}
