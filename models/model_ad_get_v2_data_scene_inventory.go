/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2DataSceneInventory
type AdGetV2DataSceneInventory string

// List of ad_get_v2_data_scene_inventory
const (
	TAIL_SCENE_AdGetV2DataSceneInventory  AdGetV2DataSceneInventory = "TAIL_SCENE"
	VIDEO_SCENE_AdGetV2DataSceneInventory AdGetV2DataSceneInventory = "VIDEO_SCENE"
	NOT_SELECT_AdGetV2DataSceneInventory  AdGetV2DataSceneInventory = "NOT_SELECT"
	FEED_SCENE_AdGetV2DataSceneInventory  AdGetV2DataSceneInventory = "FEED_SCENE"
)

// Ptr returns reference to ad_get_v2_data_scene_inventory value
func (v AdGetV2DataSceneInventory) Ptr() *AdGetV2DataSceneInventory {
	return &v
}
