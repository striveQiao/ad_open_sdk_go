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

// ToolsNoBidSuggestBidV2FilteringFlowControlMode
type ToolsNoBidSuggestBidV2FilteringFlowControlMode string

// List of tools_no_bid_suggest_bid_v2_filtering_flow_control_mode
const (
	FLOW_CONTROL_MODE_SMOOTH_ToolsNoBidSuggestBidV2FilteringFlowControlMode     ToolsNoBidSuggestBidV2FilteringFlowControlMode = "FLOW_CONTROL_MODE_SMOOTH"
	FLOW_CONTROL_MODE_TWO_PHASES_ToolsNoBidSuggestBidV2FilteringFlowControlMode ToolsNoBidSuggestBidV2FilteringFlowControlMode = "FLOW_CONTROL_MODE_TWO_PHASES"
	FLOW_CONTROL_MODE_BALANCE_ToolsNoBidSuggestBidV2FilteringFlowControlMode    ToolsNoBidSuggestBidV2FilteringFlowControlMode = "FLOW_CONTROL_MODE_BALANCE"
	FLOW_CONTROL_MODE_HOURLY_ToolsNoBidSuggestBidV2FilteringFlowControlMode     ToolsNoBidSuggestBidV2FilteringFlowControlMode = "FLOW_CONTROL_MODE_HOURLY"
	FLOW_CONTROL_MODE_FAST_ToolsNoBidSuggestBidV2FilteringFlowControlMode       ToolsNoBidSuggestBidV2FilteringFlowControlMode = "FLOW_CONTROL_MODE_FAST"
)

// All allowed values of ToolsNoBidSuggestBidV2FilteringFlowControlMode enum
var AllowedToolsNoBidSuggestBidV2FilteringFlowControlModeEnumValues = []ToolsNoBidSuggestBidV2FilteringFlowControlMode{
	"FLOW_CONTROL_MODE_SMOOTH",
	"FLOW_CONTROL_MODE_TWO_PHASES",
	"FLOW_CONTROL_MODE_BALANCE",
	"FLOW_CONTROL_MODE_HOURLY",
	"FLOW_CONTROL_MODE_FAST",
}

// NewToolsNoBidSuggestBidV2FilteringFlowControlModeFromValue returns a pointer to a valid ToolsNoBidSuggestBidV2FilteringFlowControlMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsNoBidSuggestBidV2FilteringFlowControlModeFromValue(v string) (*ToolsNoBidSuggestBidV2FilteringFlowControlMode, error) {
	ev := ToolsNoBidSuggestBidV2FilteringFlowControlMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsNoBidSuggestBidV2FilteringFlowControlMode: valid values are %v", v, AllowedToolsNoBidSuggestBidV2FilteringFlowControlModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsNoBidSuggestBidV2FilteringFlowControlMode) IsValid() bool {
	for _, existing := range AllowedToolsNoBidSuggestBidV2FilteringFlowControlModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_no_bid_suggest_bid_v2_filtering_flow_control_mode value
func (v ToolsNoBidSuggestBidV2FilteringFlowControlMode) Ptr() *ToolsNoBidSuggestBidV2FilteringFlowControlMode {
	return &v
}
