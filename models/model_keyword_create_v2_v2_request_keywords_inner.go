/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordCreateV2V2RequestKeywordsInner struct for KeywordCreateV2V2RequestKeywordsInner
type KeywordCreateV2V2RequestKeywordsInner struct {
	//
	Bid       *float64                           `json:"bid,omitempty"`
	BidType   *KeywordCreateV2V2KeywordsBidType  `json:"bid_type,omitempty"`
	MatchType KeywordCreateV2V2KeywordsMatchType `json:"match_type"`
	//
	Word string `json:"word"`
}
