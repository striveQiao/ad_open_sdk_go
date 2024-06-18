/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueFormCreateV2EnableLayer
type ClueFormCreateV2EnableLayer string

// List of clue_form_create_v2_enable_layer
const (
	Enum_0_ClueFormCreateV2EnableLayer ClueFormCreateV2EnableLayer = "0"
	Enum_1_ClueFormCreateV2EnableLayer ClueFormCreateV2EnableLayer = "1"
)

// All allowed values of ClueFormCreateV2EnableLayer enum
var AllowedClueFormCreateV2EnableLayerEnumValues = []ClueFormCreateV2EnableLayer{
	"0",
	"1",
}

// NewClueFormCreateV2EnableLayerFromValue returns a pointer to a valid ClueFormCreateV2EnableLayer
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueFormCreateV2EnableLayerFromValue(v string) (*ClueFormCreateV2EnableLayer, error) {
	ev := ClueFormCreateV2EnableLayer(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueFormCreateV2EnableLayer: valid values are %v", v, AllowedClueFormCreateV2EnableLayerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueFormCreateV2EnableLayer) IsValid() bool {
	for _, existing := range AllowedClueFormCreateV2EnableLayerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_form_create_v2_enable_layer value
func (v ClueFormCreateV2EnableLayer) Ptr() *ClueFormCreateV2EnableLayer {
	return &v
}
