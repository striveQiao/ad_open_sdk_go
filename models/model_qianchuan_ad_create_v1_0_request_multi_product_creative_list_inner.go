/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdCreateV10RequestMultiProductCreativeListInner struct for QianchuanAdCreateV10RequestMultiProductCreativeListInner
type QianchuanAdCreateV10RequestMultiProductCreativeListInner struct {
	CreativeMaterialMode QianchuanAdCreateV10MultiProductCreativeListCreativeMaterialMode         `json:"creative_material_mode"`
	CreativeSetting      *QianchuanAdCreateV10RequestMultiProductCreativeListInnerCreativeSetting `json:"creative_setting,omitempty"`
	//
	ProductId            *int64                                                                        `json:"product_id,omitempty"`
	ProgrammaticCreative *QianchuanAdCreateV10RequestMultiProductCreativeListInnerProgrammaticCreative `json:"programmatic_creative,omitempty"`
}
