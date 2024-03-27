/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
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

type ApiOpenApi2ToolsLandingGroupCreatePostRequestExample struct {
	XOrangeCaller                    string                           `json:"x-orange-caller,omitempty"`
	ToolsLandingGroupCreateV2Request ToolsLandingGroupCreateV2Request `json:"ToolsLandingGroupCreateV2Request,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/landing_group/create/ Post
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsLandingGroupCreatePostRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsLandingGroupCreateV2Api().
		Post(ctx).
		AccessToken(accessToken).
		XOrangeCaller(request.XOrangeCaller).ToolsLandingGroupCreateV2Request(request.ToolsLandingGroupCreateV2Request).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
