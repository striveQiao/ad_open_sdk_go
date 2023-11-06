/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileMaterialDetailV2ResponseDataMaterialsInner struct for FileMaterialDetailV2ResponseDataMaterialsInner
type FileMaterialDetailV2ResponseDataMaterialsInner struct {
	//
	IsAdHighQuality *bool `json:"is_ad_high_quality,omitempty"`
	//
	IsAdLowQualityMaterial *bool `json:"is_ad_low_quality_material,omitempty"`
	//
	IsEcpHighQuality *bool `json:"is_ecp_high_quality,omitempty"`
	//
	IsEcpLowQualityMaterial *bool `json:"is_ecp_low_quality_material,omitempty"`
	//
	IsFirstPublishMaterial *bool `json:"is_first_publish_material,omitempty"`
	//
	IsInefficientMaterial *bool `json:"is_inefficient_material,omitempty"`
	//
	IsSimilarExpectedQueueMaterial *bool `json:"is_similar_expected_queue_material,omitempty"`
	//
	IsSimilarMaterial *bool `json:"is_similar_material,omitempty"`
	//
	IsSimilarQueueMaterial *bool `json:"is_similar_queue_material,omitempty"`
	//
	MaterialId string `json:"material_id"`
	//
	MessageAdLowQualityMaterial []string `json:"message_ad_low_quality_material,omitempty"`
	//
	MessageEcpLowQualityMaterial []string `json:"message_ecp_low_quality_material,omitempty"`
}
