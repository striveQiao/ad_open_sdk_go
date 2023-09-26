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

// AdGetV2FilteringStatus
type AdGetV2FilteringStatus string

// List of ad_get_v2_filtering_status
const (
	AD_STATUS_AUDIT_DENY_AdGetV2FilteringStatus       AdGetV2FilteringStatus = "AD_STATUS_AUDIT_DENY"
	AD_STATUS_CAMPAIGN_EXCEED_AdGetV2FilteringStatus  AdGetV2FilteringStatus = "AD_STATUS_CAMPAIGN_EXCEED"
	AD_STATUS_BALANCE_EXCEED_AdGetV2FilteringStatus   AdGetV2FilteringStatus = "AD_STATUS_BALANCE_EXCEED"
	AD_STATUS_DONE_AdGetV2FilteringStatus             AdGetV2FilteringStatus = "AD_STATUS_DONE"
	AD_STATUS_CAMPAIGN_DISABLE_AdGetV2FilteringStatus AdGetV2FilteringStatus = "AD_STATUS_CAMPAIGN_DISABLE"
	AD_STATUS_BUDGET_EXCEED_AdGetV2FilteringStatus    AdGetV2FilteringStatus = "AD_STATUS_BUDGET_EXCEED"
	AD_STATUS_NOT_DELETE_AdGetV2FilteringStatus       AdGetV2FilteringStatus = "AD_STATUS_NOT_DELETE"
	AD_STATUS_NOT_START_AdGetV2FilteringStatus        AdGetV2FilteringStatus = "AD_STATUS_NOT_START"
	AD_STATUS_REAUDIT_AdGetV2FilteringStatus          AdGetV2FilteringStatus = "AD_STATUS_REAUDIT"
	AD_STATUS_DISABLE_AdGetV2FilteringStatus          AdGetV2FilteringStatus = "AD_STATUS_DISABLE"
	AD_STATUS_DELIVERY_OK_AdGetV2FilteringStatus      AdGetV2FilteringStatus = "AD_STATUS_DELIVERY_OK"
	AD_STATUS_NO_SCHEDULE_AdGetV2FilteringStatus      AdGetV2FilteringStatus = "AD_STATUS_NO_SCHEDULE"
	AD_STATUS_DELETE_AdGetV2FilteringStatus           AdGetV2FilteringStatus = "AD_STATUS_DELETE"
	AD_STATUS_ALL_AdGetV2FilteringStatus              AdGetV2FilteringStatus = "AD_STATUS_ALL"
	AD_STATUS_AUDIT_AdGetV2FilteringStatus            AdGetV2FilteringStatus = "AD_STATUS_AUDIT"
	AD_STATUS_CREATE_AdGetV2FilteringStatus           AdGetV2FilteringStatus = "AD_STATUS_CREATE"
)

// All allowed values of AdGetV2FilteringStatus enum
var AllowedAdGetV2FilteringStatusEnumValues = []AdGetV2FilteringStatus{
	"AD_STATUS_AUDIT_DENY",
	"AD_STATUS_CAMPAIGN_EXCEED",
	"AD_STATUS_BALANCE_EXCEED",
	"AD_STATUS_DONE",
	"AD_STATUS_CAMPAIGN_DISABLE",
	"AD_STATUS_BUDGET_EXCEED",
	"AD_STATUS_NOT_DELETE",
	"AD_STATUS_NOT_START",
	"AD_STATUS_REAUDIT",
	"AD_STATUS_DISABLE",
	"AD_STATUS_DELIVERY_OK",
	"AD_STATUS_NO_SCHEDULE",
	"AD_STATUS_DELETE",
	"AD_STATUS_ALL",
	"AD_STATUS_AUDIT",
	"AD_STATUS_CREATE",
}

// NewAdGetV2FilteringStatusFromValue returns a pointer to a valid AdGetV2FilteringStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2FilteringStatusFromValue(v string) (*AdGetV2FilteringStatus, error) {
	ev := AdGetV2FilteringStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2FilteringStatus: valid values are %v", v, AllowedAdGetV2FilteringStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2FilteringStatus) IsValid() bool {
	for _, existing := range AllowedAdGetV2FilteringStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_filtering_status value
func (v AdGetV2FilteringStatus) Ptr() *AdGetV2FilteringStatus {
	return &v
}
