/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
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

type ApiOpenApiV10QianchuanAwemeEstimateProfitGetRequestExample struct {
	AdvertiserId    int64                                          `json:"advertiser_id"`
	MarketingGoal   QianchuanAwemeEstimateProfitV10MarketingGoal   `json:"marketing_goal"`
	DeliverySetting QianchuanAwemeEstimateProfitV10DeliverySetting `json:"delivery_setting"`
	Audience        QianchuanAwemeEstimateProfitV10Audience        `json:"audience,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/qianchuan/aweme/estimate_profit/ Get
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10QianchuanAwemeEstimateProfitGetRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.QianchuanAwemeEstimateProfitV10Api().
		Get(ctx).
		AccessToken(accessToken).
		AdvertiserId(request.AdvertiserId).MarketingGoal(request.MarketingGoal).DeliverySetting(request.DeliverySetting).Audience(request.Audience).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
