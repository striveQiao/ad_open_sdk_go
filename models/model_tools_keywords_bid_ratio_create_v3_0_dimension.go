/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsKeywordsBidRatioCreateV30Dimension
type ToolsKeywordsBidRatioCreateV30Dimension string

// List of tools_keywords_bid_ratio_create_v3.0_dimension
const (
	ADVERTISER_ToolsKeywordsBidRatioCreateV30Dimension ToolsKeywordsBidRatioCreateV30Dimension = "ADVERTISER"
	PROJECT_ToolsKeywordsBidRatioCreateV30Dimension    ToolsKeywordsBidRatioCreateV30Dimension = "PROJECT"
)

// All allowed values of ToolsKeywordsBidRatioCreateV30Dimension enum
var AllowedToolsKeywordsBidRatioCreateV30DimensionEnumValues = []ToolsKeywordsBidRatioCreateV30Dimension{
	"ADVERTISER",
	"PROJECT",
}

// NewToolsKeywordsBidRatioCreateV30DimensionFromValue returns a pointer to a valid ToolsKeywordsBidRatioCreateV30Dimension
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsKeywordsBidRatioCreateV30DimensionFromValue(v string) (*ToolsKeywordsBidRatioCreateV30Dimension, error) {
	ev := ToolsKeywordsBidRatioCreateV30Dimension(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsKeywordsBidRatioCreateV30Dimension: valid values are %v", v, AllowedToolsKeywordsBidRatioCreateV30DimensionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsKeywordsBidRatioCreateV30Dimension) IsValid() bool {
	for _, existing := range AllowedToolsKeywordsBidRatioCreateV30DimensionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_keywords_bid_ratio_create_v3.0_dimension value
func (v ToolsKeywordsBidRatioCreateV30Dimension) Ptr() *ToolsKeywordsBidRatioCreateV30Dimension {
	return &v
}
