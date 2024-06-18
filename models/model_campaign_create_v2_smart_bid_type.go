/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CampaignCreateV2SmartBidType
type CampaignCreateV2SmartBidType string

// List of campaign_create_v2_smart_bid_type
const (
	BAOSHOU_CampaignCreateV2SmartBidType           CampaignCreateV2SmartBidType = "BAOSHOU"
	GD_PROMOTE_CampaignCreateV2SmartBidType        CampaignCreateV2SmartBidType = "GD_PROMOTE"
	GUARANTEED_SHOW_CampaignCreateV2SmartBidType   CampaignCreateV2SmartBidType = "GUARANTEED_SHOW"
	AWEME_LITE_PACING_CampaignCreateV2SmartBidType CampaignCreateV2SmartBidType = "AWEME_LITE_PACING"
	MAXCV_CampaignCreateV2SmartBidType             CampaignCreateV2SmartBidType = "MAXCV"
	CUSTOM_CampaignCreateV2SmartBidType            CampaignCreateV2SmartBidType = "CUSTOM"
	JIJIN_CampaignCreateV2SmartBidType             CampaignCreateV2SmartBidType = "JIJIN"
	SMART_BID_NO_BID_CampaignCreateV2SmartBidType  CampaignCreateV2SmartBidType = "SMART_BID_NO_BID"
	LITE_PACING_CampaignCreateV2SmartBidType       CampaignCreateV2SmartBidType = "LITE_PACING"
	MAX_CONVERSION_CampaignCreateV2SmartBidType    CampaignCreateV2SmartBidType = "MAX_CONVERSION"
)

// All allowed values of CampaignCreateV2SmartBidType enum
var AllowedCampaignCreateV2SmartBidTypeEnumValues = []CampaignCreateV2SmartBidType{
	"BAOSHOU",
	"GD_PROMOTE",
	"GUARANTEED_SHOW",
	"AWEME_LITE_PACING",
	"MAXCV",
	"CUSTOM",
	"JIJIN",
	"SMART_BID_NO_BID",
	"LITE_PACING",
	"MAX_CONVERSION",
}

// NewCampaignCreateV2SmartBidTypeFromValue returns a pointer to a valid CampaignCreateV2SmartBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignCreateV2SmartBidTypeFromValue(v string) (*CampaignCreateV2SmartBidType, error) {
	ev := CampaignCreateV2SmartBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignCreateV2SmartBidType: valid values are %v", v, AllowedCampaignCreateV2SmartBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignCreateV2SmartBidType) IsValid() bool {
	for _, existing := range AllowedCampaignCreateV2SmartBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign_create_v2_smart_bid_type value
func (v CampaignCreateV2SmartBidType) Ptr() *CampaignCreateV2SmartBidType {
	return &v
}
