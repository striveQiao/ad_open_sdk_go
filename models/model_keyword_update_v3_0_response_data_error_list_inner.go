/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordUpdateV30ResponseDataErrorListInner struct for KeywordUpdateV30ResponseDataErrorListInner
type KeywordUpdateV30ResponseDataErrorListInner struct {
	//
	ErrorReason *string `json:"error_reason,omitempty"`
	//
	KeywordId *int64 `json:"keyword_id,omitempty"`
	//
	Word *string `json:"word,omitempty"`
	//
	WordId *int64 `json:"word_id,omitempty"`
}
