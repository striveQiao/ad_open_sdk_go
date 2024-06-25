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

// StarReportOrderUserDistributionGetV2InteractiveType
type StarReportOrderUserDistributionGetV2InteractiveType string

// List of star_report_order_user_distribution_get_v2_interactive_type
const (
	VV_StarReportOrderUserDistributionGetV2InteractiveType        StarReportOrderUserDistributionGetV2InteractiveType = "VV"
	SHARE_StarReportOrderUserDistributionGetV2InteractiveType     StarReportOrderUserDistributionGetV2InteractiveType = "SHARE"
	LIKE_StarReportOrderUserDistributionGetV2InteractiveType      StarReportOrderUserDistributionGetV2InteractiveType = "LIKE"
	COMMENT_StarReportOrderUserDistributionGetV2InteractiveType   StarReportOrderUserDistributionGetV2InteractiveType = "COMMENT"
	COMPONENT_StarReportOrderUserDistributionGetV2InteractiveType StarReportOrderUserDistributionGetV2InteractiveType = "COMPONENT"
)

// All allowed values of StarReportOrderUserDistributionGetV2InteractiveType enum
var AllowedStarReportOrderUserDistributionGetV2InteractiveTypeEnumValues = []StarReportOrderUserDistributionGetV2InteractiveType{
	"VV",
	"SHARE",
	"LIKE",
	"COMMENT",
	"COMPONENT",
}

// NewStarReportOrderUserDistributionGetV2InteractiveTypeFromValue returns a pointer to a valid StarReportOrderUserDistributionGetV2InteractiveType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarReportOrderUserDistributionGetV2InteractiveTypeFromValue(v string) (*StarReportOrderUserDistributionGetV2InteractiveType, error) {
	ev := StarReportOrderUserDistributionGetV2InteractiveType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarReportOrderUserDistributionGetV2InteractiveType: valid values are %v", v, AllowedStarReportOrderUserDistributionGetV2InteractiveTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarReportOrderUserDistributionGetV2InteractiveType) IsValid() bool {
	for _, existing := range AllowedStarReportOrderUserDistributionGetV2InteractiveTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_report_order_user_distribution_get_v2_interactive_type value
func (v StarReportOrderUserDistributionGetV2InteractiveType) Ptr() *StarReportOrderUserDistributionGetV2InteractiveType {
	return &v
}
