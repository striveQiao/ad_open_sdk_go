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

// NativeAnchorGetDetailV30DataListGameAnchorGamePackageListGiftGiftUnit
type NativeAnchorGetDetailV30DataListGameAnchorGamePackageListGiftGiftUnit string

// List of native_anchor_get_detail_v3.0_data_list_game_anchor_game_package_list_gift_gift_unit
const (
	INDIVIDUAL_NativeAnchorGetDetailV30DataListGameAnchorGamePackageListGiftGiftUnit NativeAnchorGetDetailV30DataListGameAnchorGamePackageListGiftGiftUnit = "INDIVIDUAL"
	MYRIAD_NativeAnchorGetDetailV30DataListGameAnchorGamePackageListGiftGiftUnit     NativeAnchorGetDetailV30DataListGameAnchorGamePackageListGiftGiftUnit = "MYRIAD"
)

// All allowed values of NativeAnchorGetDetailV30DataListGameAnchorGamePackageListGiftGiftUnit enum
var AllowedNativeAnchorGetDetailV30DataListGameAnchorGamePackageListGiftGiftUnitEnumValues = []NativeAnchorGetDetailV30DataListGameAnchorGamePackageListGiftGiftUnit{
	"INDIVIDUAL",
	"MYRIAD",
}

// NewNativeAnchorGetDetailV30DataListGameAnchorGamePackageListGiftGiftUnitFromValue returns a pointer to a valid NativeAnchorGetDetailV30DataListGameAnchorGamePackageListGiftGiftUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorGetDetailV30DataListGameAnchorGamePackageListGiftGiftUnitFromValue(v string) (*NativeAnchorGetDetailV30DataListGameAnchorGamePackageListGiftGiftUnit, error) {
	ev := NativeAnchorGetDetailV30DataListGameAnchorGamePackageListGiftGiftUnit(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorGetDetailV30DataListGameAnchorGamePackageListGiftGiftUnit: valid values are %v", v, AllowedNativeAnchorGetDetailV30DataListGameAnchorGamePackageListGiftGiftUnitEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorGetDetailV30DataListGameAnchorGamePackageListGiftGiftUnit) IsValid() bool {
	for _, existing := range AllowedNativeAnchorGetDetailV30DataListGameAnchorGamePackageListGiftGiftUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_get_detail_v3.0_data_list_game_anchor_game_package_list_gift_gift_unit value
func (v NativeAnchorGetDetailV30DataListGameAnchorGamePackageListGiftGiftUnit) Ptr() *NativeAnchorGetDetailV30DataListGameAnchorGamePackageListGiftGiftUnit {
	return &v
}
