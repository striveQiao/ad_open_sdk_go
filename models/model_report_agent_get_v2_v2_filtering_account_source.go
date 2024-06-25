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

// ReportAgentGetV2V2FilteringAccountSource
type ReportAgentGetV2V2FilteringAccountSource string

// List of report_agent_get_v2_v2_filtering_account_source
const (
	LUBAN_ACCOUNT_ReportAgentGetV2V2FilteringAccountSource     ReportAgentGetV2V2FilteringAccountSource = "LUBAN_ACCOUNT"
	NORMAL_ADVERTISER_ReportAgentGetV2V2FilteringAccountSource ReportAgentGetV2V2FilteringAccountSource = "NORMAL_ADVERTISER"
)

// All allowed values of ReportAgentGetV2V2FilteringAccountSource enum
var AllowedReportAgentGetV2V2FilteringAccountSourceEnumValues = []ReportAgentGetV2V2FilteringAccountSource{
	"LUBAN_ACCOUNT",
	"NORMAL_ADVERTISER",
}

// NewReportAgentGetV2V2FilteringAccountSourceFromValue returns a pointer to a valid ReportAgentGetV2V2FilteringAccountSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAgentGetV2V2FilteringAccountSourceFromValue(v string) (*ReportAgentGetV2V2FilteringAccountSource, error) {
	ev := ReportAgentGetV2V2FilteringAccountSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAgentGetV2V2FilteringAccountSource: valid values are %v", v, AllowedReportAgentGetV2V2FilteringAccountSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAgentGetV2V2FilteringAccountSource) IsValid() bool {
	for _, existing := range AllowedReportAgentGetV2V2FilteringAccountSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_agent_get_v2_v2_filtering_account_source value
func (v ReportAgentGetV2V2FilteringAccountSource) Ptr() *ReportAgentGetV2V2FilteringAccountSource {
	return &v
}
