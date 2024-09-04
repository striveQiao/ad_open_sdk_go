/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueFormDetailV2DataFormFormType
type ClueFormDetailV2DataFormFormType string

// List of clue_form_detail_v2_data_form_form_type
const (
	NATIVE_FORM_ClueFormDetailV2DataFormFormType            ClueFormDetailV2DataFormFormType = "NATIVE_FORM"
	NORMAL_FORM_ClueFormDetailV2DataFormFormType            ClueFormDetailV2DataFormFormType = "NORMAL_FORM"
	ADVANCED_CREATIVE_FORM_ClueFormDetailV2DataFormFormType ClueFormDetailV2DataFormFormType = "ADVANCED_CREATIVE_FORM"
)

// Ptr returns reference to clue_form_detail_v2_data_form_form_type value
func (v ClueFormDetailV2DataFormFormType) Ptr() *ClueFormDetailV2DataFormFormType {
	return &v
}
