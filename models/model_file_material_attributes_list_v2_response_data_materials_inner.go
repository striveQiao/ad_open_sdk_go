/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileMaterialAttributesListV2ResponseDataMaterialsInner struct for FileMaterialAttributesListV2ResponseDataMaterialsInner
type FileMaterialAttributesListV2ResponseDataMaterialsInner struct {
	// 当该素材为AD低质素材时，返回低质原因
	AdLowQualitySuggestions []string `json:"ad_low_quality_suggestions,omitempty"`
	//
	AttributesModifyTime *string `json:"attributes_modify_time,omitempty"`
	// 当该素材为千川低质素材时，返回低质原因
	EcpLowQualitySuggestions []string `json:"ecp_low_quality_suggestions,omitempty"`
	// 是否AD优质素材
	IsAdHighQualityMaterial *bool `json:"is_ad_high_quality_material,omitempty"`
	// 是否AD低质素材
	IsAdLowQualityMaterial *bool `json:"is_ad_low_quality_material,omitempty"`
	// 是否存在搬运风险，建议入参account_type = AD 或 QIANCHUAN查询
	IsCarryMaterial *bool `json:"is_carry_material,omitempty"`
	// 是否千川优质素材
	IsEcpHighQualityMaterial *bool `json:"is_ecp_high_quality_material,omitempty"`
	// 是否千川低质素材
	IsEcpLowQualityMaterial *bool `json:"is_ecp_low_quality_material,omitempty"`
	// 是否是首发素材
	IsFirstPublishMaterial *bool `json:"is_first_publish_material,omitempty"`
	// 是否低效素材
	IsInefficientMaterial *bool `json:"is_inefficient_material,omitempty"`
	// 是否同质化素材风险-未投放预计排队素材 - 方舟/纵横不支持
	IsSimilarExpectedQueueMaterial *bool `json:"is_similar_expected_queue_material,omitempty"`
	// 是否同质化挤压严重素材 - 方舟/纵横不支持
	IsSimilarMaterial *bool `json:"is_similar_material,omitempty"`
	// 是否同质化素材风险-排队投放素材 - 方舟/纵横不支持
	IsSimilarQueueMaterial *bool `json:"is_similar_queue_material,omitempty"`
	// 素材id
	MaterialId int64 `json:"material_id"`
}
