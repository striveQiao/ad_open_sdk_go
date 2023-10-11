/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAdGetV2Filtering
type ReportAdGetV2Filtering struct {
	//
	AdIds []int64 `json:"ad_ids,omitempty"`
	//
	AdName *string `json:"ad_name,omitempty"`
	//
	CampaignId *int64 `json:"campaign_id,omitempty"`
	//
	CampaignIds []int64 `json:"campaign_ids,omitempty"`
	//
	CampaignTypes []*ReportAdGetV2FilteringCampaignTypes `json:"campaign_types,omitempty"`
	//
	CreativeIds []int64 `json:"creative_ids,omitempty"`
	//
	CreativeMaterialModes []*ReportAdGetV2FilteringCreativeMaterialModes `json:"creative_material_modes,omitempty"`
	//
	DeliveryMode []*ReportAdGetV2FilteringDeliveryMode `json:"delivery_mode,omitempty"`
	//
	ImageModes []*ReportAdGetV2FilteringImageModes `json:"image_modes,omitempty"`
	//
	InventoryTypes []*ReportAdGetV2FilteringInventoryTypes `json:"inventory_types,omitempty"`
	LandingType    *ReportAdGetV2FilteringLandingType      `json:"landing_type,omitempty"`
	//
	LandingTypes []*ReportAdGetV2FilteringLandingTypes `json:"landing_types,omitempty"`
	Pricing      *ReportAdGetV2FilteringPricing        `json:"pricing,omitempty"`
	//
	PricingCategories []*ReportAdGetV2FilteringPricingCategories `json:"pricing_categories,omitempty"`
	//
	Pricings []*ReportAdGetV2FilteringPricings `json:"pricings,omitempty"`
	Status   *ReportAdGetV2FilteringStatus     `json:"status,omitempty"`
}
