/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CampaignCreateV2DeliveryRelatedNum
type CampaignCreateV2DeliveryRelatedNum string

// List of campaign_create_v2_delivery_related_num
const (
	CAMPAIGN_DPA_SINGLE_DELIVERY_CampaignCreateV2DeliveryRelatedNum CampaignCreateV2DeliveryRelatedNum = "CAMPAIGN_DPA_SINGLE_DELIVERY"
	CAMPAIGN_DPA_MULTI_DELIVERY_CampaignCreateV2DeliveryRelatedNum  CampaignCreateV2DeliveryRelatedNum = "CAMPAIGN_DPA_MULTI_DELIVERY"
	CAMPAIGN_DPA_DEFAULT_NOT_CampaignCreateV2DeliveryRelatedNum     CampaignCreateV2DeliveryRelatedNum = "CAMPAIGN_DPA_DEFAULT_NOT"
)

// All allowed values of CampaignCreateV2DeliveryRelatedNum enum
var AllowedCampaignCreateV2DeliveryRelatedNumEnumValues = []CampaignCreateV2DeliveryRelatedNum{
	"CAMPAIGN_DPA_SINGLE_DELIVERY",
	"CAMPAIGN_DPA_MULTI_DELIVERY",
	"CAMPAIGN_DPA_DEFAULT_NOT",
}

// NewCampaignCreateV2DeliveryRelatedNumFromValue returns a pointer to a valid CampaignCreateV2DeliveryRelatedNum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignCreateV2DeliveryRelatedNumFromValue(v string) (*CampaignCreateV2DeliveryRelatedNum, error) {
	ev := CampaignCreateV2DeliveryRelatedNum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignCreateV2DeliveryRelatedNum: valid values are %v", v, AllowedCampaignCreateV2DeliveryRelatedNumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignCreateV2DeliveryRelatedNum) IsValid() bool {
	for _, existing := range AllowedCampaignCreateV2DeliveryRelatedNumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign_create_v2_delivery_related_num value
func (v CampaignCreateV2DeliveryRelatedNum) Ptr() *CampaignCreateV2DeliveryRelatedNum {
	return &v
}
