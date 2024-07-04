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

// CampaignCreateV2CampaignType
type CampaignCreateV2CampaignType string

// List of campaign_create_v2_campaign_type
const (
	FEED_CampaignCreateV2CampaignType    CampaignCreateV2CampaignType = "FEED"
	CONTENT_CampaignCreateV2CampaignType CampaignCreateV2CampaignType = "CONTENT"
	SEARCH_CampaignCreateV2CampaignType  CampaignCreateV2CampaignType = "SEARCH"
)

// All allowed values of CampaignCreateV2CampaignType enum
var AllowedCampaignCreateV2CampaignTypeEnumValues = []CampaignCreateV2CampaignType{
	"FEED",
	"CONTENT",
	"SEARCH",
}

// NewCampaignCreateV2CampaignTypeFromValue returns a pointer to a valid CampaignCreateV2CampaignType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignCreateV2CampaignTypeFromValue(v string) (*CampaignCreateV2CampaignType, error) {
	ev := CampaignCreateV2CampaignType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignCreateV2CampaignType: valid values are %v", v, AllowedCampaignCreateV2CampaignTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignCreateV2CampaignType) IsValid() bool {
	for _, existing := range AllowedCampaignCreateV2CampaignTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign_create_v2_campaign_type value
func (v CampaignCreateV2CampaignType) Ptr() *CampaignCreateV2CampaignType {
	return &v
}
