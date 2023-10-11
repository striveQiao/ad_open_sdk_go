/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdCreateV10DeliverySettingBudgetMode
type QianchuanAdCreateV10DeliverySettingBudgetMode string

// List of qianchuan_ad_create_v1.0_delivery_setting_budget_mode
const (
	BUDGET_MODE_DAY_QianchuanAdCreateV10DeliverySettingBudgetMode   QianchuanAdCreateV10DeliverySettingBudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_TOTAL_QianchuanAdCreateV10DeliverySettingBudgetMode QianchuanAdCreateV10DeliverySettingBudgetMode = "BUDGET_MODE_TOTAL"
)

// All allowed values of QianchuanAdCreateV10DeliverySettingBudgetMode enum
var AllowedQianchuanAdCreateV10DeliverySettingBudgetModeEnumValues = []QianchuanAdCreateV10DeliverySettingBudgetMode{
	"BUDGET_MODE_DAY",
	"BUDGET_MODE_TOTAL",
}

// NewQianchuanAdCreateV10DeliverySettingBudgetModeFromValue returns a pointer to a valid QianchuanAdCreateV10DeliverySettingBudgetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10DeliverySettingBudgetModeFromValue(v string) (*QianchuanAdCreateV10DeliverySettingBudgetMode, error) {
	ev := QianchuanAdCreateV10DeliverySettingBudgetMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10DeliverySettingBudgetMode: valid values are %v", v, AllowedQianchuanAdCreateV10DeliverySettingBudgetModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10DeliverySettingBudgetMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10DeliverySettingBudgetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_delivery_setting_budget_mode value
func (v QianchuanAdCreateV10DeliverySettingBudgetMode) Ptr() *QianchuanAdCreateV10DeliverySettingBudgetMode {
	return &v
}
