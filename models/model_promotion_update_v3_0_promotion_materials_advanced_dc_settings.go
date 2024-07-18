/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionUpdateV30PromotionMaterialsAdvancedDcSettings
type PromotionUpdateV30PromotionMaterialsAdvancedDcSettings string

// List of promotion_update_v3.0_promotion_materials_advanced_dc_settings
const (
	OPTIMIZE_LANDING_PAGE_PromotionUpdateV30PromotionMaterialsAdvancedDcSettings        PromotionUpdateV30PromotionMaterialsAdvancedDcSettings = "OPTIMIZE_LANDING_PAGE"
	OPTIMIZE_SEARCH_RESULTS_PAGE_PromotionUpdateV30PromotionMaterialsAdvancedDcSettings PromotionUpdateV30PromotionMaterialsAdvancedDcSettings = "OPTIMIZE_SEARCH_RESULTS_PAGE"
)

// All allowed values of PromotionUpdateV30PromotionMaterialsAdvancedDcSettings enum
var AllowedPromotionUpdateV30PromotionMaterialsAdvancedDcSettingsEnumValues = []PromotionUpdateV30PromotionMaterialsAdvancedDcSettings{
	"OPTIMIZE_LANDING_PAGE",
	"OPTIMIZE_SEARCH_RESULTS_PAGE",
}

// NewPromotionUpdateV30PromotionMaterialsAdvancedDcSettingsFromValue returns a pointer to a valid PromotionUpdateV30PromotionMaterialsAdvancedDcSettings
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionUpdateV30PromotionMaterialsAdvancedDcSettingsFromValue(v string) (*PromotionUpdateV30PromotionMaterialsAdvancedDcSettings, error) {
	ev := PromotionUpdateV30PromotionMaterialsAdvancedDcSettings(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionUpdateV30PromotionMaterialsAdvancedDcSettings: valid values are %v", v, AllowedPromotionUpdateV30PromotionMaterialsAdvancedDcSettingsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionUpdateV30PromotionMaterialsAdvancedDcSettings) IsValid() bool {
	for _, existing := range AllowedPromotionUpdateV30PromotionMaterialsAdvancedDcSettingsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_update_v3.0_promotion_materials_advanced_dc_settings value
func (v PromotionUpdateV30PromotionMaterialsAdvancedDcSettings) Ptr() *PromotionUpdateV30PromotionMaterialsAdvancedDcSettings {
	return &v
}
