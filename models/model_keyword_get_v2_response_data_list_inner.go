/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordGetV2ResponseDataListInner struct for KeywordGetV2ResponseDataListInner
type KeywordGetV2ResponseDataListInner struct {
	//
	Bid     *float64                     `json:"bid,omitempty"`
	BidType *KeywordGetV2DataListBidType `json:"bid_type,omitempty"`
	//
	IsPause *int64 `json:"is_pause,omitempty"`
	//
	KeywordId *int64                         `json:"keyword_id,omitempty"`
	MatchType *KeywordGetV2DataListMatchType `json:"match_type,omitempty"`
	//
	Word *string `json:"word,omitempty"`
	//
	WordId *int64 `json:"word_id,omitempty"`
}
