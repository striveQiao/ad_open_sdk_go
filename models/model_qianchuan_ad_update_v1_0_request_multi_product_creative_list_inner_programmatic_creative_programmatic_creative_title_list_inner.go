/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdUpdateV10RequestMultiProductCreativeListInnerProgrammaticCreativeProgrammaticCreativeTitleListInner struct for QianchuanAdUpdateV10RequestMultiProductCreativeListInnerProgrammaticCreativeProgrammaticCreativeTitleListInner
type QianchuanAdUpdateV10RequestMultiProductCreativeListInnerProgrammaticCreativeProgrammaticCreativeTitleListInner struct {
	//
	AwemeCarouselId *int64 `json:"aweme_carousel_id,omitempty"`
	//
	DynamicWords []*QianchuanAdUpdateV10RequestMultiProductCreativeListInnerProgrammaticCreativeProgrammaticCreativeTitleListInnerDynamicWordsInner `json:"dynamic_words,omitempty"`
	//
	Title     string                                                                                                  `json:"title"`
	TitleType *QianchuanAdUpdateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType `json:"title_type,omitempty"`
}
