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

// AssetsCreativeComponentUpdateV2ComponentInfoComponentType
type AssetsCreativeComponentUpdateV2ComponentInfoComponentType string

// List of assets_creative_component_update_v2_component_info_component_type
const (
	CHOICE_MAGNET_AssetsCreativeComponentUpdateV2ComponentInfoComponentType           AssetsCreativeComponentUpdateV2ComponentInfoComponentType = "CHOICE_MAGNET"
	GAME_PACK_AssetsCreativeComponentUpdateV2ComponentInfoComponentType               AssetsCreativeComponentUpdateV2ComponentInfoComponentType = "GAME_PACK"
	VOTE_MAGNET_AssetsCreativeComponentUpdateV2ComponentInfoComponentType             AssetsCreativeComponentUpdateV2ComponentInfoComponentType = "VOTE_MAGNET"
	IMAGE_MAGNET_AssetsCreativeComponentUpdateV2ComponentInfoComponentType            AssetsCreativeComponentUpdateV2ComponentInfoComponentType = "IMAGE_MAGNET"
	LIGHT_INTER_ACTIVE_AssetsCreativeComponentUpdateV2ComponentInfoComponentType      AssetsCreativeComponentUpdateV2ComponentInfoComponentType = "LIGHT_INTER_ACTIVE"
	LUCKY_BOX_AssetsCreativeComponentUpdateV2ComponentInfoComponentType               AssetsCreativeComponentUpdateV2ComponentInfoComponentType = "LUCKY_BOX"
	RESERVATION_BUTTON_AssetsCreativeComponentUpdateV2ComponentInfoComponentType      AssetsCreativeComponentUpdateV2ComponentInfoComponentType = "RESERVATION_BUTTON"
	UNION_LIGHT_INTERACTIVE_AssetsCreativeComponentUpdateV2ComponentInfoComponentType AssetsCreativeComponentUpdateV2ComponentInfoComponentType = "UNION_LIGHT_INTERACTIVE"
	COMMERCE_MAGNET_AssetsCreativeComponentUpdateV2ComponentInfoComponentType         AssetsCreativeComponentUpdateV2ComponentInfoComponentType = "COMMERCE_MAGNET"
	PROMOTION_CARD_AssetsCreativeComponentUpdateV2ComponentInfoComponentType          AssetsCreativeComponentUpdateV2ComponentInfoComponentType = "PROMOTION_CARD"
	COUPON_MAGNET_AssetsCreativeComponentUpdateV2ComponentInfoComponentType           AssetsCreativeComponentUpdateV2ComponentInfoComponentType = "COUPON_MAGNET"
)

// All allowed values of AssetsCreativeComponentUpdateV2ComponentInfoComponentType enum
var AllowedAssetsCreativeComponentUpdateV2ComponentInfoComponentTypeEnumValues = []AssetsCreativeComponentUpdateV2ComponentInfoComponentType{
	"CHOICE_MAGNET",
	"GAME_PACK",
	"VOTE_MAGNET",
	"IMAGE_MAGNET",
	"LIGHT_INTER_ACTIVE",
	"LUCKY_BOX",
	"RESERVATION_BUTTON",
	"UNION_LIGHT_INTERACTIVE",
	"COMMERCE_MAGNET",
	"PROMOTION_CARD",
	"COUPON_MAGNET",
}

// NewAssetsCreativeComponentUpdateV2ComponentInfoComponentTypeFromValue returns a pointer to a valid AssetsCreativeComponentUpdateV2ComponentInfoComponentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAssetsCreativeComponentUpdateV2ComponentInfoComponentTypeFromValue(v string) (*AssetsCreativeComponentUpdateV2ComponentInfoComponentType, error) {
	ev := AssetsCreativeComponentUpdateV2ComponentInfoComponentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AssetsCreativeComponentUpdateV2ComponentInfoComponentType: valid values are %v", v, AllowedAssetsCreativeComponentUpdateV2ComponentInfoComponentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AssetsCreativeComponentUpdateV2ComponentInfoComponentType) IsValid() bool {
	for _, existing := range AllowedAssetsCreativeComponentUpdateV2ComponentInfoComponentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to assets_creative_component_update_v2_component_info_component_type value
func (v AssetsCreativeComponentUpdateV2ComponentInfoComponentType) Ptr() *AssetsCreativeComponentUpdateV2ComponentInfoComponentType {
	return &v
}
