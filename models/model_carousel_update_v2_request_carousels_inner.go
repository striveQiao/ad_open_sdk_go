/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CarouselUpdateV2RequestCarouselsInner struct for CarouselUpdateV2RequestCarouselsInner
type CarouselUpdateV2RequestCarouselsInner struct {
	//
	FileName *string `json:"file_name,omitempty"`
	//
	Id int64 `json:"id"`
	//
	ImageSubjects []*CarouselUpdateV2RequestCarouselsInnerImageSubjectsInner `json:"image_subjects,omitempty"`
}
