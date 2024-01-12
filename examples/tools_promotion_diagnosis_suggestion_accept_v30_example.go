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

type ApiOpenApiV30ToolsPromotionDiagnosisSuggestionAcceptPostRequestExample struct {
	ToolsPromotionDiagnosisSuggestionAcceptV30Request ToolsPromotionDiagnosisSuggestionAcceptV30Request `json:"ToolsPromotionDiagnosisSuggestionAcceptV30Request,omitempty"`
}

// url: https://api.oceanengine.com/open_api/v3.0/tools/promotion_diagnosis/suggestion/accept/ Post
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiV30ToolsPromotionDiagnosisSuggestionAcceptPostRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsPromotionDiagnosisSuggestionAcceptV30Api().
		Post(ctx).
		AccessToken(accessToken).
		ToolsPromotionDiagnosisSuggestionAcceptV30Request(request.ToolsPromotionDiagnosisSuggestionAcceptV30Request).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
