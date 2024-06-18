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

// AgentAdvCostReportListQueryV2FilteringSpuLabelName
type AgentAdvCostReportListQueryV2FilteringSpuLabelName string

// List of agent_adv_cost_report_list_query_v2_filtering_spu_label_name
const (
	BRAND_CONVERSION_AgentAdvCostReportListQueryV2FilteringSpuLabelName       AgentAdvCostReportListQueryV2FilteringSpuLabelName = "BRAND_CONVERSION"
	BRAND_EXPOSURE_AgentAdvCostReportListQueryV2FilteringSpuLabelName         AgentAdvCostReportListQueryV2FilteringSpuLabelName = "BRAND_EXPOSURE"
	BRAND_GRASS_AgentAdvCostReportListQueryV2FilteringSpuLabelName            AgentAdvCostReportListQueryV2FilteringSpuLabelName = "BRAND_GRASS"
	CONCENTRATING_RESOURCE_AgentAdvCostReportListQueryV2FilteringSpuLabelName AgentAdvCostReportListQueryV2FilteringSpuLabelName = "CONCENTRATING_RESOURCE"
	STAR_RESOURCE_AgentAdvCostReportListQueryV2FilteringSpuLabelName          AgentAdvCostReportListQueryV2FilteringSpuLabelName = "STAR_RESOURCE"
)

// All allowed values of AgentAdvCostReportListQueryV2FilteringSpuLabelName enum
var AllowedAgentAdvCostReportListQueryV2FilteringSpuLabelNameEnumValues = []AgentAdvCostReportListQueryV2FilteringSpuLabelName{
	"BRAND_CONVERSION",
	"BRAND_EXPOSURE",
	"BRAND_GRASS",
	"CONCENTRATING_RESOURCE",
	"STAR_RESOURCE",
}

// NewAgentAdvCostReportListQueryV2FilteringSpuLabelNameFromValue returns a pointer to a valid AgentAdvCostReportListQueryV2FilteringSpuLabelName
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentAdvCostReportListQueryV2FilteringSpuLabelNameFromValue(v string) (*AgentAdvCostReportListQueryV2FilteringSpuLabelName, error) {
	ev := AgentAdvCostReportListQueryV2FilteringSpuLabelName(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentAdvCostReportListQueryV2FilteringSpuLabelName: valid values are %v", v, AllowedAgentAdvCostReportListQueryV2FilteringSpuLabelNameEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentAdvCostReportListQueryV2FilteringSpuLabelName) IsValid() bool {
	for _, existing := range AllowedAgentAdvCostReportListQueryV2FilteringSpuLabelNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_adv_cost_report_list_query_v2_filtering_spu_label_name value
func (v AgentAdvCostReportListQueryV2FilteringSpuLabelName) Ptr() *AgentAdvCostReportListQueryV2FilteringSpuLabelName {
	return &v
}
