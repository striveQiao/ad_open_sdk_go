/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// ToolsAbTestInfoGetV2ApiService ToolsAbTestInfoGetV2Api service
type ToolsAbTestInfoGetV2ApiService service

type ApiOpenApi2ToolsAbTestInfoGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsAbTestInfoGetV2ApiService
	advertiserId *int64
	testId       *int64
}

// 广告主ID
func (r *ApiOpenApi2ToolsAbTestInfoGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsAbTestInfoGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 实验ID
func (r *ApiOpenApi2ToolsAbTestInfoGetGetRequest) TestId(testId int64) *ApiOpenApi2ToolsAbTestInfoGetGetRequest {
	r.testId = &testId
	return r
}

func (r *ApiOpenApi2ToolsAbTestInfoGetGetRequest) Execute() (*ToolsAbTestInfoGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsAbTestInfoGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAbTestInfoGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAbTestInfoGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsAbTestInfoGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAbTestInfoGetGet Method for OpenApi2ToolsAbTestInfoGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAbTestInfoGetGetRequest
*/
func (a *ToolsAbTestInfoGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsAbTestInfoGetGetRequest {
	return &ApiOpenApi2ToolsAbTestInfoGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAbTestInfoGetV2Response
func (a *ToolsAbTestInfoGetV2ApiService) getExecute(r *ApiOpenApi2ToolsAbTestInfoGetGetRequest) (*ToolsAbTestInfoGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsAbTestInfoGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/ab_test_info/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.testId == nil {
		return localVarReturnValue, nil, ReportError("testId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "test_id", r.testId)
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
