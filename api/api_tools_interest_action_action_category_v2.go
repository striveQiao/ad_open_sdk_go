/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
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

// ToolsInterestActionActionCategoryV2ApiService ToolsInterestActionActionCategoryV2Api service
type ToolsInterestActionActionCategoryV2ApiService service

type ApiOpenApi2ToolsInterestActionActionCategoryGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsInterestActionActionCategoryV2ApiService
	actionDays   *ToolsInterestActionActionCategoryV2ActionDays
	advertiserId *int64
}

func (r *ApiOpenApi2ToolsInterestActionActionCategoryGetRequest) ActionDays(actionDays ToolsInterestActionActionCategoryV2ActionDays) *ApiOpenApi2ToolsInterestActionActionCategoryGetRequest {
	r.actionDays = &actionDays
	return r
}

func (r *ApiOpenApi2ToolsInterestActionActionCategoryGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsInterestActionActionCategoryGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsInterestActionActionCategoryGetRequest) Execute() (*ToolsInterestActionActionCategoryV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsInterestActionActionCategoryGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsInterestActionActionCategoryGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsInterestActionActionCategoryGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsInterestActionActionCategoryGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsInterestActionActionCategoryGet Method for OpenApi2ToolsInterestActionActionCategoryGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsInterestActionActionCategoryGetRequest
*/
func (a *ToolsInterestActionActionCategoryV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsInterestActionActionCategoryGetRequest {
	return &ApiOpenApi2ToolsInterestActionActionCategoryGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsInterestActionActionCategoryV2Response
func (a *ToolsInterestActionActionCategoryV2ApiService) getExecute(r *ApiOpenApi2ToolsInterestActionActionCategoryGetRequest) (*ToolsInterestActionActionCategoryV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsInterestActionActionCategoryV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/interest_action/action/category/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.actionDays != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "action_days", r.actionDays)
	}
	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
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
