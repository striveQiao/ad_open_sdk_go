/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeOrderCreateV10RequestAudience
type QianchuanAwemeOrderCreateV10RequestAudience struct {
	//
	Age          []*QianchuanAwemeOrderCreateV10AudienceAge       `json:"age,omitempty"`
	AudienceMode QianchuanAwemeOrderCreateV10AudienceAudienceMode `json:"audience_mode"`
	//
	AuthorIds []int64 `json:"author_ids,omitempty"`
	//
	Behaviors []*QianchuanAwemeOrderCreateV10AudienceBehaviors `json:"behaviors,omitempty"`
	//
	City                 []int64                                                   `json:"city,omitempty"`
	District             *QianchuanAwemeOrderCreateV10AudienceDistrict             `json:"district,omitempty"`
	ExcludeLimitedRegion *QianchuanAwemeOrderCreateV10AudienceExcludeLimitedRegion `json:"exclude_limited_region,omitempty"`
	Gender               *QianchuanAwemeOrderCreateV10AudienceGender               `json:"gender,omitempty"`
	//
	InterestCategories []int64 `json:"interest_categories,omitempty"`
}
