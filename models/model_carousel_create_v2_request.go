/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CarouselCreateV2Request struct for CarouselCreateV2Request
type CarouselCreateV2Request struct {
	//
	AdvertiserId int64                        `json:"advertiser_id"`
	CarouselType CarouselCreateV2CarouselType `json:"carousel_type"`
	//
	FileName *string `json:"file_name,omitempty"`
	//
	Images []*CarouselCreateV2RequestImagesInner `json:"images"`
	// 音频
	VideoId *string `json:"video_id,omitempty"`
}
