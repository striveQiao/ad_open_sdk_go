/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryQualificationListV30ResponseDataListInner struct for AdvertiserDeliveryQualificationListV30ResponseDataListInner
type AdvertiserDeliveryQualificationListV30ResponseDataListInner struct {
	//
	AuditTime string `json:"audit_time"`
	//
	Images []*AdvertiserDeliveryQualificationListV30ResponseDataListInnerImagesInner `json:"images,omitempty"`
	//
	QualificationId   int64                                                           `json:"qualification_id"`
	QualificationType AdvertiserDeliveryQualificationListV30DataListQualificationType `json:"qualification_type"`
	//
	RejectReason string                                               `json:"reject_reason"`
	Status       AdvertiserDeliveryQualificationListV30DataListStatus `json:"status"`
}
