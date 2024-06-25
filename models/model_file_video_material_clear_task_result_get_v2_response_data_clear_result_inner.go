/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileVideoMaterialClearTaskResultGetV2ResponseDataClearResultInner struct for FileVideoMaterialClearTaskResultGetV2ResponseDataClearResultInner
type FileVideoMaterialClearTaskResultGetV2ResponseDataClearResultInner struct {
	//
	ClearMaterialTypes []*FileVideoMaterialClearTaskResultGetV2DataClearResultClearMaterialTypes `json:"clear_material_types,omitempty"`
	ClearResult        *FileVideoMaterialClearTaskResultGetV2DataClearResultClearResult          `json:"clear_result,omitempty"`
	//
	CreativeFailCount *int64 `json:"creative_fail_count,omitempty"`
	//
	CreativeSuccessCount *int64 `json:"creative_success_count,omitempty"`
	//
	MaterialId *int64 `json:"material_id,omitempty"`
	//
	PromotionFailCount *int64 `json:"promotion_fail_count,omitempty"`
	//
	PromotionSuccessCount *int64 `json:"promotion_success_count,omitempty"`
}
