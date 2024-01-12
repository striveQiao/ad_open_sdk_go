/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportMaterialGetV10DataListVideoSource
type QianchuanReportMaterialGetV10DataListVideoSource string

// List of qianchuan_report_material_get_v1.0_data_list_video_source
const (
	ARTHUR_QianchuanReportMaterialGetV10DataListVideoSource          QianchuanReportMaterialGetV10DataListVideoSource = "ARTHUR"
	AWEME_QianchuanReportMaterialGetV10DataListVideoSource           QianchuanReportMaterialGetV10DataListVideoSource = "AWEME"
	BP_QianchuanReportMaterialGetV10DataListVideoSource              QianchuanReportMaterialGetV10DataListVideoSource = "BP"
	CREATIVE_CENTER_QianchuanReportMaterialGetV10DataListVideoSource QianchuanReportMaterialGetV10DataListVideoSource = "CREATIVE_CENTER"
	E_COMMERCE_QianchuanReportMaterialGetV10DataListVideoSource      QianchuanReportMaterialGetV10DataListVideoSource = "E_COMMERCE"
	JI_CHUANG_QianchuanReportMaterialGetV10DataListVideoSource       QianchuanReportMaterialGetV10DataListVideoSource = "JI_CHUANG"
	LIVE_HIGHLIGHT_QianchuanReportMaterialGetV10DataListVideoSource  QianchuanReportMaterialGetV10DataListVideoSource = "LIVE_HIGHLIGHT"
	STAR_QianchuanReportMaterialGetV10DataListVideoSource            QianchuanReportMaterialGetV10DataListVideoSource = "STAR"
	TADA_QianchuanReportMaterialGetV10DataListVideoSource            QianchuanReportMaterialGetV10DataListVideoSource = "TADA"
	VIDEO_CAPTURE_QianchuanReportMaterialGetV10DataListVideoSource   QianchuanReportMaterialGetV10DataListVideoSource = "VIDEO_CAPTURE"
)

// All allowed values of QianchuanReportMaterialGetV10DataListVideoSource enum
var AllowedQianchuanReportMaterialGetV10DataListVideoSourceEnumValues = []QianchuanReportMaterialGetV10DataListVideoSource{
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

// NewQianchuanReportMaterialGetV10DataListVideoSourceFromValue returns a pointer to a valid QianchuanReportMaterialGetV10DataListVideoSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportMaterialGetV10DataListVideoSourceFromValue(v string) (*QianchuanReportMaterialGetV10DataListVideoSource, error) {
	ev := QianchuanReportMaterialGetV10DataListVideoSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportMaterialGetV10DataListVideoSource: valid values are %v", v, AllowedQianchuanReportMaterialGetV10DataListVideoSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportMaterialGetV10DataListVideoSource) IsValid() bool {
	for _, existing := range AllowedQianchuanReportMaterialGetV10DataListVideoSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_material_get_v1.0_data_list_video_source value
func (v QianchuanReportMaterialGetV10DataListVideoSource) Ptr() *QianchuanReportMaterialGetV10DataListVideoSource {
	return &v
}
