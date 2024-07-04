/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeCustomCreativeUpdateV2RequestCreativeListInnerVideoMaterial
type CreativeCustomCreativeUpdateV2RequestCreativeListInnerVideoMaterial struct {
	//
	AwemeItemId *int64 `json:"aweme_item_id,omitempty"`
	//
	DpaVideoTaskIds      []string                                                                      `json:"dpa_video_task_ids,omitempty"`
	DpaVideoTemplateType *CreativeCustomCreativeUpdateV2CreativeListVideoMaterialDpaVideoTemplateType  `json:"dpa_video_template_type,omitempty"`
	ImageInfo            *CreativeCustomCreativeUpdateV2RequestCreativeListInnerVideoMaterialImageInfo `json:"image_info,omitempty"`
	VideoInfo            *CreativeCustomCreativeUpdateV2RequestCreativeListInnerVideoMaterialVideoInfo `json:"video_info,omitempty"`
}
