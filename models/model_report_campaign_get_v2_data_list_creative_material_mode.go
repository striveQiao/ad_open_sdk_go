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

// ReportCampaignGetV2DataListCreativeMaterialMode
type ReportCampaignGetV2DataListCreativeMaterialMode string

// List of report_campaign_get_v2_data_list_creative_material_mode
const (
	INTERVENE_ReportCampaignGetV2DataListCreativeMaterialMode       ReportCampaignGetV2DataListCreativeMaterialMode = "INTERVENE"
	CTR_ReportCampaignGetV2DataListCreativeMaterialMode             ReportCampaignGetV2DataListCreativeMaterialMode = "CTR"
	STATIC_ASSEMBLE_ReportCampaignGetV2DataListCreativeMaterialMode ReportCampaignGetV2DataListCreativeMaterialMode = "STATIC_ASSEMBLE"
)

// All allowed values of ReportCampaignGetV2DataListCreativeMaterialMode enum
var AllowedReportCampaignGetV2DataListCreativeMaterialModeEnumValues = []ReportCampaignGetV2DataListCreativeMaterialMode{
	"INTERVENE",
	"CTR",
	"STATIC_ASSEMBLE",
}

// NewReportCampaignGetV2DataListCreativeMaterialModeFromValue returns a pointer to a valid ReportCampaignGetV2DataListCreativeMaterialMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCampaignGetV2DataListCreativeMaterialModeFromValue(v string) (*ReportCampaignGetV2DataListCreativeMaterialMode, error) {
	ev := ReportCampaignGetV2DataListCreativeMaterialMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCampaignGetV2DataListCreativeMaterialMode: valid values are %v", v, AllowedReportCampaignGetV2DataListCreativeMaterialModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCampaignGetV2DataListCreativeMaterialMode) IsValid() bool {
	for _, existing := range AllowedReportCampaignGetV2DataListCreativeMaterialModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_campaign_get_v2_data_list_creative_material_mode value
func (v ReportCampaignGetV2DataListCreativeMaterialMode) Ptr() *ReportCampaignGetV2DataListCreativeMaterialMode {
	return &v
}
