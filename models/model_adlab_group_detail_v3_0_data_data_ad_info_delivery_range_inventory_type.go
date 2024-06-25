/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryType
type AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryType string

// List of adlab_group_detail_v3.0_data_data_ad_info_delivery_range_inventory_type
const (
	INVENTORY_AWEME_FEED_AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryType      AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryType = "INVENTORY_AWEME_FEED"
	INVENTORY_FEED_AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryType            AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryType = "INVENTORY_FEED"
	INVENTORY_HOMED_AGGREGATE_AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryType AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryType = "INVENTORY_HOMED_AGGREGATE"
	INVENTORY_HOTSOON_FEED_AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryType    AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryType = "INVENTORY_HOTSOON_FEED"
	INVENTORY_TOMATO_NOVEL_AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryType    AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryType = "INVENTORY_TOMATO_NOVEL"
	INVENTORY_UNION_SLOT_AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryType      AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryType = "INVENTORY_UNION_SLOT"
	INVENTORY_VIDEO_FEED_AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryType      AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryType = "INVENTORY_VIDEO_FEED"
	UNION_BOUTIQUE_GAME_AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryType       AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryType = "UNION_BOUTIQUE_GAME"
)

// All allowed values of AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryType enum
var AllowedAdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryTypeEnumValues = []AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryType{
	"INVENTORY_AWEME_FEED",
	"INVENTORY_FEED",
	"INVENTORY_HOMED_AGGREGATE",
	"INVENTORY_HOTSOON_FEED",
	"INVENTORY_TOMATO_NOVEL",
	"INVENTORY_UNION_SLOT",
	"INVENTORY_VIDEO_FEED",
	"UNION_BOUTIQUE_GAME",
}

// NewAdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryTypeFromValue returns a pointer to a valid AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryTypeFromValue(v string) (*AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryType, error) {
	ev := AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryType: valid values are %v", v, AllowedAdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_detail_v3.0_data_data_ad_info_delivery_range_inventory_type value
func (v AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryType) Ptr() *AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryType {
	return &v
}
