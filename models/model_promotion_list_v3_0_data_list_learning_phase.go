/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionListV30DataListLearningPhase
type PromotionListV30DataListLearningPhase string

// List of promotion_list_v3.0_data_list_learning_phase
const (
	LEARNING_PromotionListV30DataListLearningPhase     PromotionListV30DataListLearningPhase = "LEARNING"
	LEARN_FAILED_PromotionListV30DataListLearningPhase PromotionListV30DataListLearningPhase = "LEARN_FAILED"
	LEARNED_PromotionListV30DataListLearningPhase      PromotionListV30DataListLearningPhase = "LEARNED"
	DEFAULT_PromotionListV30DataListLearningPhase      PromotionListV30DataListLearningPhase = "DEFAULT"
)

// All allowed values of PromotionListV30DataListLearningPhase enum
var AllowedPromotionListV30DataListLearningPhaseEnumValues = []PromotionListV30DataListLearningPhase{
	"LEARNING",
	"LEARN_FAILED",
	"LEARNED",
	"DEFAULT",
}

// NewPromotionListV30DataListLearningPhaseFromValue returns a pointer to a valid PromotionListV30DataListLearningPhase
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListLearningPhaseFromValue(v string) (*PromotionListV30DataListLearningPhase, error) {
	ev := PromotionListV30DataListLearningPhase(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListLearningPhase: valid values are %v", v, AllowedPromotionListV30DataListLearningPhaseEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListLearningPhase) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListLearningPhaseEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_learning_phase value
func (v PromotionListV30DataListLearningPhase) Ptr() *PromotionListV30DataListLearningPhase {
	return &v
}
