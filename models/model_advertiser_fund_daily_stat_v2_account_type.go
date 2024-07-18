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

// AdvertiserFundDailyStatV2AccountType
type AdvertiserFundDailyStatV2AccountType string

// List of advertiser_fund_daily_stat_v2_account_type
const (
	AD_AdvertiserFundDailyStatV2AccountType    AdvertiserFundDailyStatV2AccountType = "AD"
	LOCAL_AdvertiserFundDailyStatV2AccountType AdvertiserFundDailyStatV2AccountType = "LOCAL"
	STAR_AdvertiserFundDailyStatV2AccountType  AdvertiserFundDailyStatV2AccountType = "STAR"
)

// All allowed values of AdvertiserFundDailyStatV2AccountType enum
var AllowedAdvertiserFundDailyStatV2AccountTypeEnumValues = []AdvertiserFundDailyStatV2AccountType{
	"AD",
	"LOCAL",
	"STAR",
}

// NewAdvertiserFundDailyStatV2AccountTypeFromValue returns a pointer to a valid AdvertiserFundDailyStatV2AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserFundDailyStatV2AccountTypeFromValue(v string) (*AdvertiserFundDailyStatV2AccountType, error) {
	ev := AdvertiserFundDailyStatV2AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserFundDailyStatV2AccountType: valid values are %v", v, AllowedAdvertiserFundDailyStatV2AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserFundDailyStatV2AccountType) IsValid() bool {
	for _, existing := range AllowedAdvertiserFundDailyStatV2AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_fund_daily_stat_v2_account_type value
func (v AdvertiserFundDailyStatV2AccountType) Ptr() *AdvertiserFundDailyStatV2AccountType {
	return &v
}
