/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordGetV2ResponseDataAdsPrivativeInner struct for ToolsPrivativeWordGetV2ResponseDataAdsPrivativeInner
type ToolsPrivativeWordGetV2ResponseDataAdsPrivativeInner struct {
	//
	AdId *int64 `json:"ad_id,omitempty"`
	//
	PhraseWords []string `json:"phrase_words,omitempty"`
	//
	PreciseWords []string `json:"precise_words,omitempty"`
}
