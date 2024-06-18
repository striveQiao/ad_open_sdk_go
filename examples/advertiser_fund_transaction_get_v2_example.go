/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	ad_open_sdk_go "github.com/oceanengine/ad_open_sdk_go"
	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

type ApiOpenApi2AdvertiserFundTransactionGetGetRequestExample struct {
	AdvertiserId    int64                                         `json:"advertiser_id"`
	StartDate       string                                        `json:"start_date"`
	EndDate         string                                        `json:"end_date"`
	TransactionType AdvertiserFundTransactionGetV2TransactionType `json:"transaction_type"`
	Page            int32                                         `json:"page,omitempty"`
	PageSize        int32                                         `json:"page_size,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/advertiser/fund/transaction/get/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2AdvertiserFundTransactionGetGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.AdvertiserFundTransactionGetV2Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).StartDate(request.StartDate).EndDate(request.EndDate).TransactionType(request.TransactionType).Page(request.Page).PageSize(request.PageSize).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
