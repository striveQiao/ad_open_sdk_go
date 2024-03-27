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

// AgentAdvCostReportListQueryV2FilteringAppName
type AgentAdvCostReportListQueryV2FilteringAppName int64

// List of agent_adv_cost_report_list_query_v2_filtering_app_name
const (
	Enum_1_AgentAdvCostReportListQueryV2FilteringAppName  AgentAdvCostReportListQueryV2FilteringAppName = 1
	Enum_10_AgentAdvCostReportListQueryV2FilteringAppName AgentAdvCostReportListQueryV2FilteringAppName = 10
	Enum_11_AgentAdvCostReportListQueryV2FilteringAppName AgentAdvCostReportListQueryV2FilteringAppName = 11
	Enum_12_AgentAdvCostReportListQueryV2FilteringAppName AgentAdvCostReportListQueryV2FilteringAppName = 12
	Enum_13_AgentAdvCostReportListQueryV2FilteringAppName AgentAdvCostReportListQueryV2FilteringAppName = 13
	Enum_14_AgentAdvCostReportListQueryV2FilteringAppName AgentAdvCostReportListQueryV2FilteringAppName = 14
	Enum_15_AgentAdvCostReportListQueryV2FilteringAppName AgentAdvCostReportListQueryV2FilteringAppName = 15
	Enum_16_AgentAdvCostReportListQueryV2FilteringAppName AgentAdvCostReportListQueryV2FilteringAppName = 16
	Enum_17_AgentAdvCostReportListQueryV2FilteringAppName AgentAdvCostReportListQueryV2FilteringAppName = 17
	Enum_2_AgentAdvCostReportListQueryV2FilteringAppName  AgentAdvCostReportListQueryV2FilteringAppName = 2
	Enum_3_AgentAdvCostReportListQueryV2FilteringAppName  AgentAdvCostReportListQueryV2FilteringAppName = 3
	Enum_4_AgentAdvCostReportListQueryV2FilteringAppName  AgentAdvCostReportListQueryV2FilteringAppName = 4
	Enum_5_AgentAdvCostReportListQueryV2FilteringAppName  AgentAdvCostReportListQueryV2FilteringAppName = 5
	Enum_6_AgentAdvCostReportListQueryV2FilteringAppName  AgentAdvCostReportListQueryV2FilteringAppName = 6
	Enum_7_AgentAdvCostReportListQueryV2FilteringAppName  AgentAdvCostReportListQueryV2FilteringAppName = 7
	Enum_8_AgentAdvCostReportListQueryV2FilteringAppName  AgentAdvCostReportListQueryV2FilteringAppName = 8
	Enum_9_AgentAdvCostReportListQueryV2FilteringAppName  AgentAdvCostReportListQueryV2FilteringAppName = 9
)

// All allowed values of AgentAdvCostReportListQueryV2FilteringAppName enum
var AllowedAgentAdvCostReportListQueryV2FilteringAppNameEnumValues = []AgentAdvCostReportListQueryV2FilteringAppName{
	1,
	10,
	11,
	12,
	13,
	14,
	15,
	16,
	17,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
}

// NewAgentAdvCostReportListQueryV2FilteringAppNameFromValue returns a pointer to a valid AgentAdvCostReportListQueryV2FilteringAppName
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentAdvCostReportListQueryV2FilteringAppNameFromValue(v int64) (*AgentAdvCostReportListQueryV2FilteringAppName, error) {
	ev := AgentAdvCostReportListQueryV2FilteringAppName(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentAdvCostReportListQueryV2FilteringAppName: valid values are %v", v, AllowedAgentAdvCostReportListQueryV2FilteringAppNameEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentAdvCostReportListQueryV2FilteringAppName) IsValid() bool {
	for _, existing := range AllowedAgentAdvCostReportListQueryV2FilteringAppNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_adv_cost_report_list_query_v2_filtering_app_name value
func (v AgentAdvCostReportListQueryV2FilteringAppName) Ptr() *AgentAdvCostReportListQueryV2FilteringAppName {
	return &v
}
