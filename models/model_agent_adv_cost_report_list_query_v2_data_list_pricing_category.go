/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AgentAdvCostReportListQueryV2DataListPricingCategory
type AgentAdvCostReportListQueryV2DataListPricingCategory int64

// List of agent_adv_cost_report_list_query_v2_data_list_pricing_category
const (
	Enum_0_AgentAdvCostReportListQueryV2DataListPricingCategory AgentAdvCostReportListQueryV2DataListPricingCategory = 0
	Enum_1_AgentAdvCostReportListQueryV2DataListPricingCategory AgentAdvCostReportListQueryV2DataListPricingCategory = 1
	Enum_2_AgentAdvCostReportListQueryV2DataListPricingCategory AgentAdvCostReportListQueryV2DataListPricingCategory = 2
)

// All allowed values of AgentAdvCostReportListQueryV2DataListPricingCategory enum
var AllowedAgentAdvCostReportListQueryV2DataListPricingCategoryEnumValues = []AgentAdvCostReportListQueryV2DataListPricingCategory{
	0,
	1,
	2,
}

// NewAgentAdvCostReportListQueryV2DataListPricingCategoryFromValue returns a pointer to a valid AgentAdvCostReportListQueryV2DataListPricingCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentAdvCostReportListQueryV2DataListPricingCategoryFromValue(v int64) (*AgentAdvCostReportListQueryV2DataListPricingCategory, error) {
	ev := AgentAdvCostReportListQueryV2DataListPricingCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentAdvCostReportListQueryV2DataListPricingCategory: valid values are %v", v, AllowedAgentAdvCostReportListQueryV2DataListPricingCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentAdvCostReportListQueryV2DataListPricingCategory) IsValid() bool {
	for _, existing := range AllowedAgentAdvCostReportListQueryV2DataListPricingCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_adv_cost_report_list_query_v2_data_list_pricing_category value
func (v AgentAdvCostReportListQueryV2DataListPricingCategory) Ptr() *AgentAdvCostReportListQueryV2DataListPricingCategory {
	return &v
}
