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

// ToolsUnionFlowPackagePromotionReportV30OrderType
type ToolsUnionFlowPackagePromotionReportV30OrderType string

// List of tools_union_flow_package_promotion_report_v3.0_order_type
const (
	ASC_ToolsUnionFlowPackagePromotionReportV30OrderType  ToolsUnionFlowPackagePromotionReportV30OrderType = "ASC"
	DESC_ToolsUnionFlowPackagePromotionReportV30OrderType ToolsUnionFlowPackagePromotionReportV30OrderType = "DESC"
)

// All allowed values of ToolsUnionFlowPackagePromotionReportV30OrderType enum
var AllowedToolsUnionFlowPackagePromotionReportV30OrderTypeEnumValues = []ToolsUnionFlowPackagePromotionReportV30OrderType{
	"ASC",
	"DESC",
}

// NewToolsUnionFlowPackagePromotionReportV30OrderTypeFromValue returns a pointer to a valid ToolsUnionFlowPackagePromotionReportV30OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsUnionFlowPackagePromotionReportV30OrderTypeFromValue(v string) (*ToolsUnionFlowPackagePromotionReportV30OrderType, error) {
	ev := ToolsUnionFlowPackagePromotionReportV30OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsUnionFlowPackagePromotionReportV30OrderType: valid values are %v", v, AllowedToolsUnionFlowPackagePromotionReportV30OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsUnionFlowPackagePromotionReportV30OrderType) IsValid() bool {
	for _, existing := range AllowedToolsUnionFlowPackagePromotionReportV30OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_union_flow_package_promotion_report_v3.0_order_type value
func (v ToolsUnionFlowPackagePromotionReportV30OrderType) Ptr() *ToolsUnionFlowPackagePromotionReportV30OrderType {
	return &v
}
