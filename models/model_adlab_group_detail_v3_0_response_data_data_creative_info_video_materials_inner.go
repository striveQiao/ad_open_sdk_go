/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupDetailV30ResponseDataDataCreativeInfoVideoMaterialsInner struct for AdlabGroupDetailV30ResponseDataDataCreativeInfoVideoMaterialsInner
type AdlabGroupDetailV30ResponseDataDataCreativeInfoVideoMaterialsInner struct {
	//
	ImageInfo []*AdlabGroupDetailV30ResponseDataDataCreativeInfoVideoMaterialsInnerImageInfoInner `json:"image_info,omitempty"`
	ImageMode *AdlabGroupDetailV30DataDataCreativeInfoVideoMaterialsImageMode                     `json:"image_mode,omitempty"`
	VideoInfo *AdlabGroupDetailV30ResponseDataDataCreativeInfoVideoMaterialsInnerVideoInfo        `json:"video_info,omitempty"`
}
