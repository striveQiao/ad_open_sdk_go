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

// DouplusOrderCreateV30TargetingAge
type DouplusOrderCreateV30TargetingAge string

// List of douplus_order_create_v3.0_targeting_age
const (
	AGE_BETWEEN_18_23_DouplusOrderCreateV30TargetingAge DouplusOrderCreateV30TargetingAge = "AGE_BETWEEN_18_23"
	AGE_BETWEEN_24_30_DouplusOrderCreateV30TargetingAge DouplusOrderCreateV30TargetingAge = "AGE_BETWEEN_24_30"
	AGE_BETWEEN_31_40_DouplusOrderCreateV30TargetingAge DouplusOrderCreateV30TargetingAge = "AGE_BETWEEN_31_40"
	AGE_BETWEEN_41_49_DouplusOrderCreateV30TargetingAge DouplusOrderCreateV30TargetingAge = "AGE_BETWEEN_41_49"
)

// All allowed values of DouplusOrderCreateV30TargetingAge enum
var AllowedDouplusOrderCreateV30TargetingAgeEnumValues = []DouplusOrderCreateV30TargetingAge{
	"AGE_BETWEEN_18_23",
	"AGE_BETWEEN_24_30",
	"AGE_BETWEEN_31_40",
	"AGE_BETWEEN_41_49",
}

// NewDouplusOrderCreateV30TargetingAgeFromValue returns a pointer to a valid DouplusOrderCreateV30TargetingAge
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDouplusOrderCreateV30TargetingAgeFromValue(v string) (*DouplusOrderCreateV30TargetingAge, error) {
	ev := DouplusOrderCreateV30TargetingAge(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DouplusOrderCreateV30TargetingAge: valid values are %v", v, AllowedDouplusOrderCreateV30TargetingAgeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DouplusOrderCreateV30TargetingAge) IsValid() bool {
	for _, existing := range AllowedDouplusOrderCreateV30TargetingAgeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to douplus_order_create_v3.0_targeting_age value
func (v DouplusOrderCreateV30TargetingAge) Ptr() *DouplusOrderCreateV30TargetingAge {
	return &v
}
