/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SuggWordsV30RequestPromotionMaterials
type SuggWordsV30RequestPromotionMaterials struct {
	//
	Abstracts []string `json:"abstracts,omitempty"`
	//
	AppName *string `json:"app_name,omitempty"`
	//
	QualityWords []string `json:"quality_words,omitempty"`
	//
	Source *string `json:"source,omitempty"`
	//
	Titles []string `json:"titles,omitempty"`
}
