/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionUpdateV30ShopMultiRoiGoalsShopPlatform
type PromotionUpdateV30ShopMultiRoiGoalsShopPlatform string

// List of promotion_update_v3.0_shop_multi_roi_goals_shop_platform
const (
	JD_PromotionUpdateV30ShopMultiRoiGoalsShopPlatform    PromotionUpdateV30ShopMultiRoiGoalsShopPlatform = "JD"
	OTHER_PromotionUpdateV30ShopMultiRoiGoalsShopPlatform PromotionUpdateV30ShopMultiRoiGoalsShopPlatform = "OTHER"
	PDD_PromotionUpdateV30ShopMultiRoiGoalsShopPlatform   PromotionUpdateV30ShopMultiRoiGoalsShopPlatform = "PDD"
	TB_PromotionUpdateV30ShopMultiRoiGoalsShopPlatform    PromotionUpdateV30ShopMultiRoiGoalsShopPlatform = "TB"
)

// All allowed values of PromotionUpdateV30ShopMultiRoiGoalsShopPlatform enum
var AllowedPromotionUpdateV30ShopMultiRoiGoalsShopPlatformEnumValues = []PromotionUpdateV30ShopMultiRoiGoalsShopPlatform{
	"JD",
	"OTHER",
	"PDD",
	"TB",
}

// NewPromotionUpdateV30ShopMultiRoiGoalsShopPlatformFromValue returns a pointer to a valid PromotionUpdateV30ShopMultiRoiGoalsShopPlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionUpdateV30ShopMultiRoiGoalsShopPlatformFromValue(v string) (*PromotionUpdateV30ShopMultiRoiGoalsShopPlatform, error) {
	ev := PromotionUpdateV30ShopMultiRoiGoalsShopPlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionUpdateV30ShopMultiRoiGoalsShopPlatform: valid values are %v", v, AllowedPromotionUpdateV30ShopMultiRoiGoalsShopPlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionUpdateV30ShopMultiRoiGoalsShopPlatform) IsValid() bool {
	for _, existing := range AllowedPromotionUpdateV30ShopMultiRoiGoalsShopPlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_update_v3.0_shop_multi_roi_goals_shop_platform value
func (v PromotionUpdateV30ShopMultiRoiGoalsShopPlatform) Ptr() *PromotionUpdateV30ShopMultiRoiGoalsShopPlatform {
	return &v
}
