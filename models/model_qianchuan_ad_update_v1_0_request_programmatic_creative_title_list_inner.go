/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdUpdateV10RequestProgrammaticCreativeTitleListInner struct for QianchuanAdUpdateV10RequestProgrammaticCreativeTitleListInner
type QianchuanAdUpdateV10RequestProgrammaticCreativeTitleListInner struct {
	//
	DynamicWords []*QianchuanAdCreateV10RequestProgrammaticCreativeTitleListInnerDynamicWordsInner `json:"dynamic_words,omitempty"`
	//
	Title     string                                                      `json:"title"`
	TitleType *QianchuanAdUpdateV10ProgrammaticCreativeTitleListTitleType `json:"title_type,omitempty"`
}
