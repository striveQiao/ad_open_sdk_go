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

// CampaignCreateV2DeliveryMode
type CampaignCreateV2DeliveryMode string

// List of campaign_create_v2_delivery_mode
const (
	MANUAL_CampaignCreateV2DeliveryMode     CampaignCreateV2DeliveryMode = "MANUAL"
	PROCEDURAL_CampaignCreateV2DeliveryMode CampaignCreateV2DeliveryMode = "PROCEDURAL"
)

// All allowed values of CampaignCreateV2DeliveryMode enum
var AllowedCampaignCreateV2DeliveryModeEnumValues = []CampaignCreateV2DeliveryMode{
	"MANUAL",
	"PROCEDURAL",
}

// NewCampaignCreateV2DeliveryModeFromValue returns a pointer to a valid CampaignCreateV2DeliveryMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignCreateV2DeliveryModeFromValue(v string) (*CampaignCreateV2DeliveryMode, error) {
	ev := CampaignCreateV2DeliveryMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignCreateV2DeliveryMode: valid values are %v", v, AllowedCampaignCreateV2DeliveryModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignCreateV2DeliveryMode) IsValid() bool {
	for _, existing := range AllowedCampaignCreateV2DeliveryModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign_create_v2_delivery_mode value
func (v CampaignCreateV2DeliveryMode) Ptr() *CampaignCreateV2DeliveryMode {
	return &v
}
