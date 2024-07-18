/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserFundTransactionGetV2ResponseDataListInner struct for AdvertiserFundTransactionGetV2ResponseDataListInner
type AdvertiserFundTransactionGetV2ResponseDataListInner struct {
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	Amount *float64 `json:"amount,omitempty"`
	//
	Cash *float64 `json:"cash,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	Dealbase *float64 `json:"dealbase,omitempty"`
	//
	Frozen *float64 `json:"frozen,omitempty"`
	//
	Grant *float64 `json:"grant,omitempty"`
	//
	Payee *int64 `json:"payee,omitempty"`
	//
	Remitter *int64 `json:"remitter,omitempty"`
	//
	ReturnGoods *float64 `json:"return_goods,omitempty"`
	//
	TransactionSeq  *int64                                                 `json:"transaction_seq,omitempty"`
	TransactionType *AdvertiserFundTransactionGetV2DataListTransactionType `json:"transaction_type,omitempty"`
}
