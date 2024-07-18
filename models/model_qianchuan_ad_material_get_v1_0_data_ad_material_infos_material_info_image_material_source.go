/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource
type QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource string

// List of qianchuan_ad_material_get_v1.0_data_ad_material_infos_material_info_image_material_source
const (
	AGENT_QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource           QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource = "AGENT"
	AI_GENERATE_QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource     QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource = "AI_GENERATE"
	ARTHUR_QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource          QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource = "ARTHUR"
	AUTO_GENERATE_QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource   QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource = "AUTO_GENERATE"
	AWEME_QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource           QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource = "AWEME"
	BP_QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource              QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource = "BP"
	CREATIVE_CENTER_QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource = "CREATIVE_CENTER"
	E_COMMERCE_QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource      QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource = "E_COMMERCE"
	JIANYING_QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource        QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource = "JIANYING"
	JI_CHUANG_QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource       QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource = "JI_CHUANG"
	LIVE_HIGHLIGHT_QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource  QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource = "LIVE_HIGHLIGHT"
	QUNFENG_QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource         QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource = "QUNFENG"
	SQUARE_QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource          QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource = "SQUARE"
	STAR_QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource            QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource = "STAR"
	TADA_QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource            QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource = "TADA"
	VIDEO_CAPTURE_QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource   QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource = "VIDEO_CAPTURE"
)

// All allowed values of QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource enum
var AllowedQianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSourceEnumValues = []QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource{
	"AGENT",
	"AI_GENERATE",
	"ARTHUR",
	"AUTO_GENERATE",
	"AWEME",
	"BP",
	"CREATIVE_CENTER",
	"E_COMMERCE",
	"JIANYING",
	"JI_CHUANG",
	"LIVE_HIGHLIGHT",
	"QUNFENG",
	"SQUARE",
	"STAR",
	"TADA",
	"VIDEO_CAPTURE",
}

// NewQianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSourceFromValue returns a pointer to a valid QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSourceFromValue(v string) (*QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource, error) {
	ev := QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource: valid values are %v", v, AllowedQianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource) IsValid() bool {
	for _, existing := range AllowedQianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_material_get_v1.0_data_ad_material_infos_material_info_image_material_source value
func (v QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource) Ptr() *QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoImageMaterialSource {
	return &v
}
