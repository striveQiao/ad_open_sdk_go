/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionListV30ResponseDataListInnerKeywordsInner struct for PromotionListV30ResponseDataListInnerKeywordsInner
type PromotionListV30ResponseDataListInnerKeywordsInner struct {
	//
	Bid       *float64                                   `json:"bid,omitempty"`
	BidType   *PromotionListV30DataListKeywordsBidType   `json:"bid_type,omitempty"`
	MatchType *PromotionListV30DataListKeywordsMatchType `json:"match_type,omitempty"`
	//
	Word *string `json:"word,omitempty"`
}
