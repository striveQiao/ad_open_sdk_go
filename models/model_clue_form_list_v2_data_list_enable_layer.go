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

// ClueFormListV2DataListEnableLayer
type ClueFormListV2DataListEnableLayer string

// List of clue_form_list_v2_data_list_enable_layer
const (
	Enum_0_ClueFormListV2DataListEnableLayer ClueFormListV2DataListEnableLayer = "0"
	Enum_1_ClueFormListV2DataListEnableLayer ClueFormListV2DataListEnableLayer = "1"
)

// All allowed values of ClueFormListV2DataListEnableLayer enum
var AllowedClueFormListV2DataListEnableLayerEnumValues = []ClueFormListV2DataListEnableLayer{
	"0",
	"1",
}

// NewClueFormListV2DataListEnableLayerFromValue returns a pointer to a valid ClueFormListV2DataListEnableLayer
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueFormListV2DataListEnableLayerFromValue(v string) (*ClueFormListV2DataListEnableLayer, error) {
	ev := ClueFormListV2DataListEnableLayer(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueFormListV2DataListEnableLayer: valid values are %v", v, AllowedClueFormListV2DataListEnableLayerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueFormListV2DataListEnableLayer) IsValid() bool {
	for _, existing := range AllowedClueFormListV2DataListEnableLayerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_form_list_v2_data_list_enable_layer value
func (v ClueFormListV2DataListEnableLayer) Ptr() *ClueFormListV2DataListEnableLayer {
	return &v
}
