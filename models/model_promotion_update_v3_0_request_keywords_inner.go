/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionUpdateV30RequestKeywordsInner struct for PromotionUpdateV30RequestKeywordsInner
type PromotionUpdateV30RequestKeywordsInner struct {
	//
	Bid       *float64                             `json:"bid,omitempty"`
	BidType   *PromotionUpdateV30KeywordsBidType   `json:"bid_type,omitempty"`
	MatchType *PromotionUpdateV30KeywordsMatchType `json:"match_type,omitempty"`
	//
	Word string `json:"word"`
}
