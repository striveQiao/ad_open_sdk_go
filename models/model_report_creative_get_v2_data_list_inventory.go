/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCreativeGetV2DataListInventory
type ReportCreativeGetV2DataListInventory string

// List of report_creative_get_v2_data_list_inventory
const (
	INVENTORY_BEAUTY_ReportCreativeGetV2DataListInventory       ReportCreativeGetV2DataListInventory = "INVENTORY_BEAUTY"
	INVENTORY_FACE_U_ReportCreativeGetV2DataListInventory       ReportCreativeGetV2DataListInventory = "INVENTORY_FACE_U"
	INVENTORY_TOMATO_NOVEL_ReportCreativeGetV2DataListInventory ReportCreativeGetV2DataListInventory = "INVENTORY_TOMATO_NOVEL"
	INVENTORY_SEARCH_ReportCreativeGetV2DataListInventory       ReportCreativeGetV2DataListInventory = "INVENTORY_SEARCH"
	INVENTORY_FURNISH_ReportCreativeGetV2DataListInventory      ReportCreativeGetV2DataListInventory = "INVENTORY_FURNISH"
	INVENTORY_UNION_SLOT_ReportCreativeGetV2DataListInventory   ReportCreativeGetV2DataListInventory = "INVENTORY_UNION_SLOT"
	INVENTORY_HOTSOON_FEED_ReportCreativeGetV2DataListInventory ReportCreativeGetV2DataListInventory = "INVENTORY_HOTSOON_FEED"
	INVENTORY_AWEME_FEED_ReportCreativeGetV2DataListInventory   ReportCreativeGetV2DataListInventory = "INVENTORY_AWEME_FEED"
	INVENTORY_STUDY_ReportCreativeGetV2DataListInventory        ReportCreativeGetV2DataListInventory = "INVENTORY_STUDY"
	INVENTORY_UNIVERSAL_ReportCreativeGetV2DataListInventory    ReportCreativeGetV2DataListInventory = "INVENTORY_UNIVERSAL"
	INVENTORY_PIPIXIA_ReportCreativeGetV2DataListInventory      ReportCreativeGetV2DataListInventory = "INVENTORY_PIPIXIA"
	INVENTORY_VIDEO_FEED_ReportCreativeGetV2DataListInventory   ReportCreativeGetV2DataListInventory = "INVENTORY_VIDEO_FEED"
	INVENTORY_FEED_ReportCreativeGetV2DataListInventory         ReportCreativeGetV2DataListInventory = "INVENTORY_FEED"
	UNION_BOUTIQUE_GAME_ReportCreativeGetV2DataListInventory    ReportCreativeGetV2DataListInventory = "UNION_BOUTIQUE_GAME"
	INVENTORY_AUTOMOBILE_ReportCreativeGetV2DataListInventory   ReportCreativeGetV2DataListInventory = "INVENTORY_AUTOMOBILE"
)

// All allowed values of ReportCreativeGetV2DataListInventory enum
var AllowedReportCreativeGetV2DataListInventoryEnumValues = []ReportCreativeGetV2DataListInventory{
	"INVENTORY_BEAUTY",
	"INVENTORY_FACE_U",
	"INVENTORY_TOMATO_NOVEL",
	"INVENTORY_SEARCH",
	"INVENTORY_FURNISH",
	"INVENTORY_UNION_SLOT",
	"INVENTORY_HOTSOON_FEED",
	"INVENTORY_AWEME_FEED",
	"INVENTORY_STUDY",
	"INVENTORY_UNIVERSAL",
	"INVENTORY_PIPIXIA",
	"INVENTORY_VIDEO_FEED",
	"INVENTORY_FEED",
	"UNION_BOUTIQUE_GAME",
	"INVENTORY_AUTOMOBILE",
}

// NewReportCreativeGetV2DataListInventoryFromValue returns a pointer to a valid ReportCreativeGetV2DataListInventory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCreativeGetV2DataListInventoryFromValue(v string) (*ReportCreativeGetV2DataListInventory, error) {
	ev := ReportCreativeGetV2DataListInventory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCreativeGetV2DataListInventory: valid values are %v", v, AllowedReportCreativeGetV2DataListInventoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCreativeGetV2DataListInventory) IsValid() bool {
	for _, existing := range AllowedReportCreativeGetV2DataListInventoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_creative_get_v2_data_list_inventory value
func (v ReportCreativeGetV2DataListInventory) Ptr() *ReportCreativeGetV2DataListInventory {
	return &v
}
