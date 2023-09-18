/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanCampaignListGetV10DataListMarketingGoal
type QianchuanCampaignListGetV10DataListMarketingGoal string

// List of qianchuan_campaign_list_get_v1.0_data_list_marketing_goal
const (
	ALL_QianchuanCampaignListGetV10DataListMarketingGoal              QianchuanCampaignListGetV10DataListMarketingGoal = "ALL"
	LIVE_PROM_GOODS_QianchuanCampaignListGetV10DataListMarketingGoal  QianchuanCampaignListGetV10DataListMarketingGoal = "LIVE_PROM_GOODS"
	VIDEO_PROM_GOODS_QianchuanCampaignListGetV10DataListMarketingGoal QianchuanCampaignListGetV10DataListMarketingGoal = "VIDEO_PROM_GOODS"
)

// All allowed values of QianchuanCampaignListGetV10DataListMarketingGoal enum
var AllowedQianchuanCampaignListGetV10DataListMarketingGoalEnumValues = []QianchuanCampaignListGetV10DataListMarketingGoal{
	"ALL",
	"LIVE_PROM_GOODS",
	"VIDEO_PROM_GOODS",
}

// NewQianchuanCampaignListGetV10DataListMarketingGoalFromValue returns a pointer to a valid QianchuanCampaignListGetV10DataListMarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCampaignListGetV10DataListMarketingGoalFromValue(v string) (*QianchuanCampaignListGetV10DataListMarketingGoal, error) {
	ev := QianchuanCampaignListGetV10DataListMarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCampaignListGetV10DataListMarketingGoal: valid values are %v", v, AllowedQianchuanCampaignListGetV10DataListMarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCampaignListGetV10DataListMarketingGoal) IsValid() bool {
	for _, existing := range AllowedQianchuanCampaignListGetV10DataListMarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_campaign_list_get_v1.0_data_list_marketing_goal value
func (v QianchuanCampaignListGetV10DataListMarketingGoal) Ptr() *QianchuanCampaignListGetV10DataListMarketingGoal {
	return &v
}
