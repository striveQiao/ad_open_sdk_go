/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPlayableListGetV2Status
type ToolsPlayableListGetV2Status string

// List of tools_playable_list_get_v2_status
const (
	VALIDATING_ToolsPlayableListGetV2Status       ToolsPlayableListGetV2Status = "VALIDATING"
	AUDIT_FAIL_ToolsPlayableListGetV2Status       ToolsPlayableListGetV2Status = "AUDIT_FAIL"
	VALIDATE_SUCCESS_ToolsPlayableListGetV2Status ToolsPlayableListGetV2Status = "VALIDATE_SUCCESS"
	AUDIT_SUCCESS_ToolsPlayableListGetV2Status    ToolsPlayableListGetV2Status = "AUDIT_SUCCESS"
	VALIDATE_FAIL_ToolsPlayableListGetV2Status    ToolsPlayableListGetV2Status = "VALIDATE_FAIL"
)

// Ptr returns reference to tools_playable_list_get_v2_status value
func (v ToolsPlayableListGetV2Status) Ptr() *ToolsPlayableListGetV2Status {
	return &v
}
