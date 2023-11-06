/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdDetailGetV10ResponseDataProgrammaticCreativeCard
type QianchuanAdDetailGetV10ResponseDataProgrammaticCreativeCard struct {
	//
	ComponentId *int64 `json:"component_id,omitempty"`
	//
	PromotionCardActionButton            *string                                                                                  `json:"promotion_card_action_button,omitempty"`
	PromotionCardButtonSmartOptimization *QianchuanAdDetailGetV10DataProgrammaticCreativeCardPromotionCardButtonSmartOptimization `json:"promotion_card_button_smart_optimization,omitempty"`
	//
	PromotionCardId *int64 `json:"promotion_card_id,omitempty"`
	//
	PromotionCardImageId *string `json:"promotion_card_image_id,omitempty"`
	//
	PromotionCardSellingPoints []string `json:"promotion_card_selling_points,omitempty"`
	//
	PromotionCardTitle *string `json:"promotion_card_title,omitempty"`
}
