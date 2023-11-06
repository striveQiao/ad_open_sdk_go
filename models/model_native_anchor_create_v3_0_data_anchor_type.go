/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// NativeAnchorCreateV30DataAnchorType
type NativeAnchorCreateV30DataAnchorType string

// List of native_anchor_create_v3.0_data_anchor_type
const (
	APP_GAME_NativeAnchorCreateV30DataAnchorType             NativeAnchorCreateV30DataAnchorType = "APP_GAME"
	APP_INTERNET_SERVICE_NativeAnchorCreateV30DataAnchorType NativeAnchorCreateV30DataAnchorType = "APP_INTERNET_SERVICE"
	APP_SHOP_NativeAnchorCreateV30DataAnchorType             NativeAnchorCreateV30DataAnchorType = "APP_SHOP"
	MICRO_APP_NativeAnchorCreateV30DataAnchorType            NativeAnchorCreateV30DataAnchorType = "MICRO_APP"
	ONLINE_SUBSCRIBE_NativeAnchorCreateV30DataAnchorType     NativeAnchorCreateV30DataAnchorType = "ONLINE_SUBSCRIBE"
	PRIVATE_CHAT_NativeAnchorCreateV30DataAnchorType         NativeAnchorCreateV30DataAnchorType = "PRIVATE_CHAT"
	SHOPPING_CART_NativeAnchorCreateV30DataAnchorType        NativeAnchorCreateV30DataAnchorType = "SHOPPING_CART"
)

// All allowed values of NativeAnchorCreateV30DataAnchorType enum
var AllowedNativeAnchorCreateV30DataAnchorTypeEnumValues = []NativeAnchorCreateV30DataAnchorType{
	"APP_GAME",
	"APP_INTERNET_SERVICE",
	"APP_SHOP",
	"MICRO_APP",
	"ONLINE_SUBSCRIBE",
	"PRIVATE_CHAT",
	"SHOPPING_CART",
}

// NewNativeAnchorCreateV30DataAnchorTypeFromValue returns a pointer to a valid NativeAnchorCreateV30DataAnchorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorCreateV30DataAnchorTypeFromValue(v string) (*NativeAnchorCreateV30DataAnchorType, error) {
	ev := NativeAnchorCreateV30DataAnchorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorCreateV30DataAnchorType: valid values are %v", v, AllowedNativeAnchorCreateV30DataAnchorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorCreateV30DataAnchorType) IsValid() bool {
	for _, existing := range AllowedNativeAnchorCreateV30DataAnchorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_create_v3.0_data_anchor_type value
func (v NativeAnchorCreateV30DataAnchorType) Ptr() *NativeAnchorCreateV30DataAnchorType {
	return &v
}
