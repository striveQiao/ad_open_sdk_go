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

// CampaignGetV2FilteringStatus
type CampaignGetV2FilteringStatus string

// List of campaign_get_v2_filtering_status
const (
	CAMPAIGN_STATUS_ALL_CampaignGetV2FilteringStatus        CampaignGetV2FilteringStatus = "CAMPAIGN_STATUS_ALL"
	CAMPAIGN_STATUS_ENABLE_CampaignGetV2FilteringStatus     CampaignGetV2FilteringStatus = "CAMPAIGN_STATUS_ENABLE"
	CAMPAIGN_STATUS_DISABLE_CampaignGetV2FilteringStatus    CampaignGetV2FilteringStatus = "CAMPAIGN_STATUS_DISABLE"
	CAMPAIGN_STATUS_NOT_DELETE_CampaignGetV2FilteringStatus CampaignGetV2FilteringStatus = "CAMPAIGN_STATUS_NOT_DELETE"
	CAMPAIGN_STATUS_DELETE_CampaignGetV2FilteringStatus     CampaignGetV2FilteringStatus = "CAMPAIGN_STATUS_DELETE"
)

// All allowed values of CampaignGetV2FilteringStatus enum
var AllowedCampaignGetV2FilteringStatusEnumValues = []CampaignGetV2FilteringStatus{
	"CAMPAIGN_STATUS_ALL",
	"CAMPAIGN_STATUS_ENABLE",
	"CAMPAIGN_STATUS_DISABLE",
	"CAMPAIGN_STATUS_NOT_DELETE",
	"CAMPAIGN_STATUS_DELETE",
}

// NewCampaignGetV2FilteringStatusFromValue returns a pointer to a valid CampaignGetV2FilteringStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignGetV2FilteringStatusFromValue(v string) (*CampaignGetV2FilteringStatus, error) {
	ev := CampaignGetV2FilteringStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignGetV2FilteringStatus: valid values are %v", v, AllowedCampaignGetV2FilteringStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignGetV2FilteringStatus) IsValid() bool {
	for _, existing := range AllowedCampaignGetV2FilteringStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign_get_v2_filtering_status value
func (v CampaignGetV2FilteringStatus) Ptr() *CampaignGetV2FilteringStatus {
	return &v
}
