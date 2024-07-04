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

// ToolsBidsSuggestV30DeepBidType
type ToolsBidsSuggestV30DeepBidType string

// List of tools_bids_suggest_v3.0_deep_bid_type
const (
	ALI_FL_ToolsBidsSuggestV30DeepBidType                       ToolsBidsSuggestV30DeepBidType = "ALI_FL"
	AUTO_MIN_SECOND_STAGE_ToolsBidsSuggestV30DeepBidType        ToolsBidsSuggestV30DeepBidType = "AUTO_MIN_SECOND_STAGE"
	BID_PER_ACTION_ToolsBidsSuggestV30DeepBidType               ToolsBidsSuggestV30DeepBidType = "BID_PER_ACTION"
	DEEP_BID_DEFAULT_ToolsBidsSuggestV30DeepBidType             ToolsBidsSuggestV30DeepBidType = "DEEP_BID_DEFAULT"
	DEEP_BID_MIN_ToolsBidsSuggestV30DeepBidType                 ToolsBidsSuggestV30DeepBidType = "DEEP_BID_MIN"
	DEEP_BID_PACING_ToolsBidsSuggestV30DeepBidType              ToolsBidsSuggestV30DeepBidType = "DEEP_BID_PACING"
	DEEP_BID_TYPE_RETENTION_DAYS_ToolsBidsSuggestV30DeepBidType ToolsBidsSuggestV30DeepBidType = "DEEP_BID_TYPE_RETENTION_DAYS"
	FORM_BID_ToolsBidsSuggestV30DeepBidType                     ToolsBidsSuggestV30DeepBidType = "FORM_BID"
	MIN_SECOND_STAGE_ToolsBidsSuggestV30DeepBidType             ToolsBidsSuggestV30DeepBidType = "MIN_SECOND_STAGE"
	PACING_SECOND_STAGE_ToolsBidsSuggestV30DeepBidType          ToolsBidsSuggestV30DeepBidType = "PACING_SECOND_STAGE"
	PHONE_CONNECT_BID_ToolsBidsSuggestV30DeepBidType            ToolsBidsSuggestV30DeepBidType = "PHONE_CONNECT_BID"
	ROI_COEFFICIENT_ToolsBidsSuggestV30DeepBidType              ToolsBidsSuggestV30DeepBidType = "ROI_COEFFICIENT"
	ROI_DIRECT_MAIL_ToolsBidsSuggestV30DeepBidType              ToolsBidsSuggestV30DeepBidType = "ROI_DIRECT_MAIL"
	ROI_PACING_ToolsBidsSuggestV30DeepBidType                   ToolsBidsSuggestV30DeepBidType = "ROI_PACING"
	SMARTBID_ToolsBidsSuggestV30DeepBidType                     ToolsBidsSuggestV30DeepBidType = "SMARTBID"
	SOCIAL_ROI_ToolsBidsSuggestV30DeepBidType                   ToolsBidsSuggestV30DeepBidType = "SOCIAL_ROI"
)

// All allowed values of ToolsBidsSuggestV30DeepBidType enum
var AllowedToolsBidsSuggestV30DeepBidTypeEnumValues = []ToolsBidsSuggestV30DeepBidType{
	"ALI_FL",
	"AUTO_MIN_SECOND_STAGE",
	"BID_PER_ACTION",
	"DEEP_BID_DEFAULT",
	"DEEP_BID_MIN",
	"DEEP_BID_PACING",
	"DEEP_BID_TYPE_RETENTION_DAYS",
	"FORM_BID",
	"MIN_SECOND_STAGE",
	"PACING_SECOND_STAGE",
	"PHONE_CONNECT_BID",
	"ROI_COEFFICIENT",
	"ROI_DIRECT_MAIL",
	"ROI_PACING",
	"SMARTBID",
	"SOCIAL_ROI",
}

// NewToolsBidsSuggestV30DeepBidTypeFromValue returns a pointer to a valid ToolsBidsSuggestV30DeepBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidsSuggestV30DeepBidTypeFromValue(v string) (*ToolsBidsSuggestV30DeepBidType, error) {
	ev := ToolsBidsSuggestV30DeepBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidsSuggestV30DeepBidType: valid values are %v", v, AllowedToolsBidsSuggestV30DeepBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidsSuggestV30DeepBidType) IsValid() bool {
	for _, existing := range AllowedToolsBidsSuggestV30DeepBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bids_suggest_v3.0_deep_bid_type value
func (v ToolsBidsSuggestV30DeepBidType) Ptr() *ToolsBidsSuggestV30DeepBidType {
	return &v
}
