/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdCreateV10RequestCreativeListInner struct for QianchuanAdCreateV10RequestCreativeListInner
type QianchuanAdCreateV10RequestCreativeListInner struct {
	CarouselMaterial      *QianchuanAdCreateV10RequestCreativeListInnerCarouselMaterial      `json:"carousel_material,omitempty"`
	ImageMaterial         *QianchuanAdCreateV10RequestCreativeListInnerImageMaterial         `json:"image_material,omitempty"`
	ImageMode             QianchuanAdCreateV10CreativeListImageMode                          `json:"image_mode"`
	PromotionCardMaterial *QianchuanAdCreateV10RequestCreativeListInnerPromotionCardMaterial `json:"promotion_card_material,omitempty"`
	TitleMaterial         *QianchuanAdCreateV10RequestCreativeListInnerTitleMaterial         `json:"title_material,omitempty"`
	VideoMaterial         *QianchuanAdCreateV10RequestCreativeListInnerVideoMaterial         `json:"video_material,omitempty"`
}
