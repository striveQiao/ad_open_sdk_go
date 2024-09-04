/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentInfoV2DataAccountStatus
type AgentInfoV2DataAccountStatus string

// List of agent_info_v2_data_account_status
const (
	STATUS_DISABLE_AgentInfoV2DataAccountStatus                AgentInfoV2DataAccountStatus = "STATUS_DISABLE"
	STATUS_SELF_SERVICE_UNAUDITED_AgentInfoV2DataAccountStatus AgentInfoV2DataAccountStatus = "STATUS_SELF_SERVICE_UNAUDITED"
	STATUS_PENDING_VERIFIED_AgentInfoV2DataAccountStatus       AgentInfoV2DataAccountStatus = "STATUS_PENDING_VERIFIED"
	STATUS_LIMIT_AgentInfoV2DataAccountStatus                  AgentInfoV2DataAccountStatus = "STATUS_LIMIT"
	STATUS_PENDING_CONFIRM_MODIFY_AgentInfoV2DataAccountStatus AgentInfoV2DataAccountStatus = "STATUS_PENDING_CONFIRM_MODIFY"
	STATUS_CONFIRM_FAIL_AgentInfoV2DataAccountStatus           AgentInfoV2DataAccountStatus = "STATUS_CONFIRM_FAIL"
	STATUS_CONFIRM_MODIFY_FAIL_AgentInfoV2DataAccountStatus    AgentInfoV2DataAccountStatus = "STATUS_CONFIRM_MODIFY_FAIL"
	STATUS_ENABLE_AgentInfoV2DataAccountStatus                 AgentInfoV2DataAccountStatus = "STATUS_ENABLE"
	STATUS_PENDING_CONFIRM_AgentInfoV2DataAccountStatus        AgentInfoV2DataAccountStatus = "STATUS_PENDING_CONFIRM"
	STATUS_WAIT_FOR_PUBLIC_AUTH_AgentInfoV2DataAccountStatus   AgentInfoV2DataAccountStatus = "STATUS_WAIT_FOR_PUBLIC_AUTH"
	STATUS_CONFIRM_FAIL_END_AgentInfoV2DataAccountStatus       AgentInfoV2DataAccountStatus = "STATUS_CONFIRM_FAIL_END"
	STATUS_WAIT_FOR_BPM_AUDIT_AgentInfoV2DataAccountStatus     AgentInfoV2DataAccountStatus = "STATUS_WAIT_FOR_BPM_AUDIT"
)

// Ptr returns reference to agent_info_v2_data_account_status value
func (v AgentInfoV2DataAccountStatus) Ptr() *AgentInfoV2DataAccountStatus {
	return &v
}
