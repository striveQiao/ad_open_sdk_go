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

// QianchuanToolsAllowCouponV10MarketingScene
type QianchuanToolsAllowCouponV10MarketingScene string

// List of qianchuan_tools_allow_coupon_v1.0_marketing_scene
const (
	ALL_QianchuanToolsAllowCouponV10MarketingScene           QianchuanToolsAllowCouponV10MarketingScene = "ALL"
	FEED_QianchuanToolsAllowCouponV10MarketingScene          QianchuanToolsAllowCouponV10MarketingScene = "FEED"
	SEARCH_QianchuanToolsAllowCouponV10MarketingScene        QianchuanToolsAllowCouponV10MarketingScene = "SEARCH"
	SHOPPING_MALL_QianchuanToolsAllowCouponV10MarketingScene QianchuanToolsAllowCouponV10MarketingScene = "SHOPPING_MALL"
)

// All allowed values of QianchuanToolsAllowCouponV10MarketingScene enum
var AllowedQianchuanToolsAllowCouponV10MarketingSceneEnumValues = []QianchuanToolsAllowCouponV10MarketingScene{
	"ALL",
	"FEED",
	"SEARCH",
	"SHOPPING_MALL",
}

// NewQianchuanToolsAllowCouponV10MarketingSceneFromValue returns a pointer to a valid QianchuanToolsAllowCouponV10MarketingScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanToolsAllowCouponV10MarketingSceneFromValue(v string) (*QianchuanToolsAllowCouponV10MarketingScene, error) {
	ev := QianchuanToolsAllowCouponV10MarketingScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanToolsAllowCouponV10MarketingScene: valid values are %v", v, AllowedQianchuanToolsAllowCouponV10MarketingSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanToolsAllowCouponV10MarketingScene) IsValid() bool {
	for _, existing := range AllowedQianchuanToolsAllowCouponV10MarketingSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_tools_allow_coupon_v1.0_marketing_scene value
func (v QianchuanToolsAllowCouponV10MarketingScene) Ptr() *QianchuanToolsAllowCouponV10MarketingScene {
	return &v
}
