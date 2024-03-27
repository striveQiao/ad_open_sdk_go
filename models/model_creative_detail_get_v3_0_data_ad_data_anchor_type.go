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

// CreativeDetailGetV30DataAdDataAnchorType
type CreativeDetailGetV30DataAdDataAnchorType string

// List of creative_detail_get_v3.0_data_ad_data_anchor_type
const (
	APP_GAME_CreativeDetailGetV30DataAdDataAnchorType             CreativeDetailGetV30DataAdDataAnchorType = "APP_GAME"
	APP_INTERNET_SERVICE_CreativeDetailGetV30DataAdDataAnchorType CreativeDetailGetV30DataAdDataAnchorType = "APP_INTERNET_SERVICE"
	APP_SHOP_CreativeDetailGetV30DataAdDataAnchorType             CreativeDetailGetV30DataAdDataAnchorType = "APP_SHOP"
	INSURANCE_CreativeDetailGetV30DataAdDataAnchorType            CreativeDetailGetV30DataAdDataAnchorType = "INSURANCE"
	MICRO_APP_CreativeDetailGetV30DataAdDataAnchorType            CreativeDetailGetV30DataAdDataAnchorType = "MICRO_APP"
	ONLINE_SUBSCRIBE_CreativeDetailGetV30DataAdDataAnchorType     CreativeDetailGetV30DataAdDataAnchorType = "ONLINE_SUBSCRIBE"
	PRIVATE_CHAT_CreativeDetailGetV30DataAdDataAnchorType         CreativeDetailGetV30DataAdDataAnchorType = "PRIVATE_CHAT"
	SHOPPING_CART_CreativeDetailGetV30DataAdDataAnchorType        CreativeDetailGetV30DataAdDataAnchorType = "SHOPPING_CART"
)

// All allowed values of CreativeDetailGetV30DataAdDataAnchorType enum
var AllowedCreativeDetailGetV30DataAdDataAnchorTypeEnumValues = []CreativeDetailGetV30DataAdDataAnchorType{
	"APP_GAME",
	"APP_INTERNET_SERVICE",
	"APP_SHOP",
	"INSURANCE",
	"MICRO_APP",
	"ONLINE_SUBSCRIBE",
	"PRIVATE_CHAT",
	"SHOPPING_CART",
}

// NewCreativeDetailGetV30DataAdDataAnchorTypeFromValue returns a pointer to a valid CreativeDetailGetV30DataAdDataAnchorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeDetailGetV30DataAdDataAnchorTypeFromValue(v string) (*CreativeDetailGetV30DataAdDataAnchorType, error) {
	ev := CreativeDetailGetV30DataAdDataAnchorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeDetailGetV30DataAdDataAnchorType: valid values are %v", v, AllowedCreativeDetailGetV30DataAdDataAnchorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeDetailGetV30DataAdDataAnchorType) IsValid() bool {
	for _, existing := range AllowedCreativeDetailGetV30DataAdDataAnchorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_detail_get_v3.0_data_ad_data_anchor_type value
func (v CreativeDetailGetV30DataAdDataAnchorType) Ptr() *CreativeDetailGetV30DataAdDataAnchorType {
	return &v
}
