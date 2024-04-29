/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanCreativeGetV10FilteringStatus
type QianchuanCreativeGetV10FilteringStatus string

// List of qianchuan_creative_get_v1.0_filtering_status
const (
	ADVERTISER_OFFLINE_BUDGET_QianchuanCreativeGetV10FilteringStatus QianchuanCreativeGetV10FilteringStatus = "ADVERTISER_OFFLINE_BUDGET"
	AD_DISABLE_QianchuanCreativeGetV10FilteringStatus                QianchuanCreativeGetV10FilteringStatus = "AD_DISABLE"
	AD_OFFLINE_BUDGET_QianchuanCreativeGetV10FilteringStatus         QianchuanCreativeGetV10FilteringStatus = "AD_OFFLINE_BUDGET"
	ALL_INCLUDE_DELETED_QianchuanCreativeGetV10FilteringStatus       QianchuanCreativeGetV10FilteringStatus = "ALL_INCLUDE_DELETED"
	AUDIT_QianchuanCreativeGetV10FilteringStatus                     QianchuanCreativeGetV10FilteringStatus = "AUDIT"
	ALL_QianchuanCreativeGetV10FilteringStatus                       QianchuanCreativeGetV10FilteringStatus = "All"
	CAMPAIGN_DISABLE_QianchuanCreativeGetV10FilteringStatus          QianchuanCreativeGetV10FilteringStatus = "CAMPAIGN_DISABLE"
	CAMPAIGN_OFFLINE_BUDGET_QianchuanCreativeGetV10FilteringStatus   QianchuanCreativeGetV10FilteringStatus = "CAMPAIGN_OFFLINE_BUDGET"
	DELETED_QianchuanCreativeGetV10FilteringStatus                   QianchuanCreativeGetV10FilteringStatus = "DELETED"
	DELIVERY_OK_QianchuanCreativeGetV10FilteringStatus               QianchuanCreativeGetV10FilteringStatus = "DELIVERY_OK"
	DISABLE_QianchuanCreativeGetV10FilteringStatus                   QianchuanCreativeGetV10FilteringStatus = "DISABLE"
	EXTERNAL_URL_DISABLE_QianchuanCreativeGetV10FilteringStatus      QianchuanCreativeGetV10FilteringStatus = "EXTERNAL_URL_DISABLE"
	FROZEN_QianchuanCreativeGetV10FilteringStatus                    QianchuanCreativeGetV10FilteringStatus = "FROZEN"
	LIVE_ROOM_OFF_QianchuanCreativeGetV10FilteringStatus             QianchuanCreativeGetV10FilteringStatus = "LIVE_ROOM_OFF"
	NO_SCHEDULE_QianchuanCreativeGetV10FilteringStatus               QianchuanCreativeGetV10FilteringStatus = "NO_SCHEDULE"
	OFFLINE_AUDIT_QianchuanCreativeGetV10FilteringStatus             QianchuanCreativeGetV10FilteringStatus = "OFFLINE_AUDIT"
	OFFLINE_BALANCE_QianchuanCreativeGetV10FilteringStatus           QianchuanCreativeGetV10FilteringStatus = "OFFLINE_BALANCE"
	REAUDIT_QianchuanCreativeGetV10FilteringStatus                   QianchuanCreativeGetV10FilteringStatus = "REAUDIT"
	TIME_DONE_QianchuanCreativeGetV10FilteringStatus                 QianchuanCreativeGetV10FilteringStatus = "TIME_DONE"
	TIME_NO_REACH_QianchuanCreativeGetV10FilteringStatus             QianchuanCreativeGetV10FilteringStatus = "TIME_NO_REACH"
)

// All allowed values of QianchuanCreativeGetV10FilteringStatus enum
var AllowedQianchuanCreativeGetV10FilteringStatusEnumValues = []QianchuanCreativeGetV10FilteringStatus{
	"ADVERTISER_OFFLINE_BUDGET",
	"AD_DISABLE",
	"AD_OFFLINE_BUDGET",
	"ALL_INCLUDE_DELETED",
	"AUDIT",
	"All",
	"CAMPAIGN_DISABLE",
	"CAMPAIGN_OFFLINE_BUDGET",
	"DELETED",
	"DELIVERY_OK",
	"DISABLE",
	"EXTERNAL_URL_DISABLE",
	"FROZEN",
	"LIVE_ROOM_OFF",
	"NO_SCHEDULE",
	"OFFLINE_AUDIT",
	"OFFLINE_BALANCE",
	"REAUDIT",
	"TIME_DONE",
	"TIME_NO_REACH",
}

// NewQianchuanCreativeGetV10FilteringStatusFromValue returns a pointer to a valid QianchuanCreativeGetV10FilteringStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCreativeGetV10FilteringStatusFromValue(v string) (*QianchuanCreativeGetV10FilteringStatus, error) {
	ev := QianchuanCreativeGetV10FilteringStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCreativeGetV10FilteringStatus: valid values are %v", v, AllowedQianchuanCreativeGetV10FilteringStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCreativeGetV10FilteringStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanCreativeGetV10FilteringStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_creative_get_v1.0_filtering_status value
func (v QianchuanCreativeGetV10FilteringStatus) Ptr() *QianchuanCreativeGetV10FilteringStatus {
	return &v
}
