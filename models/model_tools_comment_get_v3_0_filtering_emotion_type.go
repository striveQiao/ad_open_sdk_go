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

// ToolsCommentGetV30FilteringEmotionType
type ToolsCommentGetV30FilteringEmotionType string

// List of tools_comment_get_v3.0_filtering_emotion_type
const (
	NEGATIVE_ToolsCommentGetV30FilteringEmotionType ToolsCommentGetV30FilteringEmotionType = "NEGATIVE"
	NEUTRAL_ToolsCommentGetV30FilteringEmotionType  ToolsCommentGetV30FilteringEmotionType = "NEUTRAL"
	POSITIVE_ToolsCommentGetV30FilteringEmotionType ToolsCommentGetV30FilteringEmotionType = "POSITIVE"
)

// All allowed values of ToolsCommentGetV30FilteringEmotionType enum
var AllowedToolsCommentGetV30FilteringEmotionTypeEnumValues = []ToolsCommentGetV30FilteringEmotionType{
	"NEGATIVE",
	"NEUTRAL",
	"POSITIVE",
}

// NewToolsCommentGetV30FilteringEmotionTypeFromValue returns a pointer to a valid ToolsCommentGetV30FilteringEmotionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsCommentGetV30FilteringEmotionTypeFromValue(v string) (*ToolsCommentGetV30FilteringEmotionType, error) {
	ev := ToolsCommentGetV30FilteringEmotionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsCommentGetV30FilteringEmotionType: valid values are %v", v, AllowedToolsCommentGetV30FilteringEmotionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsCommentGetV30FilteringEmotionType) IsValid() bool {
	for _, existing := range AllowedToolsCommentGetV30FilteringEmotionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_comment_get_v3.0_filtering_emotion_type value
func (v ToolsCommentGetV30FilteringEmotionType) Ptr() *ToolsCommentGetV30FilteringEmotionType {
	return &v
}
