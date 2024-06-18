/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCampaignGetV2Filtering
type ReportCampaignGetV2Filtering struct {
	//
	AdIds []int64 `json:"ad_ids,omitempty"`
	//
	CampaignIds []int64 `json:"campaign_ids,omitempty"`
	//
	CampaignName *string `json:"campaign_name,omitempty"`
	//
	CampaignTypes []*ReportCampaignGetV2FilteringCampaignTypes `json:"campaign_types,omitempty"`
	//
	CreativeIds []int64 `json:"creative_ids,omitempty"`
	//
	CreativeMaterialModes []*ReportCampaignGetV2FilteringCreativeMaterialModes `json:"creative_material_modes,omitempty"`
	//
	DeliveryMode []*ReportCampaignGetV2FilteringDeliveryMode `json:"delivery_mode,omitempty"`
	//
	ImageModes []*ReportCampaignGetV2FilteringImageModes `json:"image_modes,omitempty"`
	//
	InventoryTypes []*ReportCampaignGetV2FilteringInventoryTypes `json:"inventory_types,omitempty"`
	LandingType    *ReportCampaignGetV2FilteringLandingType      `json:"landing_type,omitempty"`
	//
	LandingTypes []*ReportCampaignGetV2FilteringLandingTypes `json:"landing_types,omitempty"`
	//
	PricingCategories []*ReportCampaignGetV2FilteringPricingCategories `json:"pricing_categories,omitempty"`
	//
	Pricings []*ReportCampaignGetV2FilteringPricings `json:"pricings,omitempty"`
	Status   *ReportCampaignGetV2FilteringStatus     `json:"status,omitempty"`
}
