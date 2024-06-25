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

// QianchuanAdGetV10FilteringMarketingScene
type QianchuanAdGetV10FilteringMarketingScene string

// List of qianchuan_ad_get_v1.0_filtering_marketing_scene
const (
	ALL_QianchuanAdGetV10FilteringMarketingScene           QianchuanAdGetV10FilteringMarketingScene = "ALL"
	FEED_QianchuanAdGetV10FilteringMarketingScene          QianchuanAdGetV10FilteringMarketingScene = "FEED"
	SEARCH_QianchuanAdGetV10FilteringMarketingScene        QianchuanAdGetV10FilteringMarketingScene = "SEARCH"
	SHOPPING_MALL_QianchuanAdGetV10FilteringMarketingScene QianchuanAdGetV10FilteringMarketingScene = "SHOPPING_MALL"
)

// All allowed values of QianchuanAdGetV10FilteringMarketingScene enum
var AllowedQianchuanAdGetV10FilteringMarketingSceneEnumValues = []QianchuanAdGetV10FilteringMarketingScene{
	"ALL",
	"FEED",
	"SEARCH",
	"SHOPPING_MALL",
}

// NewQianchuanAdGetV10FilteringMarketingSceneFromValue returns a pointer to a valid QianchuanAdGetV10FilteringMarketingScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdGetV10FilteringMarketingSceneFromValue(v string) (*QianchuanAdGetV10FilteringMarketingScene, error) {
	ev := QianchuanAdGetV10FilteringMarketingScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdGetV10FilteringMarketingScene: valid values are %v", v, AllowedQianchuanAdGetV10FilteringMarketingSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdGetV10FilteringMarketingScene) IsValid() bool {
	for _, existing := range AllowedQianchuanAdGetV10FilteringMarketingSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_get_v1.0_filtering_marketing_scene value
func (v QianchuanAdGetV10FilteringMarketingScene) Ptr() *QianchuanAdGetV10FilteringMarketingScene {
	return &v
}
