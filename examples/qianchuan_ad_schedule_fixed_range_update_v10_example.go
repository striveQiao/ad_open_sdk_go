/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
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

type ApiOpenApiV10QianchuanAdScheduleFixedRangeUpdatePostRequestExample struct {
	QianchuanAdScheduleFixedRangeUpdateV10Request QianchuanAdScheduleFixedRangeUpdateV10Request `json:"QianchuanAdScheduleFixedRangeUpdateV10Request,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v1.0/qianchuan/ad/schedule_fixed_range/update/ Post
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV10QianchuanAdScheduleFixedRangeUpdatePostRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.QianchuanAdScheduleFixedRangeUpdateV10Api().
		Post(ctx).
		AccessToken(accessToken).
		QianchuanAdScheduleFixedRangeUpdateV10Request(request.QianchuanAdScheduleFixedRangeUpdateV10Request).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
