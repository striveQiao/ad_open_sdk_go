/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CarouselListV2ResponseDataCarouselsInnerImagesInner struct for CarouselListV2ResponseDataCarouselsInnerImagesInner
type CarouselListV2ResponseDataCarouselsInnerImagesInner struct {
	//
	Height *int64 `json:"height,omitempty"`
	//
	ImageId      *string                                                          `json:"image_id,omitempty"`
	ImageSubject *CarouselListV2ResponseDataCarouselsInnerImagesInnerImageSubject `json:"image_subject,omitempty"`
	//
	MaterialId *int64 `json:"material_id,omitempty"`
	//
	Ratio *float64 `json:"ratio,omitempty"`
	//
	Url *string `json:"url,omitempty"`
	//
	Width *int64 `json:"width,omitempty"`
}
