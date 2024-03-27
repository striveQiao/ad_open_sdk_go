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

// PromotionListV30DataListStatus
type PromotionListV30DataListStatus string

// List of promotion_list_v3.0_data_list_status
const (
	ADVERTISER_OFFLINE_BUDGET_PromotionListV30DataListStatus    PromotionListV30DataListStatus = "ADVERTISER_OFFLINE_BUDGET"
	ADVERTISER_PREOFFLINE_BUDGET_PromotionListV30DataListStatus PromotionListV30DataListStatus = "ADVERTISER_PREOFFLINE_BUDGET"
	ALL_PromotionListV30DataListStatus                          PromotionListV30DataListStatus = "ALL"
	AUDIT_PromotionListV30DataListStatus                        PromotionListV30DataListStatus = "AUDIT"
	AUDIT_DENY_PromotionListV30DataListStatus                   PromotionListV30DataListStatus = "AUDIT_DENY"
	AWEME_ACCOUNT_DISABLED_PromotionListV30DataListStatus       PromotionListV30DataListStatus = "AWEME_ACCOUNT_DISABLED"
	AWEME_ANCHOR_DISABLED_PromotionListV30DataListStatus        PromotionListV30DataListStatus = "AWEME_ANCHOR_DISABLED"
	CREATE_PromotionListV30DataListStatus                       PromotionListV30DataListStatus = "CREATE"
	DELETED_PromotionListV30DataListStatus                      PromotionListV30DataListStatus = "DELETED"
	DISABLED_PromotionListV30DataListStatus                     PromotionListV30DataListStatus = "DISABLED"
	DISABLE_BY_QUOTA_PromotionListV30DataListStatus             PromotionListV30DataListStatus = "DISABLE_BY_QUOTA"
	ERROR_DEFAULT_PromotionListV30DataListStatus                PromotionListV30DataListStatus = "ERROR_DEFAULT"
	FROZEN_PromotionListV30DataListStatus                       PromotionListV30DataListStatus = "FROZEN"
	LIVE_ROOM_OFF_PromotionListV30DataListStatus                PromotionListV30DataListStatus = "LIVE_ROOM_OFF"
	NOT_DELETED_PromotionListV30DataListStatus                  PromotionListV30DataListStatus = "NOT_DELETED"
	NO_SCHEDULE_PromotionListV30DataListStatus                  PromotionListV30DataListStatus = "NO_SCHEDULE"
	OFFLINE_BALANCE_PromotionListV30DataListStatus              PromotionListV30DataListStatus = "OFFLINE_BALANCE"
	OFFLINE_BUDGET_PromotionListV30DataListStatus               PromotionListV30DataListStatus = "OFFLINE_BUDGET"
	OK_PromotionListV30DataListStatus                           PromotionListV30DataListStatus = "OK"
	PREOFFLINE_BUDGET_PromotionListV30DataListStatus            PromotionListV30DataListStatus = "PREOFFLINE_BUDGET"
	PRODUCT_OFFLINE_PromotionListV30DataListStatus              PromotionListV30DataListStatus = "PRODUCT_OFFLINE"
	PROJECT_DISABLED_PromotionListV30DataListStatus             PromotionListV30DataListStatus = "PROJECT_DISABLED"
	PROJECT_OFFLINE_BUDGET_PromotionListV30DataListStatus       PromotionListV30DataListStatus = "PROJECT_OFFLINE_BUDGET"
	PROJECT_PREOFFLINE_BUDGET_PromotionListV30DataListStatus    PromotionListV30DataListStatus = "PROJECT_PREOFFLINE_BUDGET"
	REAUDIT_PromotionListV30DataListStatus                      PromotionListV30DataListStatus = "REAUDIT"
	TIME_DONE_PromotionListV30DataListStatus                    PromotionListV30DataListStatus = "TIME_DONE"
	TIME_NO_REACH_PromotionListV30DataListStatus                PromotionListV30DataListStatus = "TIME_NO_REACH"
)

// All allowed values of PromotionListV30DataListStatus enum
var AllowedPromotionListV30DataListStatusEnumValues = []PromotionListV30DataListStatus{
	"ADVERTISER_OFFLINE_BUDGET",
	"ADVERTISER_PREOFFLINE_BUDGET",
	"ALL",
	"AUDIT",
	"AUDIT_DENY",
	"AWEME_ACCOUNT_DISABLED",
	"AWEME_ANCHOR_DISABLED",
	"CREATE",
	"DELETED",
	"DISABLED",
	"DISABLE_BY_QUOTA",
	"ERROR_DEFAULT",
	"FROZEN",
	"LIVE_ROOM_OFF",
	"NOT_DELETED",
	"NO_SCHEDULE",
	"OFFLINE_BALANCE",
	"OFFLINE_BUDGET",
	"OK",
	"PREOFFLINE_BUDGET",
	"PRODUCT_OFFLINE",
	"PROJECT_DISABLED",
	"PROJECT_OFFLINE_BUDGET",
	"PROJECT_PREOFFLINE_BUDGET",
	"REAUDIT",
	"TIME_DONE",
	"TIME_NO_REACH",
}

// NewPromotionListV30DataListStatusFromValue returns a pointer to a valid PromotionListV30DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListStatusFromValue(v string) (*PromotionListV30DataListStatus, error) {
	ev := PromotionListV30DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListStatus: valid values are %v", v, AllowedPromotionListV30DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListStatus) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_status value
func (v PromotionListV30DataListStatus) Ptr() *PromotionListV30DataListStatus {
	return &v
}
