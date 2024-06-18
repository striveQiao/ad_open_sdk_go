/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordListV30ResponseDataListInner struct for KeywordListV30ResponseDataListInner
type KeywordListV30ResponseDataListInner struct {
	//
	Bid     *float64                       `json:"bid,omitempty"`
	BidType *KeywordListV30DataListBidType `json:"bid_type,omitempty"`
	//
	IsPause *int64 `json:"is_pause,omitempty"`
	//
	KeywordId *int64                           `json:"keyword_id,omitempty"`
	MatchType *KeywordListV30DataListMatchType `json:"match_type,omitempty"`
	Status    *KeywordListV30DataListStatus    `json:"status,omitempty"`
	//
	Word *string `json:"word,omitempty"`
	//
	WordId *int64 `json:"word_id,omitempty"`
}
