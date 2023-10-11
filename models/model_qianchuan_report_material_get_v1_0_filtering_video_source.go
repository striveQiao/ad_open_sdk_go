/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportMaterialGetV10FilteringVideoSource
type QianchuanReportMaterialGetV10FilteringVideoSource string

// List of qianchuan_report_material_get_v1.0_filtering_video_source
const (
	ARTHUR_QianchuanReportMaterialGetV10FilteringVideoSource          QianchuanReportMaterialGetV10FilteringVideoSource = "ARTHUR"
	AWEME_QianchuanReportMaterialGetV10FilteringVideoSource           QianchuanReportMaterialGetV10FilteringVideoSource = "AWEME"
	BP_QianchuanReportMaterialGetV10FilteringVideoSource              QianchuanReportMaterialGetV10FilteringVideoSource = "BP"
	CREATIVE_CENTER_QianchuanReportMaterialGetV10FilteringVideoSource QianchuanReportMaterialGetV10FilteringVideoSource = "CREATIVE_CENTER"
	E_COMMERCE_QianchuanReportMaterialGetV10FilteringVideoSource      QianchuanReportMaterialGetV10FilteringVideoSource = "E_COMMERCE"
	JI_CHUANG_QianchuanReportMaterialGetV10FilteringVideoSource       QianchuanReportMaterialGetV10FilteringVideoSource = "JI_CHUANG"
	LIVE_HIGHLIGHT_QianchuanReportMaterialGetV10FilteringVideoSource  QianchuanReportMaterialGetV10FilteringVideoSource = "LIVE_HIGHLIGHT"
	STAR_QianchuanReportMaterialGetV10FilteringVideoSource            QianchuanReportMaterialGetV10FilteringVideoSource = "STAR"
	TADA_QianchuanReportMaterialGetV10FilteringVideoSource            QianchuanReportMaterialGetV10FilteringVideoSource = "TADA"
	VIDEO_CAPTURE_QianchuanReportMaterialGetV10FilteringVideoSource   QianchuanReportMaterialGetV10FilteringVideoSource = "VIDEO_CAPTURE"
)

// All allowed values of QianchuanReportMaterialGetV10FilteringVideoSource enum
var AllowedQianchuanReportMaterialGetV10FilteringVideoSourceEnumValues = []QianchuanReportMaterialGetV10FilteringVideoSource{
	"ARTHUR",
	"AWEME",
	"BP",
	"CREATIVE_CENTER",
	"E_COMMERCE",
	"JI_CHUANG",
	"LIVE_HIGHLIGHT",
	"STAR",
	"TADA",
	"VIDEO_CAPTURE",
}

// NewQianchuanReportMaterialGetV10FilteringVideoSourceFromValue returns a pointer to a valid QianchuanReportMaterialGetV10FilteringVideoSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportMaterialGetV10FilteringVideoSourceFromValue(v string) (*QianchuanReportMaterialGetV10FilteringVideoSource, error) {
	ev := QianchuanReportMaterialGetV10FilteringVideoSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportMaterialGetV10FilteringVideoSource: valid values are %v", v, AllowedQianchuanReportMaterialGetV10FilteringVideoSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportMaterialGetV10FilteringVideoSource) IsValid() bool {
	for _, existing := range AllowedQianchuanReportMaterialGetV10FilteringVideoSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_material_get_v1.0_filtering_video_source value
func (v QianchuanReportMaterialGetV10FilteringVideoSource) Ptr() *QianchuanReportMaterialGetV10FilteringVideoSource {
	return &v
}
