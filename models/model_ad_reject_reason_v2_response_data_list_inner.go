/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdRejectReasonV2ResponseDataListInner struct for AdRejectReasonV2ResponseDataListInner
type AdRejectReasonV2ResponseDataListInner struct {
	AdReject *AdRejectReasonV2ResponseDataListInnerAdReject `json:"ad_reject,omitempty"`
	//
	CreativeReject []*AdRejectReasonV2ResponseDataListInnerCreativeRejectInner `json:"creative_reject,omitempty"`
	//
	MaterialReject []*AdRejectReasonV2ResponseDataListInnerCreativeRejectInnerMaterialRejectInner `json:"material_reject,omitempty"`
}
