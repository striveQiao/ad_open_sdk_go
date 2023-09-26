/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsCommentGetV30PlatformVersion
type ToolsCommentGetV30PlatformVersion string

// List of tools_comment_get_v3.0_platform_version
const (
	V1_ToolsCommentGetV30PlatformVersion ToolsCommentGetV30PlatformVersion = "V1"
	V2_ToolsCommentGetV30PlatformVersion ToolsCommentGetV30PlatformVersion = "V2"
)

// All allowed values of ToolsCommentGetV30PlatformVersion enum
var AllowedToolsCommentGetV30PlatformVersionEnumValues = []ToolsCommentGetV30PlatformVersion{
	"V1",
	"V2",
}

// NewToolsCommentGetV30PlatformVersionFromValue returns a pointer to a valid ToolsCommentGetV30PlatformVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsCommentGetV30PlatformVersionFromValue(v string) (*ToolsCommentGetV30PlatformVersion, error) {
	ev := ToolsCommentGetV30PlatformVersion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsCommentGetV30PlatformVersion: valid values are %v", v, AllowedToolsCommentGetV30PlatformVersionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsCommentGetV30PlatformVersion) IsValid() bool {
	for _, existing := range AllowedToolsCommentGetV30PlatformVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_comment_get_v3.0_platform_version value
func (v ToolsCommentGetV30PlatformVersion) Ptr() *ToolsCommentGetV30PlatformVersion {
	return &v
}
