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

// QianchuanAwemeAuthorizedGetV10DataAwemeIdListBindType
type QianchuanAwemeAuthorizedGetV10DataAwemeIdListBindType string

// List of qianchuan_aweme_authorized_get_v1.0_data_aweme_id_list_bind_type
const (
	AWEME_COOPERATOR_QianchuanAwemeAuthorizedGetV10DataAwemeIdListBindType   QianchuanAwemeAuthorizedGetV10DataAwemeIdListBindType = "AWEME_COOPERATOR"
	DOU_SHOP_LIVE_SELF_QianchuanAwemeAuthorizedGetV10DataAwemeIdListBindType QianchuanAwemeAuthorizedGetV10DataAwemeIdListBindType = "DOU_SHOP_LIVE_SELF"
	EXPIRED_QianchuanAwemeAuthorizedGetV10DataAwemeIdListBindType            QianchuanAwemeAuthorizedGetV10DataAwemeIdListBindType = "EXPIRED"
	OFFICIAL_QianchuanAwemeAuthorizedGetV10DataAwemeIdListBindType           QianchuanAwemeAuthorizedGetV10DataAwemeIdListBindType = "OFFICIAL"
	SELF_QianchuanAwemeAuthorizedGetV10DataAwemeIdListBindType               QianchuanAwemeAuthorizedGetV10DataAwemeIdListBindType = "SELF"
	TALENT_USER_OWNER_QianchuanAwemeAuthorizedGetV10DataAwemeIdListBindType  QianchuanAwemeAuthorizedGetV10DataAwemeIdListBindType = "TALENT_USER_OWNER"
	UNAUTHORIZED_QianchuanAwemeAuthorizedGetV10DataAwemeIdListBindType       QianchuanAwemeAuthorizedGetV10DataAwemeIdListBindType = "UNAUTHORIZED"
	UNKNOWN_QianchuanAwemeAuthorizedGetV10DataAwemeIdListBindType            QianchuanAwemeAuthorizedGetV10DataAwemeIdListBindType = "UNKNOWN"
)

// All allowed values of QianchuanAwemeAuthorizedGetV10DataAwemeIdListBindType enum
var AllowedQianchuanAwemeAuthorizedGetV10DataAwemeIdListBindTypeEnumValues = []QianchuanAwemeAuthorizedGetV10DataAwemeIdListBindType{
	"AWEME_COOPERATOR",
	"DOU_SHOP_LIVE_SELF",
	"EXPIRED",
	"OFFICIAL",
	"SELF",
	"TALENT_USER_OWNER",
	"UNAUTHORIZED",
	"UNKNOWN",
}

// NewQianchuanAwemeAuthorizedGetV10DataAwemeIdListBindTypeFromValue returns a pointer to a valid QianchuanAwemeAuthorizedGetV10DataAwemeIdListBindType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeAuthorizedGetV10DataAwemeIdListBindTypeFromValue(v string) (*QianchuanAwemeAuthorizedGetV10DataAwemeIdListBindType, error) {
	ev := QianchuanAwemeAuthorizedGetV10DataAwemeIdListBindType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeAuthorizedGetV10DataAwemeIdListBindType: valid values are %v", v, AllowedQianchuanAwemeAuthorizedGetV10DataAwemeIdListBindTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeAuthorizedGetV10DataAwemeIdListBindType) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeAuthorizedGetV10DataAwemeIdListBindTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_authorized_get_v1.0_data_aweme_id_list_bind_type value
func (v QianchuanAwemeAuthorizedGetV10DataAwemeIdListBindType) Ptr() *QianchuanAwemeAuthorizedGetV10DataAwemeIdListBindType {
	return &v
}
