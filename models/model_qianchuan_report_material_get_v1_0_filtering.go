/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportMaterialGetV10Filtering
type QianchuanReportMaterialGetV10Filtering struct {
	//
	AnalysisType []*QianchuanReportMaterialGetV10FilteringAnalysisType `json:"analysis_type,omitempty"`
	//
	CarouselSource []*QianchuanReportMaterialGetV10FilteringCarouselSource `json:"carousel_source,omitempty"`
	//
	ImageSource []*QianchuanReportMaterialGetV10FilteringImageSource `json:"image_source,omitempty"`
	//
	MaterialId []int64 `json:"material_id,omitempty"`
	//
	MaterialMode []*QianchuanReportMaterialGetV10FilteringMaterialMode `json:"material_mode,omitempty"`
	MaterialType QianchuanReportMaterialGetV10FilteringMaterialType    `json:"material_type"`
	//
	VideoSource []*QianchuanReportMaterialGetV10FilteringVideoSource `json:"video_source,omitempty"`
}
