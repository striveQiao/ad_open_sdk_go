/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdUpdateV10RequestMultiProductCreativeListInner struct for QianchuanAdUpdateV10RequestMultiProductCreativeListInner
type QianchuanAdUpdateV10RequestMultiProductCreativeListInner struct {
	CreativeMaterialMode QianchuanAdUpdateV10MultiProductCreativeListCreativeMaterialMode         `json:"creative_material_mode"`
	CreativeSetting      *QianchuanAdUpdateV10RequestMultiProductCreativeListInnerCreativeSetting `json:"creative_setting,omitempty"`
	//
	ProductId            *int64                                                                        `json:"product_id,omitempty"`
	ProgrammaticCreative *QianchuanAdUpdateV10RequestMultiProductCreativeListInnerProgrammaticCreative `json:"programmatic_creative,omitempty"`
}
