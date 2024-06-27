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

// CampaignCreateV2Operation
type CampaignCreateV2Operation string

// List of campaign_create_v2_operation
const (
	ENABLE_CampaignCreateV2Operation  CampaignCreateV2Operation = "enable"
	DISABLE_CampaignCreateV2Operation CampaignCreateV2Operation = "disable"
)

// All allowed values of CampaignCreateV2Operation enum
var AllowedCampaignCreateV2OperationEnumValues = []CampaignCreateV2Operation{
	"enable",
	"disable",
}

// NewCampaignCreateV2OperationFromValue returns a pointer to a valid CampaignCreateV2Operation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignCreateV2OperationFromValue(v string) (*CampaignCreateV2Operation, error) {
	ev := CampaignCreateV2Operation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignCreateV2Operation: valid values are %v", v, AllowedCampaignCreateV2OperationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignCreateV2Operation) IsValid() bool {
	for _, existing := range AllowedCampaignCreateV2OperationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign_create_v2_operation value
func (v CampaignCreateV2Operation) Ptr() *CampaignCreateV2Operation {
	return &v
}
