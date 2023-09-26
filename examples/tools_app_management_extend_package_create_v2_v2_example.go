/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
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

type ApiOpenApi2ToolsAppManagementExtendPackageCreateV2PostRequestExample struct {
	ToolsAppManagementExtendPackageCreateV2V2Request ToolsAppManagementExtendPackageCreateV2V2Request `json:"ToolsAppManagementExtendPackageCreateV2V2Request,omitempty"`
}

// url: https://api.oceanengine.com/open_api/2/tools/app_management/extend_package/create_v2/ Post
func main() {
	const demoreq = ``
	const accessToken = "ACCESS_TOKEN"
	ctx := context.Background()

	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApi2ToolsAppManagementExtendPackageCreateV2PostRequestExample
	err := json.Unmarshal([]byte(demoreq), &request)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpRes, err := apiClient.ToolsAppManagementExtendPackageCreateV2V2Api().
		Post(ctx).
		AccessToken(accessToken).
		ToolsAppManagementExtendPackageCreateV2V2Request(request.ToolsAppManagementExtendPackageCreateV2V2Request).
		Execute()
	fmt.Println(ToJsonString(resp))
	resBytes, _ := io.ReadAll(httpRes.Body)
	fmt.Println(string(resBytes))
	fmt.Println(err)

}
