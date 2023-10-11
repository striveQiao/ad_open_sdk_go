/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPromotionDiagnosisSuggestionAcceptV30Request struct for ToolsPromotionDiagnosisSuggestionAcceptV30Request
type ToolsPromotionDiagnosisSuggestionAcceptV30Request struct {
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	DiagnosisId *string `json:"diagnosis_id,omitempty"`
	//
	PromotionId *int64 `json:"promotion_id,omitempty"`
	//
	Tools []*ToolsPromotionDiagnosisSuggestionAcceptV30RequestToolsInner `json:"tools,omitempty"`
}
