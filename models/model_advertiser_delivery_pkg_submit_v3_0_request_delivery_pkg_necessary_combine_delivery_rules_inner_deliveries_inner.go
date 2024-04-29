/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryPkgSubmitV30RequestDeliveryPkgNecessaryCombineDeliveryRulesInnerDeliveriesInner struct for AdvertiserDeliveryPkgSubmitV30RequestDeliveryPkgNecessaryCombineDeliveryRulesInnerDeliveriesInner
type AdvertiserDeliveryPkgSubmitV30RequestDeliveryPkgNecessaryCombineDeliveryRulesInnerDeliveriesInner struct {
	// 资质图片附件，数组长度1~50
	Attachments []int64 `json:"attachments"`
	// 资质类型id，来自【推广产品资质规则配置查询接口】
	QualType int64 `json:"qual_type"`
	// 新增时传0；编辑时传入前置查询到的资质id
	QualificationId int64 `json:"qualification_id"`
}
