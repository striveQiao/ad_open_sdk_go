/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanCampaignListGetV10Filter
type QianchuanCampaignListGetV10Filter struct {
	//
	Ids            []int64                                          `json:"ids,omitempty"`
	MarketingGoal  QianchuanCampaignListGetV10FilterMarketingGoal   `json:"marketing_goal"`
	MarketingScene *QianchuanCampaignListGetV10FilterMarketingScene `json:"marketing_scene,omitempty"`
	//
	Name   *string                                  `json:"name,omitempty"`
	Status *QianchuanCampaignListGetV10FilterStatus `json:"status,omitempty"`
}
