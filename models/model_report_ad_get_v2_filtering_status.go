/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportAdGetV2FilteringStatus
type ReportAdGetV2FilteringStatus string

// List of report_ad_get_v2_filtering_status
const (
	AD_STATUS_AUDIT_ReportAdGetV2FilteringStatus            ReportAdGetV2FilteringStatus = "AD_STATUS_AUDIT"
	AD_STATUS_AUDIT_DENY_ReportAdGetV2FilteringStatus       ReportAdGetV2FilteringStatus = "AD_STATUS_AUDIT_DENY"
	AD_STATUS_DISABLE_ReportAdGetV2FilteringStatus          ReportAdGetV2FilteringStatus = "AD_STATUS_DISABLE"
	AD_STATUS_BUDGET_EXCEED_ReportAdGetV2FilteringStatus    ReportAdGetV2FilteringStatus = "AD_STATUS_BUDGET_EXCEED"
	AD_STATUS_ALL_ReportAdGetV2FilteringStatus              ReportAdGetV2FilteringStatus = "AD_STATUS_ALL"
	AD_STATUS_NOT_START_ReportAdGetV2FilteringStatus        ReportAdGetV2FilteringStatus = "AD_STATUS_NOT_START"
	AD_STATUS_BALANCE_EXCEED_ReportAdGetV2FilteringStatus   ReportAdGetV2FilteringStatus = "AD_STATUS_BALANCE_EXCEED"
	AD_STATUS_DELETE_ReportAdGetV2FilteringStatus           ReportAdGetV2FilteringStatus = "AD_STATUS_DELETE"
	AD_STATUS_NO_SCHEDULE_ReportAdGetV2FilteringStatus      ReportAdGetV2FilteringStatus = "AD_STATUS_NO_SCHEDULE"
	AD_STATUS_CAMPAIGN_DISABLE_ReportAdGetV2FilteringStatus ReportAdGetV2FilteringStatus = "AD_STATUS_CAMPAIGN_DISABLE"
	AD_STATUS_CAMPAIGN_EXCEED_ReportAdGetV2FilteringStatus  ReportAdGetV2FilteringStatus = "AD_STATUS_CAMPAIGN_EXCEED"
	AD_STATUS_REAUDIT_ReportAdGetV2FilteringStatus          ReportAdGetV2FilteringStatus = "AD_STATUS_REAUDIT"
	AD_STATUS_DELIVERY_OK_ReportAdGetV2FilteringStatus      ReportAdGetV2FilteringStatus = "AD_STATUS_DELIVERY_OK"
	AD_STATUS_NOT_DELETE_ReportAdGetV2FilteringStatus       ReportAdGetV2FilteringStatus = "AD_STATUS_NOT_DELETE"
	AD_STATUS_DONE_ReportAdGetV2FilteringStatus             ReportAdGetV2FilteringStatus = "AD_STATUS_DONE"
	AD_STATUS_CREATE_ReportAdGetV2FilteringStatus           ReportAdGetV2FilteringStatus = "AD_STATUS_CREATE"
)

// All allowed values of ReportAdGetV2FilteringStatus enum
var AllowedReportAdGetV2FilteringStatusEnumValues = []ReportAdGetV2FilteringStatus{
	"AD_STATUS_AUDIT",
	"AD_STATUS_AUDIT_DENY",
	"AD_STATUS_DISABLE",
	"AD_STATUS_BUDGET_EXCEED",
	"AD_STATUS_ALL",
	"AD_STATUS_NOT_START",
	"AD_STATUS_BALANCE_EXCEED",
	"AD_STATUS_DELETE",
	"AD_STATUS_NO_SCHEDULE",
	"AD_STATUS_CAMPAIGN_DISABLE",
	"AD_STATUS_CAMPAIGN_EXCEED",
	"AD_STATUS_REAUDIT",
	"AD_STATUS_DELIVERY_OK",
	"AD_STATUS_NOT_DELETE",
	"AD_STATUS_DONE",
	"AD_STATUS_CREATE",
}

// NewReportAdGetV2FilteringStatusFromValue returns a pointer to a valid ReportAdGetV2FilteringStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdGetV2FilteringStatusFromValue(v string) (*ReportAdGetV2FilteringStatus, error) {
	ev := ReportAdGetV2FilteringStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdGetV2FilteringStatus: valid values are %v", v, AllowedReportAdGetV2FilteringStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdGetV2FilteringStatus) IsValid() bool {
	for _, existing := range AllowedReportAdGetV2FilteringStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_ad_get_v2_filtering_status value
func (v ReportAdGetV2FilteringStatus) Ptr() *ReportAdGetV2FilteringStatus {
	return &v
}
