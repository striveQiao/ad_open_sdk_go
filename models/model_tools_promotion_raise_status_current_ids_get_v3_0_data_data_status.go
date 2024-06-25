/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatus
type ToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatus string

// List of tools_promotion_raise_status_current_ids_get_v3.0_data_data_status
const (
	DISABLE_RAISE_ToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatus   ToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatus = "DISABLE_RAISE"
	ENABLE_PRERAISE_ToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatus ToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatus = "ENABLE_PRERAISE"
	ENABLE_RAISE_ToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatus    ToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatus = "ENABLE_RAISE"
	FINISH_RAISE_ToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatus    ToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatus = "FINISH_RAISE"
	HAS_PRERAISE_ToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatus    ToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatus = "HAS_PRERAISE"
	RAISING_ToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatus         ToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatus = "RAISING"
)

// All allowed values of ToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatus enum
var AllowedToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatusEnumValues = []ToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatus{
	"DISABLE_RAISE",
	"ENABLE_PRERAISE",
	"ENABLE_RAISE",
	"FINISH_RAISE",
	"HAS_PRERAISE",
	"RAISING",
}

// NewToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatusFromValue returns a pointer to a valid ToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatusFromValue(v string) (*ToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatus, error) {
	ev := ToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatus: valid values are %v", v, AllowedToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatus) IsValid() bool {
	for _, existing := range AllowedToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_promotion_raise_status_current_ids_get_v3.0_data_data_status value
func (v ToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatus) Ptr() *ToolsPromotionRaiseStatusCurrentIdsGetV30DataDataStatus {
	return &v
}
