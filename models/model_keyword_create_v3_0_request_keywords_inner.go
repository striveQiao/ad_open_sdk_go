/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordCreateV30RequestKeywordsInner struct for KeywordCreateV30RequestKeywordsInner
type KeywordCreateV30RequestKeywordsInner struct {
	//
	Bid       *float64                          `json:"bid,omitempty"`
	MatchType KeywordCreateV30KeywordsMatchType `json:"match_type"`
	//
	Word string `json:"word"`
}
