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

// ToolsAbTestListGetV2DataTestListConclusion
type ToolsAbTestListGetV2DataTestListConclusion string

// List of tools_ab_test_list_get_v2_data_test_list_conclusion
const (
	FAILED_ToolsAbTestListGetV2DataTestListConclusion       ToolsAbTestListGetV2DataTestListConclusion = "FAILED"
	INSUFFICIENT_ToolsAbTestListGetV2DataTestListConclusion ToolsAbTestListGetV2DataTestListConclusion = "INSUFFICIENT"
	NOT_START_ToolsAbTestListGetV2DataTestListConclusion    ToolsAbTestListGetV2DataTestListConclusion = "NOT_START"
	OBVIOUS_ToolsAbTestListGetV2DataTestListConclusion      ToolsAbTestListGetV2DataTestListConclusion = "OBVIOUS"
	TMP_OBVIOUS_ToolsAbTestListGetV2DataTestListConclusion  ToolsAbTestListGetV2DataTestListConclusion = "TMP_OBVIOUS"
)

// All allowed values of ToolsAbTestListGetV2DataTestListConclusion enum
var AllowedToolsAbTestListGetV2DataTestListConclusionEnumValues = []ToolsAbTestListGetV2DataTestListConclusion{
	"FAILED",
	"INSUFFICIENT",
	"NOT_START",
	"OBVIOUS",
	"TMP_OBVIOUS",
}

// NewToolsAbTestListGetV2DataTestListConclusionFromValue returns a pointer to a valid ToolsAbTestListGetV2DataTestListConclusion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAbTestListGetV2DataTestListConclusionFromValue(v string) (*ToolsAbTestListGetV2DataTestListConclusion, error) {
	ev := ToolsAbTestListGetV2DataTestListConclusion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAbTestListGetV2DataTestListConclusion: valid values are %v", v, AllowedToolsAbTestListGetV2DataTestListConclusionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAbTestListGetV2DataTestListConclusion) IsValid() bool {
	for _, existing := range AllowedToolsAbTestListGetV2DataTestListConclusionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ab_test_list_get_v2_data_test_list_conclusion value
func (v ToolsAbTestListGetV2DataTestListConclusion) Ptr() *ToolsAbTestListGetV2DataTestListConclusion {
	return &v
}
