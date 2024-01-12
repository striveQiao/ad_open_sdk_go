/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CampaignGetV2DataListLandingType
type CampaignGetV2DataListLandingType string

// List of campaign_get_v2_data_list_landing_type
const (
	QUICK_APP_CampaignGetV2DataListLandingType      CampaignGetV2DataListLandingType = "QUICK_APP"
	ARTICLE_CampaignGetV2DataListLandingType        CampaignGetV2DataListLandingType = "ARTICLE"
	SHOP_CampaignGetV2DataListLandingType           CampaignGetV2DataListLandingType = "SHOP"
	GOODS_CampaignGetV2DataListLandingType          CampaignGetV2DataListLandingType = "GOODS"
	LINK_CampaignGetV2DataListLandingType           CampaignGetV2DataListLandingType = "LINK"
	AWEME_CampaignGetV2DataListLandingType          CampaignGetV2DataListLandingType = "AWEME"
	LIVE_CampaignGetV2DataListLandingType           CampaignGetV2DataListLandingType = "LIVE"
	BRAND_EXTERNAL_CampaignGetV2DataListLandingType CampaignGetV2DataListLandingType = "BRAND_EXTERNAL"
	STORE_CampaignGetV2DataListLandingType          CampaignGetV2DataListLandingType = "STORE"
	APP_CampaignGetV2DataListLandingType            CampaignGetV2DataListLandingType = "APP"
	DPA_CampaignGetV2DataListLandingType            CampaignGetV2DataListLandingType = "DPA"
)

// All allowed values of CampaignGetV2DataListLandingType enum
var AllowedCampaignGetV2DataListLandingTypeEnumValues = []CampaignGetV2DataListLandingType{
	"QUICK_APP",
	"ARTICLE",
	"SHOP",
	"GOODS",
	"LINK",
	"AWEME",
	"LIVE",
	"BRAND_EXTERNAL",
	"STORE",
	"APP",
	"DPA",
}

// NewCampaignGetV2DataListLandingTypeFromValue returns a pointer to a valid CampaignGetV2DataListLandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignGetV2DataListLandingTypeFromValue(v string) (*CampaignGetV2DataListLandingType, error) {
	ev := CampaignGetV2DataListLandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignGetV2DataListLandingType: valid values are %v", v, AllowedCampaignGetV2DataListLandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignGetV2DataListLandingType) IsValid() bool {
	for _, existing := range AllowedCampaignGetV2DataListLandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign_get_v2_data_list_landing_type value
func (v CampaignGetV2DataListLandingType) Ptr() *CampaignGetV2DataListLandingType {
	return &v
}
