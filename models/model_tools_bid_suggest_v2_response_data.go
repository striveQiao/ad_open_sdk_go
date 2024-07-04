/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsBidSuggestV2ResponseData
type ToolsBidSuggestV2ResponseData struct {
	//
	BidHigh50 *float64 `json:"bid_high_50,omitempty"`
	//
	BidHigh90 *float64 `json:"bid_high_90,omitempty"`
	//
	SmartBidRange []float64 `json:"smart_bid_range,omitempty"`
	//
	SmartBidSuggestBudget *int64 `json:"smart_bid_suggest_budget,omitempty"`
	//
	SmartBudgetRange []int64 `json:"smart_budget_range,omitempty"`
	//
	SuggestedBid *float64 `json:"suggested_bid,omitempty"`
}
