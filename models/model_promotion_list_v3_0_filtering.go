/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionListV30Filtering
type PromotionListV30Filtering struct {
	BlueFlowPackageSetting *PromotionListV30FilteringBlueFlowPackageSetting `json:"blue_flow_package_setting,omitempty"`
	DeliveryMode           *PromotionListV30FilteringDeliveryMode           `json:"delivery_mode,omitempty"`
	HasCarryMaterial       *PromotionListV30FilteringHasCarryMaterial       `json:"has_carry_material,omitempty"`
	//
	Ids []int64 `json:"ids,omitempty"`
	//
	LearningPhase []*PromotionListV30FilteringLearningPhase `json:"learning_phase,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	ProjectId *int64 `json:"project_id,omitempty"`
	//
	PromotionCreateTime *string `json:"promotion_create_time,omitempty"`
	//
	PromotionModifyTime *string `json:"promotion_modify_time,omitempty"`
	//
	RejectReasonType []*PromotionListV30FilteringRejectReasonType `json:"reject_reason_type,omitempty"`
	StarDeliveryType *PromotionListV30FilteringStarDeliveryType   `json:"star_delivery_type,omitempty"`
	Status           *PromotionListV30FilteringStatus             `json:"status,omitempty"`
	StatusFirst      *PromotionListV30FilteringStatusFirst        `json:"status_first,omitempty"`
	StatusSecond     *PromotionListV30FilteringStatusSecond       `json:"status_second,omitempty"`
}
