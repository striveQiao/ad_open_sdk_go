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

// QianchuanAdGetV10DataListCampaignScene
type QianchuanAdGetV10DataListCampaignScene string

// List of qianchuan_ad_get_v1.0_data_list_campaign_scene
const (
	DAILY_SALE_QianchuanAdGetV10DataListCampaignScene                  QianchuanAdGetV10DataListCampaignScene = "DAILY_SALE"
	LIVE_HEAT_QianchuanAdGetV10DataListCampaignScene                   QianchuanAdGetV10DataListCampaignScene = "LIVE_HEAT"
	NEW_CUSTOMER_TRANSFORMATION_QianchuanAdGetV10DataListCampaignScene QianchuanAdGetV10DataListCampaignScene = "NEW_CUSTOMER_TRANSFORMATION"
	NEW_PRODUCT_BOOST_QianchuanAdGetV10DataListCampaignScene           QianchuanAdGetV10DataListCampaignScene = "NEW_PRODUCT_BOOST"
	PLANT_GRASS_QianchuanAdGetV10DataListCampaignScene                 QianchuanAdGetV10DataListCampaignScene = "PLANT_GRASS"
	PRODUCT_HEAT_QianchuanAdGetV10DataListCampaignScene                QianchuanAdGetV10DataListCampaignScene = "PRODUCT_HEAT"
)

// All allowed values of QianchuanAdGetV10DataListCampaignScene enum
var AllowedQianchuanAdGetV10DataListCampaignSceneEnumValues = []QianchuanAdGetV10DataListCampaignScene{
	"DAILY_SALE",
	"LIVE_HEAT",
	"NEW_CUSTOMER_TRANSFORMATION",
	"NEW_PRODUCT_BOOST",
	"PLANT_GRASS",
	"PRODUCT_HEAT",
}

// NewQianchuanAdGetV10DataListCampaignSceneFromValue returns a pointer to a valid QianchuanAdGetV10DataListCampaignScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdGetV10DataListCampaignSceneFromValue(v string) (*QianchuanAdGetV10DataListCampaignScene, error) {
	ev := QianchuanAdGetV10DataListCampaignScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdGetV10DataListCampaignScene: valid values are %v", v, AllowedQianchuanAdGetV10DataListCampaignSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdGetV10DataListCampaignScene) IsValid() bool {
	for _, existing := range AllowedQianchuanAdGetV10DataListCampaignSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_get_v1.0_data_list_campaign_scene value
func (v QianchuanAdGetV10DataListCampaignScene) Ptr() *QianchuanAdGetV10DataListCampaignScene {
	return &v
}
