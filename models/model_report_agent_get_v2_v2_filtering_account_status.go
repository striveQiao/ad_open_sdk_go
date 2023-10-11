/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportAgentGetV2V2FilteringAccountStatus
type ReportAgentGetV2V2FilteringAccountStatus string

// List of report_agent_get_v2_v2_filtering_account_status
const (
	STATUS_ENABLE_AND_AVATAR_AUDITING_ReportAgentGetV2V2FilteringAccountStatus ReportAgentGetV2V2FilteringAccountStatus = "STATUS_ENABLE_AND_AVATAR_AUDITING"
	STATUS_PENDING_CONFIRM_MODIFY_ReportAgentGetV2V2FilteringAccountStatus     ReportAgentGetV2V2FilteringAccountStatus = "STATUS_PENDING_CONFIRM_MODIFY"
	STATUS_PENDING_CONFIRM_ReportAgentGetV2V2FilteringAccountStatus            ReportAgentGetV2V2FilteringAccountStatus = "STATUS_PENDING_CONFIRM"
	STATUS_WAIT_FOR_BPM_FILE_CONTACT_ReportAgentGetV2V2FilteringAccountStatus  ReportAgentGetV2V2FilteringAccountStatus = "STATUS_WAIT_FOR_BPM_FILE_CONTACT"
	STATUS_PUNISH_ReportAgentGetV2V2FilteringAccountStatus                     ReportAgentGetV2V2FilteringAccountStatus = "STATUS_PUNISH"
	STATUS_CONFIRM_FAIL_ReportAgentGetV2V2FilteringAccountStatus               ReportAgentGetV2V2FilteringAccountStatus = "STATUS_CONFIRM_FAIL"
	STATUS_WAIT_FOR_BPM_AUDIT_ReportAgentGetV2V2FilteringAccountStatus         ReportAgentGetV2V2FilteringAccountStatus = "STATUS_WAIT_FOR_BPM_AUDIT"
	STATUS_PENDING_VERIFIED_ReportAgentGetV2V2FilteringAccountStatus           ReportAgentGetV2V2FilteringAccountStatus = "STATUS_PENDING_VERIFIED"
	STATUS_SELF_SERVICE_UNAUDITED_ReportAgentGetV2V2FilteringAccountStatus     ReportAgentGetV2V2FilteringAccountStatus = "STATUS_SELF_SERVICE_UNAUDITED"
	STATUS_DISABLE_ReportAgentGetV2V2FilteringAccountStatus                    ReportAgentGetV2V2FilteringAccountStatus = "STATUS_DISABLE"
	STATUS_ENABLE_ReportAgentGetV2V2FilteringAccountStatus                     ReportAgentGetV2V2FilteringAccountStatus = "STATUS_ENABLE"
	STATUS_CONFIRM_FAIL_END_ReportAgentGetV2V2FilteringAccountStatus           ReportAgentGetV2V2FilteringAccountStatus = "STATUS_CONFIRM_FAIL_END"
	STATUS_WAIT_FOR_ACCOUNT_FEE_ReportAgentGetV2V2FilteringAccountStatus       ReportAgentGetV2V2FilteringAccountStatus = "STATUS_WAIT_FOR_ACCOUNT_FEE"
	STATUS_CONFIRM_MODIFY_FAIL_ReportAgentGetV2V2FilteringAccountStatus        ReportAgentGetV2V2FilteringAccountStatus = "STATUS_CONFIRM_MODIFY_FAIL"
	STATUS_WAIT_FOR_PUBLIC_AUTH_ReportAgentGetV2V2FilteringAccountStatus       ReportAgentGetV2V2FilteringAccountStatus = "STATUS_WAIT_FOR_PUBLIC_AUTH"
)

// All allowed values of ReportAgentGetV2V2FilteringAccountStatus enum
var AllowedReportAgentGetV2V2FilteringAccountStatusEnumValues = []ReportAgentGetV2V2FilteringAccountStatus{
	"STATUS_ENABLE_AND_AVATAR_AUDITING",
	"STATUS_PENDING_CONFIRM_MODIFY",
	"STATUS_PENDING_CONFIRM",
	"STATUS_WAIT_FOR_BPM_FILE_CONTACT",
	"STATUS_PUNISH",
	"STATUS_CONFIRM_FAIL",
	"STATUS_WAIT_FOR_BPM_AUDIT",
	"STATUS_PENDING_VERIFIED",
	"STATUS_SELF_SERVICE_UNAUDITED",
	"STATUS_DISABLE",
	"STATUS_ENABLE",
	"STATUS_CONFIRM_FAIL_END",
	"STATUS_WAIT_FOR_ACCOUNT_FEE",
	"STATUS_CONFIRM_MODIFY_FAIL",
	"STATUS_WAIT_FOR_PUBLIC_AUTH",
}

// NewReportAgentGetV2V2FilteringAccountStatusFromValue returns a pointer to a valid ReportAgentGetV2V2FilteringAccountStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAgentGetV2V2FilteringAccountStatusFromValue(v string) (*ReportAgentGetV2V2FilteringAccountStatus, error) {
	ev := ReportAgentGetV2V2FilteringAccountStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAgentGetV2V2FilteringAccountStatus: valid values are %v", v, AllowedReportAgentGetV2V2FilteringAccountStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAgentGetV2V2FilteringAccountStatus) IsValid() bool {
	for _, existing := range AllowedReportAgentGetV2V2FilteringAccountStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_agent_get_v2_v2_filtering_account_status value
func (v ReportAgentGetV2V2FilteringAccountStatus) Ptr() *ReportAgentGetV2V2FilteringAccountStatus {
	return &v
}
