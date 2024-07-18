/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordUpdateV2V2ResponseDataErrorListInner struct for KeywordUpdateV2V2ResponseDataErrorListInner
type KeywordUpdateV2V2ResponseDataErrorListInner struct {
	//
	Bid     *float64                               `json:"bid,omitempty"`
	BidType *KeywordUpdateV2V2DataErrorListBidType `json:"bid_type,omitempty"`
	//
	ErrorReason *string `json:"error_reason,omitempty"`
	//
	IsPause *int64 `json:"is_pause,omitempty"`
	//
	KeywordId *int64                                   `json:"keyword_id,omitempty"`
	MatchType *KeywordUpdateV2V2DataErrorListMatchType `json:"match_type,omitempty"`
	//
	WordId *int64 `json:"word_id,omitempty"`
}
