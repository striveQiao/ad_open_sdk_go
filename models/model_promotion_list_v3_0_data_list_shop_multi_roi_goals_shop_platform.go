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

// PromotionListV30DataListShopMultiRoiGoalsShopPlatform
type PromotionListV30DataListShopMultiRoiGoalsShopPlatform string

// List of promotion_list_v3.0_data_list_shop_multi_roi_goals_shop_platform
const (
	JD_PromotionListV30DataListShopMultiRoiGoalsShopPlatform    PromotionListV30DataListShopMultiRoiGoalsShopPlatform = "JD"
	OTHER_PromotionListV30DataListShopMultiRoiGoalsShopPlatform PromotionListV30DataListShopMultiRoiGoalsShopPlatform = "OTHER"
	PDD_PromotionListV30DataListShopMultiRoiGoalsShopPlatform   PromotionListV30DataListShopMultiRoiGoalsShopPlatform = "PDD"
	TB_PromotionListV30DataListShopMultiRoiGoalsShopPlatform    PromotionListV30DataListShopMultiRoiGoalsShopPlatform = "TB"
)

// All allowed values of PromotionListV30DataListShopMultiRoiGoalsShopPlatform enum
var AllowedPromotionListV30DataListShopMultiRoiGoalsShopPlatformEnumValues = []PromotionListV30DataListShopMultiRoiGoalsShopPlatform{
	"JD",
	"OTHER",
	"PDD",
	"TB",
}

// NewPromotionListV30DataListShopMultiRoiGoalsShopPlatformFromValue returns a pointer to a valid PromotionListV30DataListShopMultiRoiGoalsShopPlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListShopMultiRoiGoalsShopPlatformFromValue(v string) (*PromotionListV30DataListShopMultiRoiGoalsShopPlatform, error) {
	ev := PromotionListV30DataListShopMultiRoiGoalsShopPlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListShopMultiRoiGoalsShopPlatform: valid values are %v", v, AllowedPromotionListV30DataListShopMultiRoiGoalsShopPlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListShopMultiRoiGoalsShopPlatform) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListShopMultiRoiGoalsShopPlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_shop_multi_roi_goals_shop_platform value
func (v PromotionListV30DataListShopMultiRoiGoalsShopPlatform) Ptr() *PromotionListV30DataListShopMultiRoiGoalsShopPlatform {
	return &v
}
