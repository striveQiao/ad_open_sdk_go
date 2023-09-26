/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAccountBudgetUpdateV10Request struct for QianchuanAccountBudgetUpdateV10Request
type QianchuanAccountBudgetUpdateV10Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Budget     *float64                                  `json:"budget,omitempty"`
	BudgetMode QianchuanAccountBudgetUpdateV10BudgetMode `json:"budget_mode"`
}
