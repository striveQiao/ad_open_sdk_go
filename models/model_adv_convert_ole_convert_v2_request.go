/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvConvertOleConvertV2Request struct for AdvConvertOleConvertV2Request
type AdvConvertOleConvertV2Request struct {
	// adv_id
	AdvertiserId int64 `json:"advertiser_id"`
	// 订单所在自然日范围 结束时间，格式 yyyy-MM-dd。假设获取8.16数据，则填写2023-08-16。获取8.16-8.17数据，则填写2023-08-17。
	DateEndTime string `json:"date_end_time"`
	// 订单所在自然日范围 开始时间，格式 yyyy-MM-dd。假设获取8.16数据，则填写2023-08-16。获取8.16-8.17数据，则填写2023-08-16。
	DateStartTime string `json:"date_start_time"`
	// 下单时间范围 结束时间，可精确到分钟。格式 yyyy-MM-dd HH:mm:ss，示例 2023-08-16 13:00:00
	OrderEndTime string `json:"order_end_time"`
	// 下单时间范围 开始时间，可精确到分钟。格式 yyyy-MM-dd HH:mm:ss，示例 2023-08-16 12:00:00
	OrderStartTime string `json:"order_start_time"`
}
