/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsInterestActionKeywordSuggestV2ActionDays
type ToolsInterestActionKeywordSuggestV2ActionDays int64

// List of tools_interest_action_keyword_suggest_v2_action_days
const (
	Enum_15_ToolsInterestActionKeywordSuggestV2ActionDays  ToolsInterestActionKeywordSuggestV2ActionDays = 15
	Enum_180_ToolsInterestActionKeywordSuggestV2ActionDays ToolsInterestActionKeywordSuggestV2ActionDays = 180
	Enum_30_ToolsInterestActionKeywordSuggestV2ActionDays  ToolsInterestActionKeywordSuggestV2ActionDays = 30
	Enum_365_ToolsInterestActionKeywordSuggestV2ActionDays ToolsInterestActionKeywordSuggestV2ActionDays = 365
	Enum_60_ToolsInterestActionKeywordSuggestV2ActionDays  ToolsInterestActionKeywordSuggestV2ActionDays = 60
	Enum_7_ToolsInterestActionKeywordSuggestV2ActionDays   ToolsInterestActionKeywordSuggestV2ActionDays = 7
	Enum_90_ToolsInterestActionKeywordSuggestV2ActionDays  ToolsInterestActionKeywordSuggestV2ActionDays = 90
)

// All allowed values of ToolsInterestActionKeywordSuggestV2ActionDays enum
var AllowedToolsInterestActionKeywordSuggestV2ActionDaysEnumValues = []ToolsInterestActionKeywordSuggestV2ActionDays{
	15,
	180,
	30,
	365,
	60,
	7,
	90,
}

// NewToolsInterestActionKeywordSuggestV2ActionDaysFromValue returns a pointer to a valid ToolsInterestActionKeywordSuggestV2ActionDays
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsInterestActionKeywordSuggestV2ActionDaysFromValue(v int64) (*ToolsInterestActionKeywordSuggestV2ActionDays, error) {
	ev := ToolsInterestActionKeywordSuggestV2ActionDays(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsInterestActionKeywordSuggestV2ActionDays: valid values are %v", v, AllowedToolsInterestActionKeywordSuggestV2ActionDaysEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsInterestActionKeywordSuggestV2ActionDays) IsValid() bool {
	for _, existing := range AllowedToolsInterestActionKeywordSuggestV2ActionDaysEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_interest_action_keyword_suggest_v2_action_days value
func (v ToolsInterestActionKeywordSuggestV2ActionDays) Ptr() *ToolsInterestActionKeywordSuggestV2ActionDays {
	return &v
}
