/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanCarouselGetV10ResponseData
type QianchuanCarouselGetV10ResponseData struct {
	//
	Carousels []*QianchuanCarouselGetV10ResponseDataCarouselsInner `json:"carousels,omitempty"`
	PageInfo  *QianchuanCarouselGetV10ResponseDataPageInfo         `json:"page_info,omitempty"`
}
