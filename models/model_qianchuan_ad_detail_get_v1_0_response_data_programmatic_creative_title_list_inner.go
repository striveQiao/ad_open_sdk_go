/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdDetailGetV10ResponseDataProgrammaticCreativeTitleListInner struct for QianchuanAdDetailGetV10ResponseDataProgrammaticCreativeTitleListInner
type QianchuanAdDetailGetV10ResponseDataProgrammaticCreativeTitleListInner struct {
	//
	AwemeCarouselId *int64 `json:"aweme_carousel_id,omitempty"`
	//
	DynamicWords []*QianchuanAdDetailGetV10ResponseDataProgrammaticCreativeTitleListInnerDynamicWordsInner `json:"dynamic_words,omitempty"`
	//
	Title     *string                                                            `json:"title,omitempty"`
	TitleType *QianchuanAdDetailGetV10DataProgrammaticCreativeTitleListTitleType `json:"title_type,omitempty"`
}
