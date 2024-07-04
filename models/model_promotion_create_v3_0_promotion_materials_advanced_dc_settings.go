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

// PromotionCreateV30PromotionMaterialsAdvancedDcSettings
type PromotionCreateV30PromotionMaterialsAdvancedDcSettings string

// List of promotion_create_v3.0_promotion_materials_advanced_dc_settings
const (
	OPTIMIZE_LANDING_PAGE_PromotionCreateV30PromotionMaterialsAdvancedDcSettings        PromotionCreateV30PromotionMaterialsAdvancedDcSettings = "OPTIMIZE_LANDING_PAGE"
	OPTIMIZE_SEARCH_RESULTS_PAGE_PromotionCreateV30PromotionMaterialsAdvancedDcSettings PromotionCreateV30PromotionMaterialsAdvancedDcSettings = "OPTIMIZE_SEARCH_RESULTS_PAGE"
)

// All allowed values of PromotionCreateV30PromotionMaterialsAdvancedDcSettings enum
var AllowedPromotionCreateV30PromotionMaterialsAdvancedDcSettingsEnumValues = []PromotionCreateV30PromotionMaterialsAdvancedDcSettings{
	"OPTIMIZE_LANDING_PAGE",
	"OPTIMIZE_SEARCH_RESULTS_PAGE",
}

// NewPromotionCreateV30PromotionMaterialsAdvancedDcSettingsFromValue returns a pointer to a valid PromotionCreateV30PromotionMaterialsAdvancedDcSettings
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30PromotionMaterialsAdvancedDcSettingsFromValue(v string) (*PromotionCreateV30PromotionMaterialsAdvancedDcSettings, error) {
	ev := PromotionCreateV30PromotionMaterialsAdvancedDcSettings(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30PromotionMaterialsAdvancedDcSettings: valid values are %v", v, AllowedPromotionCreateV30PromotionMaterialsAdvancedDcSettingsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30PromotionMaterialsAdvancedDcSettings) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30PromotionMaterialsAdvancedDcSettingsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_promotion_materials_advanced_dc_settings value
func (v PromotionCreateV30PromotionMaterialsAdvancedDcSettings) Ptr() *PromotionCreateV30PromotionMaterialsAdvancedDcSettings {
	return &v
}
