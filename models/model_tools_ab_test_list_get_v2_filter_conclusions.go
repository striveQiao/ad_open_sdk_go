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

// ToolsAbTestListGetV2FilterConclusions
type ToolsAbTestListGetV2FilterConclusions string

// List of tools_ab_test_list_get_v2_filter_conclusions
const (
	FAILED_ToolsAbTestListGetV2FilterConclusions       ToolsAbTestListGetV2FilterConclusions = "FAILED"
	INSUFFICIENT_ToolsAbTestListGetV2FilterConclusions ToolsAbTestListGetV2FilterConclusions = "INSUFFICIENT"
	NOT_START_ToolsAbTestListGetV2FilterConclusions    ToolsAbTestListGetV2FilterConclusions = "NOT_START"
	OBVIOUS_ToolsAbTestListGetV2FilterConclusions      ToolsAbTestListGetV2FilterConclusions = "OBVIOUS"
	TMP_OBVIOUS_ToolsAbTestListGetV2FilterConclusions  ToolsAbTestListGetV2FilterConclusions = "TMP_OBVIOUS"
)

// All allowed values of ToolsAbTestListGetV2FilterConclusions enum
var AllowedToolsAbTestListGetV2FilterConclusionsEnumValues = []ToolsAbTestListGetV2FilterConclusions{
	"FAILED",
	"INSUFFICIENT",
	"NOT_START",
	"OBVIOUS",
	"TMP_OBVIOUS",
}

// NewToolsAbTestListGetV2FilterConclusionsFromValue returns a pointer to a valid ToolsAbTestListGetV2FilterConclusions
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAbTestListGetV2FilterConclusionsFromValue(v string) (*ToolsAbTestListGetV2FilterConclusions, error) {
	ev := ToolsAbTestListGetV2FilterConclusions(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAbTestListGetV2FilterConclusions: valid values are %v", v, AllowedToolsAbTestListGetV2FilterConclusionsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAbTestListGetV2FilterConclusions) IsValid() bool {
	for _, existing := range AllowedToolsAbTestListGetV2FilterConclusionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ab_test_list_get_v2_filter_conclusions value
func (v ToolsAbTestListGetV2FilterConclusions) Ptr() *ToolsAbTestListGetV2FilterConclusions {
	return &v
}
