/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordFeedadsSuggestV2ResponseDataListInner struct for KeywordFeedadsSuggestV2ResponseDataListInner
type KeywordFeedadsSuggestV2ResponseDataListInner struct {
	//
	Keyword   *string                                   `json:"keyword,omitempty"`
	MatchType *KeywordFeedadsSuggestV2DataListMatchType `json:"match_type,omitempty"`
}
