/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
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

// ToolsAwemeMultiLevelCategoryGetV2ApiService ToolsAwemeMultiLevelCategoryGetV2Api service
type ToolsAwemeMultiLevelCategoryGetV2ApiService service

type ApiOpenApi2ToolsAwemeMultiLevelCategoryGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsAwemeMultiLevelCategoryGetV2ApiService
	advertiserId *int64
	behaviors    *[]*ToolsAwemeMultiLevelCategoryGetV2Behaviors
}

func (r *ApiOpenApi2ToolsAwemeMultiLevelCategoryGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsAwemeMultiLevelCategoryGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsAwemeMultiLevelCategoryGetGetRequest) Behaviors(behaviors []*ToolsAwemeMultiLevelCategoryGetV2Behaviors) *ApiOpenApi2ToolsAwemeMultiLevelCategoryGetGetRequest {
	r.behaviors = &behaviors
	return r
}

func (r *ApiOpenApi2ToolsAwemeMultiLevelCategoryGetGetRequest) Execute() (*ToolsAwemeMultiLevelCategoryGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsAwemeMultiLevelCategoryGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAwemeMultiLevelCategoryGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAwemeMultiLevelCategoryGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsAwemeMultiLevelCategoryGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAwemeMultiLevelCategoryGetGet Method for OpenApi2ToolsAwemeMultiLevelCategoryGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAwemeMultiLevelCategoryGetGetRequest
*/
func (a *ToolsAwemeMultiLevelCategoryGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsAwemeMultiLevelCategoryGetGetRequest {
	return &ApiOpenApi2ToolsAwemeMultiLevelCategoryGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAwemeMultiLevelCategoryGetV2Response
func (a *ToolsAwemeMultiLevelCategoryGetV2ApiService) getExecute(r *ApiOpenApi2ToolsAwemeMultiLevelCategoryGetGetRequest) (*ToolsAwemeMultiLevelCategoryGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsAwemeMultiLevelCategoryGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/aweme_multi_level_category/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.behaviors != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "behaviors", r.behaviors)
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
