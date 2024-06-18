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

// AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryType
type AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryType string

// List of adlab_group_list_v3.0_data_groups_ad_info_delivery_range_inventory_type
const (
	INVENTORY_AWEME_FEED_AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryType      AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryType = "INVENTORY_AWEME_FEED"
	INVENTORY_FEED_AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryType            AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryType = "INVENTORY_FEED"
	INVENTORY_HOMED_AGGREGATE_AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryType AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryType = "INVENTORY_HOMED_AGGREGATE"
	INVENTORY_HOTSOON_FEED_AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryType    AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryType = "INVENTORY_HOTSOON_FEED"
	INVENTORY_TOMATO_NOVEL_AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryType    AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryType = "INVENTORY_TOMATO_NOVEL"
	INVENTORY_UNION_SLOT_AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryType      AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryType = "INVENTORY_UNION_SLOT"
	INVENTORY_VIDEO_FEED_AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryType      AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryType = "INVENTORY_VIDEO_FEED"
	UNION_BOUTIQUE_GAME_AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryType       AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryType = "UNION_BOUTIQUE_GAME"
)

// All allowed values of AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryType enum
var AllowedAdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryTypeEnumValues = []AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryType{
	"INVENTORY_AWEME_FEED",
	"INVENTORY_FEED",
	"INVENTORY_HOMED_AGGREGATE",
	"INVENTORY_HOTSOON_FEED",
	"INVENTORY_TOMATO_NOVEL",
	"INVENTORY_UNION_SLOT",
	"INVENTORY_VIDEO_FEED",
	"UNION_BOUTIQUE_GAME",
}

// NewAdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryTypeFromValue returns a pointer to a valid AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryTypeFromValue(v string) (*AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryType, error) {
	ev := AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryType: valid values are %v", v, AllowedAdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_list_v3.0_data_groups_ad_info_delivery_range_inventory_type value
func (v AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryType) Ptr() *AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryType {
	return &v
}
