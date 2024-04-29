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

// PromotionRejectReasonGetV30DataListMaterialRejectType
type PromotionRejectReasonGetV30DataListMaterialRejectType string

// List of promotion_reject_reason_get_v3.0_data_list_material_reject_type
const (
	ANCHOR_PromotionRejectReasonGetV30DataListMaterialRejectType                 PromotionRejectReasonGetV30DataListMaterialRejectType = "ANCHOR"
	AWEME_NICK_NAME_PromotionRejectReasonGetV30DataListMaterialRejectType        PromotionRejectReasonGetV30DataListMaterialRejectType = "AWEME_NICK_NAME"
	CALL_TO_ACTION_PromotionRejectReasonGetV30DataListMaterialRejectType         PromotionRejectReasonGetV30DataListMaterialRejectType = "CALL_TO_ACTION"
	CREATIVE_CAROUSEL_PromotionRejectReasonGetV30DataListMaterialRejectType      PromotionRejectReasonGetV30DataListMaterialRejectType = "CREATIVE_CAROUSEL"
	CREATIVE_COMPONENT_PromotionRejectReasonGetV30DataListMaterialRejectType     PromotionRejectReasonGetV30DataListMaterialRejectType = "CREATIVE_COMPONENT"
	CREATIVE_IMAGE_PromotionRejectReasonGetV30DataListMaterialRejectType         PromotionRejectReasonGetV30DataListMaterialRejectType = "CREATIVE_IMAGE"
	CREATIVE_TITLE_PromotionRejectReasonGetV30DataListMaterialRejectType         PromotionRejectReasonGetV30DataListMaterialRejectType = "CREATIVE_TITLE"
	CREATIVE_URL_PromotionRejectReasonGetV30DataListMaterialRejectType           PromotionRejectReasonGetV30DataListMaterialRejectType = "CREATIVE_URL"
	CREATIVE_VIDEO_PromotionRejectReasonGetV30DataListMaterialRejectType         PromotionRejectReasonGetV30DataListMaterialRejectType = "CREATIVE_VIDEO"
	OPEN_URL_PromotionRejectReasonGetV30DataListMaterialRejectType               PromotionRejectReasonGetV30DataListMaterialRejectType = "OPEN_URL"
	PRODUCT_DESCRIBE_PromotionRejectReasonGetV30DataListMaterialRejectType       PromotionRejectReasonGetV30DataListMaterialRejectType = "PRODUCT_DESCRIBE"
	PRODUCT_IMAGE_PromotionRejectReasonGetV30DataListMaterialRejectType          PromotionRejectReasonGetV30DataListMaterialRejectType = "PRODUCT_IMAGE"
	PRODUCT_SELLING_POINTS_PromotionRejectReasonGetV30DataListMaterialRejectType PromotionRejectReasonGetV30DataListMaterialRejectType = "PRODUCT_SELLING_POINTS"
)

// All allowed values of PromotionRejectReasonGetV30DataListMaterialRejectType enum
var AllowedPromotionRejectReasonGetV30DataListMaterialRejectTypeEnumValues = []PromotionRejectReasonGetV30DataListMaterialRejectType{
	"ANCHOR",
	"AWEME_NICK_NAME",
	"CALL_TO_ACTION",
	"CREATIVE_CAROUSEL",
	"CREATIVE_COMPONENT",
	"CREATIVE_IMAGE",
	"CREATIVE_TITLE",
	"CREATIVE_URL",
	"CREATIVE_VIDEO",
	"OPEN_URL",
	"PRODUCT_DESCRIBE",
	"PRODUCT_IMAGE",
	"PRODUCT_SELLING_POINTS",
}

// NewPromotionRejectReasonGetV30DataListMaterialRejectTypeFromValue returns a pointer to a valid PromotionRejectReasonGetV30DataListMaterialRejectType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionRejectReasonGetV30DataListMaterialRejectTypeFromValue(v string) (*PromotionRejectReasonGetV30DataListMaterialRejectType, error) {
	ev := PromotionRejectReasonGetV30DataListMaterialRejectType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionRejectReasonGetV30DataListMaterialRejectType: valid values are %v", v, AllowedPromotionRejectReasonGetV30DataListMaterialRejectTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionRejectReasonGetV30DataListMaterialRejectType) IsValid() bool {
	for _, existing := range AllowedPromotionRejectReasonGetV30DataListMaterialRejectTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_reject_reason_get_v3.0_data_list_material_reject_type value
func (v PromotionRejectReasonGetV30DataListMaterialRejectType) Ptr() *PromotionRejectReasonGetV30DataListMaterialRejectType {
	return &v
}
