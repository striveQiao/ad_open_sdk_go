/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanProductAnalyseCompareCreativeV10ResponseDataSimilarProductCreativeInnerVideoMaterialInner struct for QianchuanProductAnalyseCompareCreativeV10ResponseDataSimilarProductCreativeInnerVideoMaterialInner
type QianchuanProductAnalyseCompareCreativeV10ResponseDataSimilarProductCreativeInnerVideoMaterialInner struct {
	//
	AwemeItemId *int64 `json:"aweme_item_id,omitempty"`
	//
	CoverImageHeight *int64 `json:"cover_image_height,omitempty"`
	//
	CoverImageWebUrl *string `json:"cover_image_web_url,omitempty"`
	//
	CoverImageWidth *int64                                                                                  `json:"cover_image_width,omitempty"`
	Source          *QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeVideoMaterialSource `json:"source,omitempty"`
	//
	VideoDuration *int64 `json:"video_duration,omitempty"`
	//
	VideoId *string `json:"video_id,omitempty"`
	//
	VideoUrl []string `json:"video_url,omitempty"`
}
