/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionUpdateV30NativeSettingAnchorRelatedType
type PromotionUpdateV30NativeSettingAnchorRelatedType string

// List of promotion_update_v3.0_native_setting_anchor_related_type
const (
	AUTO_PromotionUpdateV30NativeSettingAnchorRelatedType   PromotionUpdateV30NativeSettingAnchorRelatedType = "AUTO"
	OFF_PromotionUpdateV30NativeSettingAnchorRelatedType    PromotionUpdateV30NativeSettingAnchorRelatedType = "OFF"
	SELECT_PromotionUpdateV30NativeSettingAnchorRelatedType PromotionUpdateV30NativeSettingAnchorRelatedType = "SELECT"
)

// All allowed values of PromotionUpdateV30NativeSettingAnchorRelatedType enum
var AllowedPromotionUpdateV30NativeSettingAnchorRelatedTypeEnumValues = []PromotionUpdateV30NativeSettingAnchorRelatedType{
	"AUTO",
	"OFF",
	"SELECT",
}

// NewPromotionUpdateV30NativeSettingAnchorRelatedTypeFromValue returns a pointer to a valid PromotionUpdateV30NativeSettingAnchorRelatedType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionUpdateV30NativeSettingAnchorRelatedTypeFromValue(v string) (*PromotionUpdateV30NativeSettingAnchorRelatedType, error) {
	ev := PromotionUpdateV30NativeSettingAnchorRelatedType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionUpdateV30NativeSettingAnchorRelatedType: valid values are %v", v, AllowedPromotionUpdateV30NativeSettingAnchorRelatedTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionUpdateV30NativeSettingAnchorRelatedType) IsValid() bool {
	for _, existing := range AllowedPromotionUpdateV30NativeSettingAnchorRelatedTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_update_v3.0_native_setting_anchor_related_type value
func (v PromotionUpdateV30NativeSettingAnchorRelatedType) Ptr() *PromotionUpdateV30NativeSettingAnchorRelatedType {
	return &v
}
