/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2DataInventoryType
type AdGetV2DataInventoryType string

// List of ad_get_v2_data_inventory_type
const (
	INVENTORY_FEED_AdGetV2DataInventoryType              AdGetV2DataInventoryType = "INVENTORY_FEED"
	INVENTORY_SEARCH_AdGetV2DataInventoryType            AdGetV2DataInventoryType = "INVENTORY_SEARCH"
	INVENTORY_TEXT_LINK_AdGetV2DataInventoryType         AdGetV2DataInventoryType = "INVENTORY_TEXT_LINK"
	UNION_BOUTIQUE_GAME_AdGetV2DataInventoryType         AdGetV2DataInventoryType = "UNION_BOUTIQUE_GAME"
	INVENTORY_AWEME_SEARCH_AdGetV2DataInventoryType      AdGetV2DataInventoryType = "INVENTORY_AWEME_SEARCH"
	INVENTORY_ESSAY_FEED_AdGetV2DataInventoryType        AdGetV2DataInventoryType = "INVENTORY_ESSAY_FEED"
	INVENTORY_UNION_SLOT_AdGetV2DataInventoryType        AdGetV2DataInventoryType = "INVENTORY_UNION_SLOT"
	INVENTORY_TOMATO_NOVEL_AdGetV2DataInventoryType      AdGetV2DataInventoryType = "INVENTORY_TOMATO_NOVEL"
	INVENTORY_HOTSOON_FEED_AdGetV2DataInventoryType      AdGetV2DataInventoryType = "INVENTORY_HOTSOON_FEED"
	INVENTORY_AWEME_FEED_AdGetV2DataInventoryType        AdGetV2DataInventoryType = "INVENTORY_AWEME_FEED"
	INVENTORY_UNION_SPLASH_SLOT_AdGetV2DataInventoryType AdGetV2DataInventoryType = "INVENTORY_UNION_SPLASH_SLOT"
	INVENTORY_VIDEO_FEED_AdGetV2DataInventoryType        AdGetV2DataInventoryType = "INVENTORY_VIDEO_FEED"
)

// Ptr returns reference to ad_get_v2_data_inventory_type value
func (v AdGetV2DataInventoryType) Ptr() *AdGetV2DataInventoryType {
	return &v
}
