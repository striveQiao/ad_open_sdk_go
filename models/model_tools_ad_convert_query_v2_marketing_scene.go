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

// ToolsAdConvertQueryV2MarketingScene
type ToolsAdConvertQueryV2MarketingScene string

// List of tools_ad_convert_query_v2_marketing_scene
const (
	NORMAL_ToolsAdConvertQueryV2MarketingScene         ToolsAdConvertQueryV2MarketingScene = "NORMAL"
	GAME_SUBSCRIBE_ToolsAdConvertQueryV2MarketingScene ToolsAdConvertQueryV2MarketingScene = "GAME_SUBSCRIBE"
	GAME_PROMOTION_ToolsAdConvertQueryV2MarketingScene ToolsAdConvertQueryV2MarketingScene = "GAME_PROMOTION"
)

// All allowed values of ToolsAdConvertQueryV2MarketingScene enum
var AllowedToolsAdConvertQueryV2MarketingSceneEnumValues = []ToolsAdConvertQueryV2MarketingScene{
	"NORMAL",
	"GAME_SUBSCRIBE",
	"GAME_PROMOTION",
}

// NewToolsAdConvertQueryV2MarketingSceneFromValue returns a pointer to a valid ToolsAdConvertQueryV2MarketingScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdConvertQueryV2MarketingSceneFromValue(v string) (*ToolsAdConvertQueryV2MarketingScene, error) {
	ev := ToolsAdConvertQueryV2MarketingScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdConvertQueryV2MarketingScene: valid values are %v", v, AllowedToolsAdConvertQueryV2MarketingSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdConvertQueryV2MarketingScene) IsValid() bool {
	for _, existing := range AllowedToolsAdConvertQueryV2MarketingSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ad_convert_query_v2_marketing_scene value
func (v ToolsAdConvertQueryV2MarketingScene) Ptr() *ToolsAdConvertQueryV2MarketingScene {
	return &v
}
