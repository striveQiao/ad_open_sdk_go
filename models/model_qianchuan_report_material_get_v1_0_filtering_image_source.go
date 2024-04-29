/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportMaterialGetV10FilteringImageSource
type QianchuanReportMaterialGetV10FilteringImageSource string

// List of qianchuan_report_material_get_v1.0_filtering_image_source
const (
	AI_GENERATE_QianchuanReportMaterialGetV10FilteringImageSource     QianchuanReportMaterialGetV10FilteringImageSource = "AI_GENERATE"
	CREATIVE_CENTER_QianchuanReportMaterialGetV10FilteringImageSource QianchuanReportMaterialGetV10FilteringImageSource = "CREATIVE_CENTER"
	E_COMMERCE_QianchuanReportMaterialGetV10FilteringImageSource      QianchuanReportMaterialGetV10FilteringImageSource = "E_COMMERCE"
	JI_CHUANG_QianchuanReportMaterialGetV10FilteringImageSource       QianchuanReportMaterialGetV10FilteringImageSource = "JI_CHUANG"
	SQUARE_QianchuanReportMaterialGetV10FilteringImageSource          QianchuanReportMaterialGetV10FilteringImageSource = "SQUARE"
)

// All allowed values of QianchuanReportMaterialGetV10FilteringImageSource enum
var AllowedQianchuanReportMaterialGetV10FilteringImageSourceEnumValues = []QianchuanReportMaterialGetV10FilteringImageSource{
	"AI_GENERATE",
	"CREATIVE_CENTER",
	"E_COMMERCE",
	"JI_CHUANG",
	"SQUARE",
}

// NewQianchuanReportMaterialGetV10FilteringImageSourceFromValue returns a pointer to a valid QianchuanReportMaterialGetV10FilteringImageSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportMaterialGetV10FilteringImageSourceFromValue(v string) (*QianchuanReportMaterialGetV10FilteringImageSource, error) {
	ev := QianchuanReportMaterialGetV10FilteringImageSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportMaterialGetV10FilteringImageSource: valid values are %v", v, AllowedQianchuanReportMaterialGetV10FilteringImageSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportMaterialGetV10FilteringImageSource) IsValid() bool {
	for _, existing := range AllowedQianchuanReportMaterialGetV10FilteringImageSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_material_get_v1.0_filtering_image_source value
func (v QianchuanReportMaterialGetV10FilteringImageSource) Ptr() *QianchuanReportMaterialGetV10FilteringImageSource {
	return &v
}
