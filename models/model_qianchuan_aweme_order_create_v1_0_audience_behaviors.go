/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeOrderCreateV10AudienceBehaviors
type QianchuanAwemeOrderCreateV10AudienceBehaviors string

// List of qianchuan_aweme_order_create_v1.0_audience_behaviors
const (
	ALL_QianchuanAwemeOrderCreateV10AudienceBehaviors           QianchuanAwemeOrderCreateV10AudienceBehaviors = "ALL"
	FOLLOWED_USER_QianchuanAwemeOrderCreateV10AudienceBehaviors QianchuanAwemeOrderCreateV10AudienceBehaviors = "FOLLOWED_USER"
	GOODS_SHARE_QianchuanAwemeOrderCreateV10AudienceBehaviors   QianchuanAwemeOrderCreateV10AudienceBehaviors = "GOODS_SHARE"
	LIKED_USER_QianchuanAwemeOrderCreateV10AudienceBehaviors    QianchuanAwemeOrderCreateV10AudienceBehaviors = "LIKED_USER"
	LIVE_WATCH_QianchuanAwemeOrderCreateV10AudienceBehaviors    QianchuanAwemeOrderCreateV10AudienceBehaviors = "LIVE_WATCH"
)

// All allowed values of QianchuanAwemeOrderCreateV10AudienceBehaviors enum
var AllowedQianchuanAwemeOrderCreateV10AudienceBehaviorsEnumValues = []QianchuanAwemeOrderCreateV10AudienceBehaviors{
	"ALL",
	"FOLLOWED_USER",
	"GOODS_SHARE",
	"LIKED_USER",
	"LIVE_WATCH",
}

// NewQianchuanAwemeOrderCreateV10AudienceBehaviorsFromValue returns a pointer to a valid QianchuanAwemeOrderCreateV10AudienceBehaviors
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderCreateV10AudienceBehaviorsFromValue(v string) (*QianchuanAwemeOrderCreateV10AudienceBehaviors, error) {
	ev := QianchuanAwemeOrderCreateV10AudienceBehaviors(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderCreateV10AudienceBehaviors: valid values are %v", v, AllowedQianchuanAwemeOrderCreateV10AudienceBehaviorsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderCreateV10AudienceBehaviors) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderCreateV10AudienceBehaviorsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_create_v1.0_audience_behaviors value
func (v QianchuanAwemeOrderCreateV10AudienceBehaviors) Ptr() *QianchuanAwemeOrderCreateV10AudienceBehaviors {
	return &v
}
