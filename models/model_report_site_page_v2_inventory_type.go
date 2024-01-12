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

// ReportSitePageV2InventoryType
type ReportSitePageV2InventoryType string

// List of report_site_page_v2_inventory_type
const (
	WATERMELON_ReportSitePageV2InventoryType ReportSitePageV2InventoryType = "WATERMELON"
	AWEME_ReportSitePageV2InventoryType      ReportSitePageV2InventoryType = "AWEME"
	TOUTIAO_ReportSitePageV2InventoryType    ReportSitePageV2InventoryType = "TOUTIAO"
	UNION_SLOT_ReportSitePageV2InventoryType ReportSitePageV2InventoryType = "UNION_SLOT"
	HOTSOON_ReportSitePageV2InventoryType    ReportSitePageV2InventoryType = "HOTSOON"
)

// All allowed values of ReportSitePageV2InventoryType enum
var AllowedReportSitePageV2InventoryTypeEnumValues = []ReportSitePageV2InventoryType{
	"WATERMELON",
	"AWEME",
	"TOUTIAO",
	"UNION_SLOT",
	"HOTSOON",
}

// NewReportSitePageV2InventoryTypeFromValue returns a pointer to a valid ReportSitePageV2InventoryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportSitePageV2InventoryTypeFromValue(v string) (*ReportSitePageV2InventoryType, error) {
	ev := ReportSitePageV2InventoryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportSitePageV2InventoryType: valid values are %v", v, AllowedReportSitePageV2InventoryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportSitePageV2InventoryType) IsValid() bool {
	for _, existing := range AllowedReportSitePageV2InventoryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_site_page_v2_inventory_type value
func (v ReportSitePageV2InventoryType) Ptr() *ReportSitePageV2InventoryType {
	return &v
}
