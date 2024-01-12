/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionUpdateV30IsCommentDisable
type PromotionUpdateV30IsCommentDisable string

// List of promotion_update_v3.0_is_comment_disable
const (
	OFF_PromotionUpdateV30IsCommentDisable PromotionUpdateV30IsCommentDisable = "OFF"
	ON_PromotionUpdateV30IsCommentDisable  PromotionUpdateV30IsCommentDisable = "ON"
)

// All allowed values of PromotionUpdateV30IsCommentDisable enum
var AllowedPromotionUpdateV30IsCommentDisableEnumValues = []PromotionUpdateV30IsCommentDisable{
	"OFF",
	"ON",
}

// NewPromotionUpdateV30IsCommentDisableFromValue returns a pointer to a valid PromotionUpdateV30IsCommentDisable
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionUpdateV30IsCommentDisableFromValue(v string) (*PromotionUpdateV30IsCommentDisable, error) {
	ev := PromotionUpdateV30IsCommentDisable(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionUpdateV30IsCommentDisable: valid values are %v", v, AllowedPromotionUpdateV30IsCommentDisableEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionUpdateV30IsCommentDisable) IsValid() bool {
	for _, existing := range AllowedPromotionUpdateV30IsCommentDisableEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_update_v3.0_is_comment_disable value
func (v PromotionUpdateV30IsCommentDisable) Ptr() *PromotionUpdateV30IsCommentDisable {
	return &v
}
