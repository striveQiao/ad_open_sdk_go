/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponDetailV2ResponseDataForm
type ClueCouponDetailV2ResponseDataForm struct {
	//
	Elements    []*ClueCouponDetailV2ResponseDataFormElementsInner `json:"elements,omitempty"`
	EnableLayer *ClueCouponDetailV2DataFormEnableLayer             `json:"enable_layer,omitempty"`
	ExtendInfo  *ClueCouponDetailV2ResponseDataFormExtendInfo      `json:"extend_info,omitempty"`
	FormType    *ClueCouponDetailV2DataFormFormType                `json:"form_type,omitempty"`
	//
	InstanceId *int64 `json:"instance_id,omitempty"`
	//
	LayerSubmitText *string `json:"layer_submit_text,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	SubmitText   *string                                 `json:"submit_text,omitempty"`
	ValidateType *ClueCouponDetailV2DataFormValidateType `json:"validate_type,omitempty"`
}
