/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportAgentGetV2V2FilteringActive
type ReportAgentGetV2V2FilteringActive string

// List of report_agent_get_v2_v2_filtering_active
const (
	ALL_ReportAgentGetV2V2FilteringActive    ReportAgentGetV2V2FilteringActive = "ALL"
	ACTIVE_ReportAgentGetV2V2FilteringActive ReportAgentGetV2V2FilteringActive = "ACTIVE"
	SILENT_ReportAgentGetV2V2FilteringActive ReportAgentGetV2V2FilteringActive = "SILENT"
)

// All allowed values of ReportAgentGetV2V2FilteringActive enum
var AllowedReportAgentGetV2V2FilteringActiveEnumValues = []ReportAgentGetV2V2FilteringActive{
	"ALL",
	"ACTIVE",
	"SILENT",
}

// NewReportAgentGetV2V2FilteringActiveFromValue returns a pointer to a valid ReportAgentGetV2V2FilteringActive
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAgentGetV2V2FilteringActiveFromValue(v string) (*ReportAgentGetV2V2FilteringActive, error) {
	ev := ReportAgentGetV2V2FilteringActive(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAgentGetV2V2FilteringActive: valid values are %v", v, AllowedReportAgentGetV2V2FilteringActiveEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAgentGetV2V2FilteringActive) IsValid() bool {
	for _, existing := range AllowedReportAgentGetV2V2FilteringActiveEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_agent_get_v2_v2_filtering_active value
func (v ReportAgentGetV2V2FilteringActive) Ptr() *ReportAgentGetV2V2FilteringActive {
	return &v
}
