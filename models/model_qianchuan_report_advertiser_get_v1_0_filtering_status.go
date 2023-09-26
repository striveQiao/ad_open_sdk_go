/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportAdvertiserGetV10FilteringStatus
type QianchuanReportAdvertiserGetV10FilteringStatus string

// List of qianchuan_report_advertiser_get_v1.0_filtering_status
const (
	ALL_INCLUDE_DELETED_QianchuanReportAdvertiserGetV10FilteringStatus  QianchuanReportAdvertiserGetV10FilteringStatus = "ALL_INCLUDE_DELETED"
	AUDIT_QianchuanReportAdvertiserGetV10FilteringStatus                QianchuanReportAdvertiserGetV10FilteringStatus = "AUDIT"
	DELETED_QianchuanReportAdvertiserGetV10FilteringStatus              QianchuanReportAdvertiserGetV10FilteringStatus = "DELETED"
	DELIVERY_OK_QianchuanReportAdvertiserGetV10FilteringStatus          QianchuanReportAdvertiserGetV10FilteringStatus = "DELIVERY_OK"
	DISABLE_QianchuanReportAdvertiserGetV10FilteringStatus              QianchuanReportAdvertiserGetV10FilteringStatus = "DISABLE"
	EXTERNAL_URL_DISABLE_QianchuanReportAdvertiserGetV10FilteringStatus QianchuanReportAdvertiserGetV10FilteringStatus = "EXTERNAL_URL_DISABLE"
	FROZEN_QianchuanReportAdvertiserGetV10FilteringStatus               QianchuanReportAdvertiserGetV10FilteringStatus = "FROZEN"
	LIVE_ROOM_OFF_QianchuanReportAdvertiserGetV10FilteringStatus        QianchuanReportAdvertiserGetV10FilteringStatus = "LIVE_ROOM_OFF"
	NO_SCHEDULE_QianchuanReportAdvertiserGetV10FilteringStatus          QianchuanReportAdvertiserGetV10FilteringStatus = "NO_SCHEDULE"
	OFFLINE_AUDIT_QianchuanReportAdvertiserGetV10FilteringStatus        QianchuanReportAdvertiserGetV10FilteringStatus = "OFFLINE_AUDIT"
	OFFLINE_BALANCE_QianchuanReportAdvertiserGetV10FilteringStatus      QianchuanReportAdvertiserGetV10FilteringStatus = "OFFLINE_BALANCE"
	OFFLINE_BUDGET_QianchuanReportAdvertiserGetV10FilteringStatus       QianchuanReportAdvertiserGetV10FilteringStatus = "OFFLINE_BUDGET"
	QUOTA_DISABLE_QianchuanReportAdvertiserGetV10FilteringStatus        QianchuanReportAdvertiserGetV10FilteringStatus = "QUOTA_DISABLE"
	REAUDIT_QianchuanReportAdvertiserGetV10FilteringStatus              QianchuanReportAdvertiserGetV10FilteringStatus = "REAUDIT"
	SYSTEM_DISABLE_QianchuanReportAdvertiserGetV10FilteringStatus       QianchuanReportAdvertiserGetV10FilteringStatus = "SYSTEM_DISABLE"
	TIME_DONE_QianchuanReportAdvertiserGetV10FilteringStatus            QianchuanReportAdvertiserGetV10FilteringStatus = "TIME_DONE"
	TIME_NO_REACH_QianchuanReportAdvertiserGetV10FilteringStatus        QianchuanReportAdvertiserGetV10FilteringStatus = "TIME_NO_REACH"
)

// All allowed values of QianchuanReportAdvertiserGetV10FilteringStatus enum
var AllowedQianchuanReportAdvertiserGetV10FilteringStatusEnumValues = []QianchuanReportAdvertiserGetV10FilteringStatus{
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

// NewQianchuanReportAdvertiserGetV10FilteringStatusFromValue returns a pointer to a valid QianchuanReportAdvertiserGetV10FilteringStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportAdvertiserGetV10FilteringStatusFromValue(v string) (*QianchuanReportAdvertiserGetV10FilteringStatus, error) {
	ev := QianchuanReportAdvertiserGetV10FilteringStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportAdvertiserGetV10FilteringStatus: valid values are %v", v, AllowedQianchuanReportAdvertiserGetV10FilteringStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportAdvertiserGetV10FilteringStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanReportAdvertiserGetV10FilteringStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_advertiser_get_v1.0_filtering_status value
func (v QianchuanReportAdvertiserGetV10FilteringStatus) Ptr() *QianchuanReportAdvertiserGetV10FilteringStatus {
	return &v
}
