/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CampaignGetV2FilteringDedicateType
type CampaignGetV2FilteringDedicateType string

// List of campaign_get_v2_filtering_dedicate_type
const (
	UNSET_CampaignGetV2FilteringDedicateType     CampaignGetV2FilteringDedicateType = "UNSET"
	DEDICATED_CampaignGetV2FilteringDedicateType CampaignGetV2FilteringDedicateType = "DEDICATED"
)

// All allowed values of CampaignGetV2FilteringDedicateType enum
var AllowedCampaignGetV2FilteringDedicateTypeEnumValues = []CampaignGetV2FilteringDedicateType{
	"UNSET",
	"DEDICATED",
}

// NewCampaignGetV2FilteringDedicateTypeFromValue returns a pointer to a valid CampaignGetV2FilteringDedicateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignGetV2FilteringDedicateTypeFromValue(v string) (*CampaignGetV2FilteringDedicateType, error) {
	ev := CampaignGetV2FilteringDedicateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignGetV2FilteringDedicateType: valid values are %v", v, AllowedCampaignGetV2FilteringDedicateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignGetV2FilteringDedicateType) IsValid() bool {
	for _, existing := range AllowedCampaignGetV2FilteringDedicateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign_get_v2_filtering_dedicate_type value
func (v CampaignGetV2FilteringDedicateType) Ptr() *CampaignGetV2FilteringDedicateType {
	return &v
}
