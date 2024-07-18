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

// ToolsCommentGetV30DataCommentListCommentType
type ToolsCommentGetV30DataCommentListCommentType string

// List of tools_comment_get_v3.0_data_comment_list_comment_type
const (
	IMAGE_COMMENT_ToolsCommentGetV30DataCommentListCommentType      ToolsCommentGetV30DataCommentListCommentType = "IMAGE_COMMENT"
	IMAGE_TEXT_COMMENT_ToolsCommentGetV30DataCommentListCommentType ToolsCommentGetV30DataCommentListCommentType = "IMAGE_TEXT_COMMENT"
	TEXT_COMMENT_ToolsCommentGetV30DataCommentListCommentType       ToolsCommentGetV30DataCommentListCommentType = "TEXT_COMMENT"
)

// All allowed values of ToolsCommentGetV30DataCommentListCommentType enum
var AllowedToolsCommentGetV30DataCommentListCommentTypeEnumValues = []ToolsCommentGetV30DataCommentListCommentType{
	"IMAGE_COMMENT",
	"IMAGE_TEXT_COMMENT",
	"TEXT_COMMENT",
}

// NewToolsCommentGetV30DataCommentListCommentTypeFromValue returns a pointer to a valid ToolsCommentGetV30DataCommentListCommentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsCommentGetV30DataCommentListCommentTypeFromValue(v string) (*ToolsCommentGetV30DataCommentListCommentType, error) {
	ev := ToolsCommentGetV30DataCommentListCommentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsCommentGetV30DataCommentListCommentType: valid values are %v", v, AllowedToolsCommentGetV30DataCommentListCommentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsCommentGetV30DataCommentListCommentType) IsValid() bool {
	for _, existing := range AllowedToolsCommentGetV30DataCommentListCommentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_comment_get_v3.0_data_comment_list_comment_type value
func (v ToolsCommentGetV30DataCommentListCommentType) Ptr() *ToolsCommentGetV30DataCommentListCommentType {
	return &v
}
