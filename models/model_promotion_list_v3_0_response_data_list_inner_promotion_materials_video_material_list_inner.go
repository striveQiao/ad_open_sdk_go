/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionListV30ResponseDataListInnerPromotionMaterialsVideoMaterialListInner struct for PromotionListV30ResponseDataListInnerPromotionMaterialsVideoMaterialListInner
type PromotionListV30ResponseDataListInnerPromotionMaterialsVideoMaterialListInner struct {
	// 引导视频id
	GuideVideoId    *string                                                                     `json:"guide_video_id,omitempty"`
	ImageMode       *PromotionListV30DataListPromotionMaterialsVideoMaterialListImageMode       `json:"image_mode,omitempty"`
	IsCarryMaterial *PromotionListV30DataListPromotionMaterialsVideoMaterialListIsCarryMaterial `json:"is_carry_material,omitempty"`
	//
	ItemId *int64 `json:"item_id,omitempty"`
	//
	MaterialId        *int64                                                                        `json:"material_id,omitempty"`
	MaterialOptStatus *PromotionListV30DataListPromotionMaterialsVideoMaterialListMaterialOptStatus `json:"material_opt_status,omitempty"`
	MaterialStatus    *PromotionListV30DataListPromotionMaterialsVideoMaterialListMaterialStatus    `json:"material_status,omitempty"`
	//
	VideoCoverId      *string                                                                       `json:"video_cover_id,omitempty"`
	VideoHpVisibility *PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoHpVisibility `json:"video_hp_visibility,omitempty"`
	//
	VideoId *string `json:"video_id,omitempty"`
	//
	VideoTaskIds      []string                                                                      `json:"video_task_ids,omitempty"`
	VideoTemplateType *PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoTemplateType `json:"video_template_type,omitempty"`
	//
	VisibleEndDate *string `json:"visible_end_date,omitempty"`
}
