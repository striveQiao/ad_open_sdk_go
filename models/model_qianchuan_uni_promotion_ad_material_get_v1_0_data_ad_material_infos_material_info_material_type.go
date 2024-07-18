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

// QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoMaterialType
type QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoMaterialType string

// List of qianchuan_uni_promotion_ad_material_get_v1.0_data_ad_material_infos_material_info_material_type
const (
	IMAGE_QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoMaterialType     QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoMaterialType = "IMAGE"
	LIVE_ROOM_QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoMaterialType QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoMaterialType = "LIVE_ROOM"
	TITLE_QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoMaterialType     QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoMaterialType = "TITLE"
	VIDEO_QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoMaterialType     QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoMaterialType = "VIDEO"
)

// All allowed values of QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoMaterialType enum
var AllowedQianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoMaterialTypeEnumValues = []QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoMaterialType{
	"IMAGE",
	"LIVE_ROOM",
	"TITLE",
	"VIDEO",
}

// NewQianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoMaterialTypeFromValue returns a pointer to a valid QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoMaterialType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoMaterialTypeFromValue(v string) (*QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoMaterialType, error) {
	ev := QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoMaterialType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoMaterialType: valid values are %v", v, AllowedQianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoMaterialTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoMaterialType) IsValid() bool {
	for _, existing := range AllowedQianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoMaterialTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_uni_promotion_ad_material_get_v1.0_data_ad_material_infos_material_info_material_type value
func (v QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoMaterialType) Ptr() *QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoMaterialType {
	return &v
}
