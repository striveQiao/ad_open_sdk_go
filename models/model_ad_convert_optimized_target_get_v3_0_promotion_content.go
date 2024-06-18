/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdConvertOptimizedTargetGetV30PromotionContent
type AdConvertOptimizedTargetGetV30PromotionContent string

// List of ad_convert_optimized_target_get_v3.0_promotion_content
const (
	DOWNLOAD_URL_AdConvertOptimizedTargetGetV30PromotionContent AdConvertOptimizedTargetGetV30PromotionContent = "DOWNLOAD_URL"
	EXTERNAL_URL_AdConvertOptimizedTargetGetV30PromotionContent AdConvertOptimizedTargetGetV30PromotionContent = "EXTERNAL_URL"
)

// All allowed values of AdConvertOptimizedTargetGetV30PromotionContent enum
var AllowedAdConvertOptimizedTargetGetV30PromotionContentEnumValues = []AdConvertOptimizedTargetGetV30PromotionContent{
	"DOWNLOAD_URL",
	"EXTERNAL_URL",
}

// NewAdConvertOptimizedTargetGetV30PromotionContentFromValue returns a pointer to a valid AdConvertOptimizedTargetGetV30PromotionContent
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdConvertOptimizedTargetGetV30PromotionContentFromValue(v string) (*AdConvertOptimizedTargetGetV30PromotionContent, error) {
	ev := AdConvertOptimizedTargetGetV30PromotionContent(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdConvertOptimizedTargetGetV30PromotionContent: valid values are %v", v, AllowedAdConvertOptimizedTargetGetV30PromotionContentEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdConvertOptimizedTargetGetV30PromotionContent) IsValid() bool {
	for _, existing := range AllowedAdConvertOptimizedTargetGetV30PromotionContentEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_convert_optimized_target_get_v3.0_promotion_content value
func (v AdConvertOptimizedTargetGetV30PromotionContent) Ptr() *AdConvertOptimizedTargetGetV30PromotionContent {
	return &v
}
