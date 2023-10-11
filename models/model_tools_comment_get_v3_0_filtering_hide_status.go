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

// ToolsCommentGetV30FilteringHideStatus
type ToolsCommentGetV30FilteringHideStatus string

// List of tools_comment_get_v3.0_filtering_hide_status
const (
	ALL_ToolsCommentGetV30FilteringHideStatus      ToolsCommentGetV30FilteringHideStatus = "ALL"
	HIDE_ToolsCommentGetV30FilteringHideStatus     ToolsCommentGetV30FilteringHideStatus = "HIDE"
	NOT_HIDE_ToolsCommentGetV30FilteringHideStatus ToolsCommentGetV30FilteringHideStatus = "NOT_HIDE"
)

// All allowed values of ToolsCommentGetV30FilteringHideStatus enum
var AllowedToolsCommentGetV30FilteringHideStatusEnumValues = []ToolsCommentGetV30FilteringHideStatus{
	"ALL",
	"HIDE",
	"NOT_HIDE",
}

// NewToolsCommentGetV30FilteringHideStatusFromValue returns a pointer to a valid ToolsCommentGetV30FilteringHideStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsCommentGetV30FilteringHideStatusFromValue(v string) (*ToolsCommentGetV30FilteringHideStatus, error) {
	ev := ToolsCommentGetV30FilteringHideStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsCommentGetV30FilteringHideStatus: valid values are %v", v, AllowedToolsCommentGetV30FilteringHideStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsCommentGetV30FilteringHideStatus) IsValid() bool {
	for _, existing := range AllowedToolsCommentGetV30FilteringHideStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_comment_get_v3.0_filtering_hide_status value
func (v ToolsCommentGetV30FilteringHideStatus) Ptr() *ToolsCommentGetV30FilteringHideStatus {
	return &v
}
