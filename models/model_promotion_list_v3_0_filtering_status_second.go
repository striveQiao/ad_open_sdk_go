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

// PromotionListV30FilteringStatusSecond
type PromotionListV30FilteringStatusSecond string

// List of promotion_list_v3.0_filtering_status_second
const (
	AUDIT_PromotionListV30FilteringStatusSecond                       PromotionListV30FilteringStatusSecond = "AUDIT"
	AUDIT_DENY_PromotionListV30FilteringStatusSecond                  PromotionListV30FilteringStatusSecond = "AUDIT_DENY"
	AWEME_ACCOUNT_DISABLED_PromotionListV30FilteringStatusSecond      PromotionListV30FilteringStatusSecond = "AWEME_ACCOUNT_DISABLED"
	AWEME_ANCHOR_DISABLED_PromotionListV30FilteringStatusSecond       PromotionListV30FilteringStatusSecond = "AWEME_ANCHOR_DISABLED"
	BALANCE_OFFLINE_BUDGET_PromotionListV30FilteringStatusSecond      PromotionListV30FilteringStatusSecond = "BALANCE_OFFLINE_BUDGET"
	BUDGET_GROUP_OFFLINE_BUDGET_PromotionListV30FilteringStatusSecond PromotionListV30FilteringStatusSecond = "BUDGET_GROUP_OFFLINE_BUDGET"
	DISABLED_PromotionListV30FilteringStatusSecond                    PromotionListV30FilteringStatusSecond = "DISABLED"
	DISABLE_BY_QUOTA_PromotionListV30FilteringStatusSecond            PromotionListV30FilteringStatusSecond = "DISABLE_BY_QUOTA"
	LIVE_ROOM_OFF_PromotionListV30FilteringStatusSecond               PromotionListV30FilteringStatusSecond = "LIVE_ROOM_OFF"
	NO_SCHEDULE_PromotionListV30FilteringStatusSecond                 PromotionListV30FilteringStatusSecond = "NO_SCHEDULE"
	OFFLINE_BALANCE_PromotionListV30FilteringStatusSecond             PromotionListV30FilteringStatusSecond = "OFFLINE_BALANCE"
	PRODUCT_OFFLINE_PromotionListV30FilteringStatusSecond             PromotionListV30FilteringStatusSecond = "PRODUCT_OFFLINE"
	PROJECT_DISABLED_PromotionListV30FilteringStatusSecond            PromotionListV30FilteringStatusSecond = "PROJECT_DISABLED"
	PROJECT_OFFLINE_BUDGET_PromotionListV30FilteringStatusSecond      PromotionListV30FilteringStatusSecond = "PROJECT_OFFLINE_BUDGET"
	PROMOTION_OFFLINE_BUDGET_PromotionListV30FilteringStatusSecond    PromotionListV30FilteringStatusSecond = "PROMOTION_OFFLINE_BUDGET"
	REAUDIT_PromotionListV30FilteringStatusSecond                     PromotionListV30FilteringStatusSecond = "REAUDIT"
	TIME_NO_REACH_PromotionListV30FilteringStatusSecond               PromotionListV30FilteringStatusSecond = "TIME_NO_REACH"
)

// All allowed values of PromotionListV30FilteringStatusSecond enum
var AllowedPromotionListV30FilteringStatusSecondEnumValues = []PromotionListV30FilteringStatusSecond{
	"AUDIT",
	"AUDIT_DENY",
	"AWEME_ACCOUNT_DISABLED",
	"AWEME_ANCHOR_DISABLED",
	"BALANCE_OFFLINE_BUDGET",
	"BUDGET_GROUP_OFFLINE_BUDGET",
	"DISABLED",
	"DISABLE_BY_QUOTA",
	"LIVE_ROOM_OFF",
	"NO_SCHEDULE",
	"OFFLINE_BALANCE",
	"PRODUCT_OFFLINE",
	"PROJECT_DISABLED",
	"PROJECT_OFFLINE_BUDGET",
	"PROMOTION_OFFLINE_BUDGET",
	"REAUDIT",
	"TIME_NO_REACH",
}

// NewPromotionListV30FilteringStatusSecondFromValue returns a pointer to a valid PromotionListV30FilteringStatusSecond
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30FilteringStatusSecondFromValue(v string) (*PromotionListV30FilteringStatusSecond, error) {
	ev := PromotionListV30FilteringStatusSecond(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30FilteringStatusSecond: valid values are %v", v, AllowedPromotionListV30FilteringStatusSecondEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30FilteringStatusSecond) IsValid() bool {
	for _, existing := range AllowedPromotionListV30FilteringStatusSecondEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_filtering_status_second value
func (v PromotionListV30FilteringStatusSecond) Ptr() *PromotionListV30FilteringStatusSecond {
	return &v
}
