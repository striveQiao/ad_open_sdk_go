/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportAdvertiserGetV10Filtering
type QianchuanReportAdvertiserGetV10Filtering struct {
	//
	AwemeIds       []int64                                                 `json:"aweme_ids,omitempty"`
	CampaignScene  *QianchuanReportAdvertiserGetV10FilteringCampaignScene  `json:"campaign_scene,omitempty"`
	MarketingGoal  QianchuanReportAdvertiserGetV10FilteringMarketingGoal   `json:"marketing_goal"`
	MarketingScene *QianchuanReportAdvertiserGetV10FilteringMarketingScene `json:"marketing_scene,omitempty"`
	OrderPlatform  *QianchuanReportAdvertiserGetV10FilteringOrderPlatform  `json:"order_platform,omitempty"`
	SmartBidType   *QianchuanReportAdvertiserGetV10FilteringSmartBidType   `json:"smart_bid_type,omitempty"`
	Status         *QianchuanReportAdvertiserGetV10FilteringStatus         `json:"status,omitempty"`
}
