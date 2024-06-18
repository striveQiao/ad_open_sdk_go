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

// QianchuanAdDetailGetV10DataAudienceActionScene
type QianchuanAdDetailGetV10DataAudienceActionScene string

// List of qianchuan_ad_detail_get_v1.0_data_audience_action_scene
const (
	APP_QianchuanAdDetailGetV10DataAudienceActionScene        QianchuanAdDetailGetV10DataAudienceActionScene = "APP"
	E_COMMERCE_QianchuanAdDetailGetV10DataAudienceActionScene QianchuanAdDetailGetV10DataAudienceActionScene = "E-COMMERCE"
	NEWS_QianchuanAdDetailGetV10DataAudienceActionScene       QianchuanAdDetailGetV10DataAudienceActionScene = "NEWS"
	SEARCH_QianchuanAdDetailGetV10DataAudienceActionScene     QianchuanAdDetailGetV10DataAudienceActionScene = "SEARCH"
)

// All allowed values of QianchuanAdDetailGetV10DataAudienceActionScene enum
var AllowedQianchuanAdDetailGetV10DataAudienceActionSceneEnumValues = []QianchuanAdDetailGetV10DataAudienceActionScene{
	"APP",
	"E-COMMERCE",
	"NEWS",
	"SEARCH",
}

// NewQianchuanAdDetailGetV10DataAudienceActionSceneFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataAudienceActionScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataAudienceActionSceneFromValue(v string) (*QianchuanAdDetailGetV10DataAudienceActionScene, error) {
	ev := QianchuanAdDetailGetV10DataAudienceActionScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataAudienceActionScene: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataAudienceActionSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataAudienceActionScene) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataAudienceActionSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_audience_action_scene value
func (v QianchuanAdDetailGetV10DataAudienceActionScene) Ptr() *QianchuanAdDetailGetV10DataAudienceActionScene {
	return &v
}
