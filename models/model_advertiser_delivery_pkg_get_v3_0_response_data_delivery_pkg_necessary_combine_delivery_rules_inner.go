/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryPkgGetV30ResponseDataDeliveryPkgNecessaryCombineDeliveryRulesInner struct for AdvertiserDeliveryPkgGetV30ResponseDataDeliveryPkgNecessaryCombineDeliveryRulesInner
type AdvertiserDeliveryPkgGetV30ResponseDataDeliveryPkgNecessaryCombineDeliveryRulesInner struct {
	// 资质的具体信息
	Deliveries []*AdvertiserDeliveryPkgGetV30ResponseDataDeliveryPkgNecessaryCombineDeliveryRulesInnerDeliveriesInner `json:"deliveries"`
	// 原子规则id
	RuleId int64 `json:"rule_id"`
}
