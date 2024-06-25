/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordUpdateV30ResponseDataSuccessListInner struct for KeywordUpdateV30ResponseDataSuccessListInner
type KeywordUpdateV30ResponseDataSuccessListInner struct {
	//
	Bid     *float64                                `json:"bid,omitempty"`
	BidType *KeywordUpdateV30DataSuccessListBidType `json:"bid_type,omitempty"`
	//
	IsPause *int64 `json:"is_pause,omitempty"`
	//
	KeywordId *int64                                    `json:"keyword_id,omitempty"`
	MatchType *KeywordUpdateV30DataSuccessListMatchType `json:"match_type,omitempty"`
	//
	Word *string `json:"word,omitempty"`
	//
	WordId *int64 `json:"word_id,omitempty"`
}
