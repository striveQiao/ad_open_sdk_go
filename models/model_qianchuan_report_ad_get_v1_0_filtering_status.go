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

// QianchuanReportAdGetV10FilteringStatus
type QianchuanReportAdGetV10FilteringStatus string

// List of qianchuan_report_ad_get_v1.0_filtering_status
const (
	ALL_INCLUDE_DELETED_QianchuanReportAdGetV10FilteringStatus  QianchuanReportAdGetV10FilteringStatus = "ALL_INCLUDE_DELETED"
	AUDIT_QianchuanReportAdGetV10FilteringStatus                QianchuanReportAdGetV10FilteringStatus = "AUDIT"
	DELETED_QianchuanReportAdGetV10FilteringStatus              QianchuanReportAdGetV10FilteringStatus = "DELETED"
	DELIVERY_OK_QianchuanReportAdGetV10FilteringStatus          QianchuanReportAdGetV10FilteringStatus = "DELIVERY_OK"
	DISABLE_QianchuanReportAdGetV10FilteringStatus              QianchuanReportAdGetV10FilteringStatus = "DISABLE"
	EXTERNAL_URL_DISABLE_QianchuanReportAdGetV10FilteringStatus QianchuanReportAdGetV10FilteringStatus = "EXTERNAL_URL_DISABLE"
	FROZEN_QianchuanReportAdGetV10FilteringStatus               QianchuanReportAdGetV10FilteringStatus = "FROZEN"
	LIVE_ROOM_OFF_QianchuanReportAdGetV10FilteringStatus        QianchuanReportAdGetV10FilteringStatus = "LIVE_ROOM_OFF"
	NO_SCHEDULE_QianchuanReportAdGetV10FilteringStatus          QianchuanReportAdGetV10FilteringStatus = "NO_SCHEDULE"
	OFFLINE_AUDIT_QianchuanReportAdGetV10FilteringStatus        QianchuanReportAdGetV10FilteringStatus = "OFFLINE_AUDIT"
	OFFLINE_BALANCE_QianchuanReportAdGetV10FilteringStatus      QianchuanReportAdGetV10FilteringStatus = "OFFLINE_BALANCE"
	OFFLINE_BUDGET_QianchuanReportAdGetV10FilteringStatus       QianchuanReportAdGetV10FilteringStatus = "OFFLINE_BUDGET"
	QUOTA_DISABLE_QianchuanReportAdGetV10FilteringStatus        QianchuanReportAdGetV10FilteringStatus = "QUOTA_DISABLE"
	REAUDIT_QianchuanReportAdGetV10FilteringStatus              QianchuanReportAdGetV10FilteringStatus = "REAUDIT"
	SYSTEM_DISABLE_QianchuanReportAdGetV10FilteringStatus       QianchuanReportAdGetV10FilteringStatus = "SYSTEM_DISABLE"
	TIME_DONE_QianchuanReportAdGetV10FilteringStatus            QianchuanReportAdGetV10FilteringStatus = "TIME_DONE"
	TIME_NO_REACH_QianchuanReportAdGetV10FilteringStatus        QianchuanReportAdGetV10FilteringStatus = "TIME_NO_REACH"
)

// All allowed values of QianchuanReportAdGetV10FilteringStatus enum
var AllowedQianchuanReportAdGetV10FilteringStatusEnumValues = []QianchuanReportAdGetV10FilteringStatus{
	"ALL_INCLUDE_DELETED",
	"AUDIT",
	"DELETED",
	"DELIVERY_OK",
	"DISABLE",
	"EXTERNAL_URL_DISABLE",
	"FROZEN",
	"LIVE_ROOM_OFF",
	"NO_SCHEDULE",
	"OFFLINE_AUDIT",
	"OFFLINE_BALANCE",
	"OFFLINE_BUDGET",
	"QUOTA_DISABLE",
	"REAUDIT",
	"SYSTEM_DISABLE",
	"TIME_DONE",
	"TIME_NO_REACH",
}

// NewQianchuanReportAdGetV10FilteringStatusFromValue returns a pointer to a valid QianchuanReportAdGetV10FilteringStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportAdGetV10FilteringStatusFromValue(v string) (*QianchuanReportAdGetV10FilteringStatus, error) {
	ev := QianchuanReportAdGetV10FilteringStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportAdGetV10FilteringStatus: valid values are %v", v, AllowedQianchuanReportAdGetV10FilteringStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportAdGetV10FilteringStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanReportAdGetV10FilteringStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_ad_get_v1.0_filtering_status value
func (v QianchuanReportAdGetV10FilteringStatus) Ptr() *QianchuanReportAdGetV10FilteringStatus {
	return &v
}
