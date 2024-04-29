/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordPromotionUpdateV30ResponseDataErrorListInner struct for ToolsPrivativeWordPromotionUpdateV30ResponseDataErrorListInner
type ToolsPrivativeWordPromotionUpdateV30ResponseDataErrorListInner struct {
	//
	ErrorMessage *string `json:"error_message,omitempty"`
	//
	FailPhraseWords []*ToolsPrivativeWordPromotionUpdateV30ResponseDataErrorListInnerFailPhraseWordsInner `json:"fail_phrase_words,omitempty"`
	//
	FailPreciseWords []*ToolsPrivativeWordPromotionUpdateV30ResponseDataErrorListInnerFailPreciseWordsInner `json:"fail_precise_words,omitempty"`
	//
	PromotionId *int64 `json:"promotion_id,omitempty"`
}
