/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionListV30ResponseDataListInnerKeywordsInner struct for PromotionListV30ResponseDataListInnerKeywordsInner
type PromotionListV30ResponseDataListInnerKeywordsInner struct {
	//
	Bid       *float32                                   `json:"bid,omitempty"`
	MatchType *PromotionListV30DataListKeywordsMatchType `json:"match_type,omitempty"`
	//
	Word *string `json:"word,omitempty"`
}
