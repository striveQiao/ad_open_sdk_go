/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeReportOrderGetV10Filtering
type QianchuanAwemeReportOrderGetV10Filtering struct {
	// 计划id列表，最多1000
	AdIds          []int64                                                 `json:"ad_ids,omitempty"`
	ExternalAction *QianchuanAwemeReportOrderGetV10FilteringExternalAction `json:"external_action,omitempty"`
	MarketingGoal  QianchuanAwemeReportOrderGetV10FilteringMarketingGoal   `json:"marketing_goal"`
}
