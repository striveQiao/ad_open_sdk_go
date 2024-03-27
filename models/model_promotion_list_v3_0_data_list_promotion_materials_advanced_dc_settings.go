/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionListV30DataListPromotionMaterialsAdvancedDcSettings
type PromotionListV30DataListPromotionMaterialsAdvancedDcSettings string

// List of promotion_list_v3.0_data_list_promotion_materials_advanced_dc_settings
const (
	OPTIMIZE_LANDING_PAGE_PromotionListV30DataListPromotionMaterialsAdvancedDcSettings        PromotionListV30DataListPromotionMaterialsAdvancedDcSettings = "OPTIMIZE_LANDING_PAGE"
	OPTIMIZE_SEARCH_RESULTS_PAGE_PromotionListV30DataListPromotionMaterialsAdvancedDcSettings PromotionListV30DataListPromotionMaterialsAdvancedDcSettings = "OPTIMIZE_SEARCH_RESULTS_PAGE"
)

// All allowed values of PromotionListV30DataListPromotionMaterialsAdvancedDcSettings enum
var AllowedPromotionListV30DataListPromotionMaterialsAdvancedDcSettingsEnumValues = []PromotionListV30DataListPromotionMaterialsAdvancedDcSettings{
	"OPTIMIZE_LANDING_PAGE",
	"OPTIMIZE_SEARCH_RESULTS_PAGE",
}

// NewPromotionListV30DataListPromotionMaterialsAdvancedDcSettingsFromValue returns a pointer to a valid PromotionListV30DataListPromotionMaterialsAdvancedDcSettings
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListPromotionMaterialsAdvancedDcSettingsFromValue(v string) (*PromotionListV30DataListPromotionMaterialsAdvancedDcSettings, error) {
	ev := PromotionListV30DataListPromotionMaterialsAdvancedDcSettings(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListPromotionMaterialsAdvancedDcSettings: valid values are %v", v, AllowedPromotionListV30DataListPromotionMaterialsAdvancedDcSettingsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListPromotionMaterialsAdvancedDcSettings) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListPromotionMaterialsAdvancedDcSettingsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_promotion_materials_advanced_dc_settings value
func (v PromotionListV30DataListPromotionMaterialsAdvancedDcSettings) Ptr() *PromotionListV30DataListPromotionMaterialsAdvancedDcSettings {
	return &v
}
