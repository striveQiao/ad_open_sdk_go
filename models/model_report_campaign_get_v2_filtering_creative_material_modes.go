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

// ReportCampaignGetV2FilteringCreativeMaterialModes
type ReportCampaignGetV2FilteringCreativeMaterialModes string

// List of report_campaign_get_v2_filtering_creative_material_modes
const (
	INTERVENE_ReportCampaignGetV2FilteringCreativeMaterialModes       ReportCampaignGetV2FilteringCreativeMaterialModes = "INTERVENE"
	CTR_ReportCampaignGetV2FilteringCreativeMaterialModes             ReportCampaignGetV2FilteringCreativeMaterialModes = "CTR"
	STATIC_ASSEMBLE_ReportCampaignGetV2FilteringCreativeMaterialModes ReportCampaignGetV2FilteringCreativeMaterialModes = "STATIC_ASSEMBLE"
)

// All allowed values of ReportCampaignGetV2FilteringCreativeMaterialModes enum
var AllowedReportCampaignGetV2FilteringCreativeMaterialModesEnumValues = []ReportCampaignGetV2FilteringCreativeMaterialModes{
	"INTERVENE",
	"CTR",
	"STATIC_ASSEMBLE",
}

// NewReportCampaignGetV2FilteringCreativeMaterialModesFromValue returns a pointer to a valid ReportCampaignGetV2FilteringCreativeMaterialModes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCampaignGetV2FilteringCreativeMaterialModesFromValue(v string) (*ReportCampaignGetV2FilteringCreativeMaterialModes, error) {
	ev := ReportCampaignGetV2FilteringCreativeMaterialModes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCampaignGetV2FilteringCreativeMaterialModes: valid values are %v", v, AllowedReportCampaignGetV2FilteringCreativeMaterialModesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCampaignGetV2FilteringCreativeMaterialModes) IsValid() bool {
	for _, existing := range AllowedReportCampaignGetV2FilteringCreativeMaterialModesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_campaign_get_v2_filtering_creative_material_modes value
func (v ReportCampaignGetV2FilteringCreativeMaterialModes) Ptr() *ReportCampaignGetV2FilteringCreativeMaterialModes {
	return &v
}
