/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CarouselListV2ResponseDataCarouselsInner struct for CarouselListV2ResponseDataCarouselsInner
type CarouselListV2ResponseDataCarouselsInner struct {
	Audio        *CarouselListV2ResponseDataCarouselsInnerAudio `json:"audio,omitempty"`
	CarouselType *CarouselListV2DataCarouselsCarouselType       `json:"carousel_type,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	FileName *string `json:"file_name,omitempty"`
	//
	Id *int64 `json:"id,omitempty"`
	//
	Images []*CarouselListV2ResponseDataCarouselsInnerImagesInner `json:"images,omitempty"`
	//
	UpdateTime *string `json:"update_time,omitempty"`
}
