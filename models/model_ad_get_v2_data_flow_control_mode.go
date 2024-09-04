/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2DataFlowControlMode
type AdGetV2DataFlowControlMode string

// List of ad_get_v2_data_flow_control_mode
const (
	FLOW_CONTROL_MODE_TWO_PHASES_AdGetV2DataFlowControlMode AdGetV2DataFlowControlMode = "FLOW_CONTROL_MODE_TWO_PHASES"
	FLOW_CONTROL_MODE_BALANCE_AdGetV2DataFlowControlMode    AdGetV2DataFlowControlMode = "FLOW_CONTROL_MODE_BALANCE"
	FLOW_CONTROL_MODE_FAST_AdGetV2DataFlowControlMode       AdGetV2DataFlowControlMode = "FLOW_CONTROL_MODE_FAST"
	FLOW_CONTROL_MODE_SMOOTH_AdGetV2DataFlowControlMode     AdGetV2DataFlowControlMode = "FLOW_CONTROL_MODE_SMOOTH"
	FLOW_CONTROL_MODE_HOURLY_AdGetV2DataFlowControlMode     AdGetV2DataFlowControlMode = "FLOW_CONTROL_MODE_HOURLY"
)

// Ptr returns reference to ad_get_v2_data_flow_control_mode value
func (v AdGetV2DataFlowControlMode) Ptr() *AdGetV2DataFlowControlMode {
	return &v
}
