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

// CampaignGetV2DataListSmartBidType
type CampaignGetV2DataListSmartBidType string

// List of campaign_get_v2_data_list_smart_bid_type
const (
	JIJIN_CampaignGetV2DataListSmartBidType             CampaignGetV2DataListSmartBidType = "JIJIN"
	SMART_BID_NO_BID_CampaignGetV2DataListSmartBidType  CampaignGetV2DataListSmartBidType = "SMART_BID_NO_BID"
	BAOSHOU_CampaignGetV2DataListSmartBidType           CampaignGetV2DataListSmartBidType = "BAOSHOU"
	GUARANTEED_SHOW_CampaignGetV2DataListSmartBidType   CampaignGetV2DataListSmartBidType = "GUARANTEED_SHOW"
	CUSTOM_CampaignGetV2DataListSmartBidType            CampaignGetV2DataListSmartBidType = "CUSTOM"
	AWEME_LITE_PACING_CampaignGetV2DataListSmartBidType CampaignGetV2DataListSmartBidType = "AWEME_LITE_PACING"
	MAX_CONVERSION_CampaignGetV2DataListSmartBidType    CampaignGetV2DataListSmartBidType = "MAX_CONVERSION"
	LITE_PACING_CampaignGetV2DataListSmartBidType       CampaignGetV2DataListSmartBidType = "LITE_PACING"
	MAXCV_CampaignGetV2DataListSmartBidType             CampaignGetV2DataListSmartBidType = "MAXCV"
	GD_PROMOTE_CampaignGetV2DataListSmartBidType        CampaignGetV2DataListSmartBidType = "GD_PROMOTE"
)

// All allowed values of CampaignGetV2DataListSmartBidType enum
var AllowedCampaignGetV2DataListSmartBidTypeEnumValues = []CampaignGetV2DataListSmartBidType{
	"JIJIN",
	"SMART_BID_NO_BID",
	"BAOSHOU",
	"GUARANTEED_SHOW",
	"CUSTOM",
	"AWEME_LITE_PACING",
	"MAX_CONVERSION",
	"LITE_PACING",
	"MAXCV",
	"GD_PROMOTE",
}

// NewCampaignGetV2DataListSmartBidTypeFromValue returns a pointer to a valid CampaignGetV2DataListSmartBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignGetV2DataListSmartBidTypeFromValue(v string) (*CampaignGetV2DataListSmartBidType, error) {
	ev := CampaignGetV2DataListSmartBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignGetV2DataListSmartBidType: valid values are %v", v, AllowedCampaignGetV2DataListSmartBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignGetV2DataListSmartBidType) IsValid() bool {
	for _, existing := range AllowedCampaignGetV2DataListSmartBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign_get_v2_data_list_smart_bid_type value
func (v CampaignGetV2DataListSmartBidType) Ptr() *CampaignGetV2DataListSmartBidType {
	return &v
}
