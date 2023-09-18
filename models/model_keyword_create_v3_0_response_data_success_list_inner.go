/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordCreateV30ResponseDataSuccessListInner struct for KeywordCreateV30ResponseDataSuccessListInner
type KeywordCreateV30ResponseDataSuccessListInner struct {
	//
	Bid     *float64                                `json:"bid,omitempty"`
	BidType *KeywordCreateV30DataSuccessListBidType `json:"bid_type,omitempty"`
	//
	KeywordId *int64                                    `json:"keyword_id,omitempty"`
	MatchType *KeywordCreateV30DataSuccessListMatchType `json:"match_type,omitempty"`
	Status    *KeywordCreateV30DataSuccessListStatus    `json:"status,omitempty"`
	//
	Word *string `json:"word,omitempty"`
}
