/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdUpdateV10RequestCreativeListInnerTitleMaterial
type QianchuanAdUpdateV10RequestCreativeListInnerTitleMaterial struct {
	//
	DynamicWords []*QianchuanAdCreateV10RequestCreativeListInnerTitleMaterialDynamicWordsInner `json:"dynamic_words,omitempty"`
	//
	Id *int64 `json:"id,omitempty"`
	//
	Title     string                                                  `json:"title"`
	TitleType *QianchuanAdUpdateV10CreativeListTitleMaterialTitleType `json:"title_type,omitempty"`
}
