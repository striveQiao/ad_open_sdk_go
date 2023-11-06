/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportAgentGetV2V2DataListAccountSource
type ReportAgentGetV2V2DataListAccountSource string

// List of report_agent_get_v2_v2_data_list_account_source
const (
	LUBAN_ACCOUNT_ReportAgentGetV2V2DataListAccountSource     ReportAgentGetV2V2DataListAccountSource = "LUBAN_ACCOUNT"
	NORMAL_ADVERTISER_ReportAgentGetV2V2DataListAccountSource ReportAgentGetV2V2DataListAccountSource = "NORMAL_ADVERTISER"
)

// All allowed values of ReportAgentGetV2V2DataListAccountSource enum
var AllowedReportAgentGetV2V2DataListAccountSourceEnumValues = []ReportAgentGetV2V2DataListAccountSource{
	"LUBAN_ACCOUNT",
	"NORMAL_ADVERTISER",
}

// NewReportAgentGetV2V2DataListAccountSourceFromValue returns a pointer to a valid ReportAgentGetV2V2DataListAccountSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAgentGetV2V2DataListAccountSourceFromValue(v string) (*ReportAgentGetV2V2DataListAccountSource, error) {
	ev := ReportAgentGetV2V2DataListAccountSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAgentGetV2V2DataListAccountSource: valid values are %v", v, AllowedReportAgentGetV2V2DataListAccountSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAgentGetV2V2DataListAccountSource) IsValid() bool {
	for _, existing := range AllowedReportAgentGetV2V2DataListAccountSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_agent_get_v2_v2_data_list_account_source value
func (v ReportAgentGetV2V2DataListAccountSource) Ptr() *ReportAgentGetV2V2DataListAccountSource {
	return &v
}
