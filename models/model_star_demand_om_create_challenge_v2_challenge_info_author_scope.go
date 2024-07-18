/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// StarDemandOmCreateChallengeV2ChallengeInfoAuthorScope
type StarDemandOmCreateChallengeV2ChallengeInfoAuthorScope int64

// List of star_demand_om_create_challenge_v2_challenge_info_author_scope
const (
	Enum_1_StarDemandOmCreateChallengeV2ChallengeInfoAuthorScope StarDemandOmCreateChallengeV2ChallengeInfoAuthorScope = 1
	Enum_2_StarDemandOmCreateChallengeV2ChallengeInfoAuthorScope StarDemandOmCreateChallengeV2ChallengeInfoAuthorScope = 2
)

// All allowed values of StarDemandOmCreateChallengeV2ChallengeInfoAuthorScope enum
var AllowedStarDemandOmCreateChallengeV2ChallengeInfoAuthorScopeEnumValues = []StarDemandOmCreateChallengeV2ChallengeInfoAuthorScope{
	1,
	2,
}

// NewStarDemandOmCreateChallengeV2ChallengeInfoAuthorScopeFromValue returns a pointer to a valid StarDemandOmCreateChallengeV2ChallengeInfoAuthorScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarDemandOmCreateChallengeV2ChallengeInfoAuthorScopeFromValue(v int64) (*StarDemandOmCreateChallengeV2ChallengeInfoAuthorScope, error) {
	ev := StarDemandOmCreateChallengeV2ChallengeInfoAuthorScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarDemandOmCreateChallengeV2ChallengeInfoAuthorScope: valid values are %v", v, AllowedStarDemandOmCreateChallengeV2ChallengeInfoAuthorScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarDemandOmCreateChallengeV2ChallengeInfoAuthorScope) IsValid() bool {
	for _, existing := range AllowedStarDemandOmCreateChallengeV2ChallengeInfoAuthorScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_demand_om_create_challenge_v2_challenge_info_author_scope value
func (v StarDemandOmCreateChallengeV2ChallengeInfoAuthorScope) Ptr() *StarDemandOmCreateChallengeV2ChallengeInfoAuthorScope {
	return &v
}
