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

// BrandCampaignGetV30DataCampaignsContractType
type BrandCampaignGetV30DataCampaignsContractType int64

// List of brand_campaign_get_v3.0_data_campaigns_contract_type
const (
	Enum_1_BrandCampaignGetV30DataCampaignsContractType BrandCampaignGetV30DataCampaignsContractType = 1
	Enum_2_BrandCampaignGetV30DataCampaignsContractType BrandCampaignGetV30DataCampaignsContractType = 2
)

// All allowed values of BrandCampaignGetV30DataCampaignsContractType enum
var AllowedBrandCampaignGetV30DataCampaignsContractTypeEnumValues = []BrandCampaignGetV30DataCampaignsContractType{
	1,
	2,
}

// NewBrandCampaignGetV30DataCampaignsContractTypeFromValue returns a pointer to a valid BrandCampaignGetV30DataCampaignsContractType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandCampaignGetV30DataCampaignsContractTypeFromValue(v int64) (*BrandCampaignGetV30DataCampaignsContractType, error) {
	ev := BrandCampaignGetV30DataCampaignsContractType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandCampaignGetV30DataCampaignsContractType: valid values are %v", v, AllowedBrandCampaignGetV30DataCampaignsContractTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandCampaignGetV30DataCampaignsContractType) IsValid() bool {
	for _, existing := range AllowedBrandCampaignGetV30DataCampaignsContractTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_campaign_get_v3.0_data_campaigns_contract_type value
func (v BrandCampaignGetV30DataCampaignsContractType) Ptr() *BrandCampaignGetV30DataCampaignsContractType {
	return &v
}
