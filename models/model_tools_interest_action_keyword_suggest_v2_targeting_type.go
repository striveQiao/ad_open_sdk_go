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

// ToolsInterestActionKeywordSuggestV2TargetingType
type ToolsInterestActionKeywordSuggestV2TargetingType string

// List of tools_interest_action_keyword_suggest_v2_targeting_type
const (
	ACTION_ToolsInterestActionKeywordSuggestV2TargetingType   ToolsInterestActionKeywordSuggestV2TargetingType = "ACTION"
	INTEREST_ToolsInterestActionKeywordSuggestV2TargetingType ToolsInterestActionKeywordSuggestV2TargetingType = "INTEREST"
)

// All allowed values of ToolsInterestActionKeywordSuggestV2TargetingType enum
var AllowedToolsInterestActionKeywordSuggestV2TargetingTypeEnumValues = []ToolsInterestActionKeywordSuggestV2TargetingType{
	"ACTION",
	"INTEREST",
}

// NewToolsInterestActionKeywordSuggestV2TargetingTypeFromValue returns a pointer to a valid ToolsInterestActionKeywordSuggestV2TargetingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsInterestActionKeywordSuggestV2TargetingTypeFromValue(v string) (*ToolsInterestActionKeywordSuggestV2TargetingType, error) {
	ev := ToolsInterestActionKeywordSuggestV2TargetingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsInterestActionKeywordSuggestV2TargetingType: valid values are %v", v, AllowedToolsInterestActionKeywordSuggestV2TargetingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsInterestActionKeywordSuggestV2TargetingType) IsValid() bool {
	for _, existing := range AllowedToolsInterestActionKeywordSuggestV2TargetingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_interest_action_keyword_suggest_v2_targeting_type value
func (v ToolsInterestActionKeywordSuggestV2TargetingType) Ptr() *ToolsInterestActionKeywordSuggestV2TargetingType {
	return &v
}
