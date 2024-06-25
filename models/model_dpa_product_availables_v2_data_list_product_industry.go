/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DpaProductAvailablesV2DataListProductIndustry
type DpaProductAvailablesV2DataListProductIndustry string

// List of dpa_product_availables_v2_data_list_product_industry
const (
	WEALTH_DpaProductAvailablesV2DataListProductIndustry           DpaProductAvailablesV2DataListProductIndustry = "WEALTH"
	VIDEO_DpaProductAvailablesV2DataListProductIndustry            DpaProductAvailablesV2DataListProductIndustry = "VIDEO"
	OTHER_DpaProductAvailablesV2DataListProductIndustry            DpaProductAvailablesV2DataListProductIndustry = "OTHER"
	MEDICAL_SERVICE_DpaProductAvailablesV2DataListProductIndustry  DpaProductAvailablesV2DataListProductIndustry = "MEDICAL_SERVICE"
	CREDIT_DpaProductAvailablesV2DataListProductIndustry           DpaProductAvailablesV2DataListProductIndustry = "CREDIT"
	TOUR_HOTEL_DpaProductAvailablesV2DataListProductIndustry       DpaProductAvailablesV2DataListProductIndustry = "TOUR_HOTEL"
	ECOMMERCE_V2_DpaProductAvailablesV2DataListProductIndustry     DpaProductAvailablesV2DataListProductIndustry = "ECOMMERCE_V2"
	TRANSPORT_TICKET_DpaProductAvailablesV2DataListProductIndustry DpaProductAvailablesV2DataListProductIndustry = "TRANSPORT_TICKET"
	RECRUITMENT_DpaProductAvailablesV2DataListProductIndustry      DpaProductAvailablesV2DataListProductIndustry = "RECRUITMENT"
	NOVEL_DpaProductAvailablesV2DataListProductIndustry            DpaProductAvailablesV2DataListProductIndustry = "NOVEL"
	COMMUNICATION_DpaProductAvailablesV2DataListProductIndustry    DpaProductAvailablesV2DataListProductIndustry = "COMMUNICATION"
	MEDICINE_DpaProductAvailablesV2DataListProductIndustry         DpaProductAvailablesV2DataListProductIndustry = "MEDICINE"
	GENERAL_DpaProductAvailablesV2DataListProductIndustry          DpaProductAvailablesV2DataListProductIndustry = "GENERAL"
	FINANCE_DpaProductAvailablesV2DataListProductIndustry          DpaProductAvailablesV2DataListProductIndustry = "FINANCE"
	AUTO_OLD_DpaProductAvailablesV2DataListProductIndustry         DpaProductAvailablesV2DataListProductIndustry = "AUTO_OLD"
	ECOMMERCE_DpaProductAvailablesV2DataListProductIndustry        DpaProductAvailablesV2DataListProductIndustry = "ECOMMERCE"
	NEW_HOUSE_DpaProductAvailablesV2DataListProductIndustry        DpaProductAvailablesV2DataListProductIndustry = "NEW_HOUSE"
	GAME_DpaProductAvailablesV2DataListProductIndustry             DpaProductAvailablesV2DataListProductIndustry = "GAME"
	TOUR_ROUTE_DpaProductAvailablesV2DataListProductIndustry       DpaProductAvailablesV2DataListProductIndustry = "TOUR_ROUTE"
	LIVE_DpaProductAvailablesV2DataListProductIndustry             DpaProductAvailablesV2DataListProductIndustry = "LIVE"
	AUTO_NEW_DpaProductAvailablesV2DataListProductIndustry         DpaProductAvailablesV2DataListProductIndustry = "AUTO_NEW"
	MERCHANTS_DpaProductAvailablesV2DataListProductIndustry        DpaProductAvailablesV2DataListProductIndustry = "MERCHANTS"
	TOUR_TICKET_DpaProductAvailablesV2DataListProductIndustry      DpaProductAvailablesV2DataListProductIndustry = "TOUR_TICKET"
	ESTATE_DpaProductAvailablesV2DataListProductIndustry           DpaProductAvailablesV2DataListProductIndustry = "ESTATE"
	EDUCATION_DpaProductAvailablesV2DataListProductIndustry        DpaProductAvailablesV2DataListProductIndustry = "EDUCATION"
	FURNITURE_DpaProductAvailablesV2DataListProductIndustry        DpaProductAvailablesV2DataListProductIndustry = "FURNITURE"
)

// All allowed values of DpaProductAvailablesV2DataListProductIndustry enum
var AllowedDpaProductAvailablesV2DataListProductIndustryEnumValues = []DpaProductAvailablesV2DataListProductIndustry{
	"WEALTH",
	"VIDEO",
	"OTHER",
	"MEDICAL_SERVICE",
	"CREDIT",
	"TOUR_HOTEL",
	"ECOMMERCE_V2",
	"TRANSPORT_TICKET",
	"RECRUITMENT",
	"NOVEL",
	"COMMUNICATION",
	"MEDICINE",
	"GENERAL",
	"FINANCE",
	"AUTO_OLD",
	"ECOMMERCE",
	"NEW_HOUSE",
	"GAME",
	"TOUR_ROUTE",
	"LIVE",
	"AUTO_NEW",
	"MERCHANTS",
	"TOUR_TICKET",
	"ESTATE",
	"EDUCATION",
	"FURNITURE",
}

// NewDpaProductAvailablesV2DataListProductIndustryFromValue returns a pointer to a valid DpaProductAvailablesV2DataListProductIndustry
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaProductAvailablesV2DataListProductIndustryFromValue(v string) (*DpaProductAvailablesV2DataListProductIndustry, error) {
	ev := DpaProductAvailablesV2DataListProductIndustry(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaProductAvailablesV2DataListProductIndustry: valid values are %v", v, AllowedDpaProductAvailablesV2DataListProductIndustryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaProductAvailablesV2DataListProductIndustry) IsValid() bool {
	for _, existing := range AllowedDpaProductAvailablesV2DataListProductIndustryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_product_availables_v2_data_list_product_industry value
func (v DpaProductAvailablesV2DataListProductIndustry) Ptr() *DpaProductAvailablesV2DataListProductIndustry {
	return &v
}
