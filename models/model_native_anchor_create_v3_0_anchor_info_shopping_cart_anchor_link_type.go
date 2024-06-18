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

// NativeAnchorCreateV30AnchorInfoShoppingCartAnchorLinkType
type NativeAnchorCreateV30AnchorInfoShoppingCartAnchorLinkType string

// List of native_anchor_create_v3.0_anchor_info_shopping_cart_anchor_link_type
const (
	ONE_JUMP_NativeAnchorCreateV30AnchorInfoShoppingCartAnchorLinkType NativeAnchorCreateV30AnchorInfoShoppingCartAnchorLinkType = "ONE_JUMP"
	TWO_JUMP_NativeAnchorCreateV30AnchorInfoShoppingCartAnchorLinkType NativeAnchorCreateV30AnchorInfoShoppingCartAnchorLinkType = "TWO_JUMP"
)

// All allowed values of NativeAnchorCreateV30AnchorInfoShoppingCartAnchorLinkType enum
var AllowedNativeAnchorCreateV30AnchorInfoShoppingCartAnchorLinkTypeEnumValues = []NativeAnchorCreateV30AnchorInfoShoppingCartAnchorLinkType{
	"ONE_JUMP",
	"TWO_JUMP",
}

// NewNativeAnchorCreateV30AnchorInfoShoppingCartAnchorLinkTypeFromValue returns a pointer to a valid NativeAnchorCreateV30AnchorInfoShoppingCartAnchorLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorCreateV30AnchorInfoShoppingCartAnchorLinkTypeFromValue(v string) (*NativeAnchorCreateV30AnchorInfoShoppingCartAnchorLinkType, error) {
	ev := NativeAnchorCreateV30AnchorInfoShoppingCartAnchorLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorCreateV30AnchorInfoShoppingCartAnchorLinkType: valid values are %v", v, AllowedNativeAnchorCreateV30AnchorInfoShoppingCartAnchorLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorCreateV30AnchorInfoShoppingCartAnchorLinkType) IsValid() bool {
	for _, existing := range AllowedNativeAnchorCreateV30AnchorInfoShoppingCartAnchorLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_create_v3.0_anchor_info_shopping_cart_anchor_link_type value
func (v NativeAnchorCreateV30AnchorInfoShoppingCartAnchorLinkType) Ptr() *NativeAnchorCreateV30AnchorInfoShoppingCartAnchorLinkType {
	return &v
}
