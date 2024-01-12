/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
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

type ApiOpenApiV30AdvertiserDeliveryQualificationListGetRequestExample struct {
	AdvertiserId      int64                                                   `json:"advertiser_id"`
	Page              int32                                                   `json:"page"`
	PageSize          int32                                                   `json:"page_size"`
	QualificationType AdvertiserDeliveryQualificationListV30QualificationType `json:"qualification_type,omitempty"`
	Status            AdvertiserDeliveryQualificationListV30Status            `json:"status,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/advertiser/delivery_qualification/list/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30AdvertiserDeliveryQualificationListGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.AdvertiserDeliveryQualificationListV30Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).Page(request.Page).PageSize(request.PageSize).QualificationType(request.QualificationType).Status(request.Status).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
