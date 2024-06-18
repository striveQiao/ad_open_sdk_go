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

// CampaignGetV2DataListCampaignType
type CampaignGetV2DataListCampaignType string

// List of campaign_get_v2_data_list_campaign_type
const (
	SEARCH_CampaignGetV2DataListCampaignType  CampaignGetV2DataListCampaignType = "SEARCH"
	CONTENT_CampaignGetV2DataListCampaignType CampaignGetV2DataListCampaignType = "CONTENT"
	FEED_CampaignGetV2DataListCampaignType    CampaignGetV2DataListCampaignType = "FEED"
)

// All allowed values of CampaignGetV2DataListCampaignType enum
var AllowedCampaignGetV2DataListCampaignTypeEnumValues = []CampaignGetV2DataListCampaignType{
	"SEARCH",
	"CONTENT",
	"FEED",
}

// NewCampaignGetV2DataListCampaignTypeFromValue returns a pointer to a valid CampaignGetV2DataListCampaignType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignGetV2DataListCampaignTypeFromValue(v string) (*CampaignGetV2DataListCampaignType, error) {
	ev := CampaignGetV2DataListCampaignType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignGetV2DataListCampaignType: valid values are %v", v, AllowedCampaignGetV2DataListCampaignTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignGetV2DataListCampaignType) IsValid() bool {
	for _, existing := range AllowedCampaignGetV2DataListCampaignTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign_get_v2_data_list_campaign_type value
func (v CampaignGetV2DataListCampaignType) Ptr() *CampaignGetV2DataListCampaignType {
	return &v
}
