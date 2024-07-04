/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
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

// StarVasGetBoostItemGroupDetailV2ApiService StarVasGetBoostItemGroupDetailV2Api service
type StarVasGetBoostItemGroupDetailV2ApiService service

type ApiOpenApi2StarVasGetBoostItemGroupDetailGetRequest struct {
	ctx        context.Context
	ApiService *StarVasGetBoostItemGroupDetailV2ApiService
	starId     *int64
	taskId     *int64
}

// 客户ID
func (r *ApiOpenApi2StarVasGetBoostItemGroupDetailGetRequest) StarId(starId int64) *ApiOpenApi2StarVasGetBoostItemGroupDetailGetRequest {
	r.starId = &starId
	return r
}

// 助推组ID
func (r *ApiOpenApi2StarVasGetBoostItemGroupDetailGetRequest) TaskId(taskId int64) *ApiOpenApi2StarVasGetBoostItemGroupDetailGetRequest {
	r.taskId = &taskId
	return r
}

func (r *ApiOpenApi2StarVasGetBoostItemGroupDetailGetRequest) Execute() (*StarVasGetBoostItemGroupDetailV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarVasGetBoostItemGroupDetailGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarVasGetBoostItemGroupDetailGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarVasGetBoostItemGroupDetailGetRequest) WithLog(enable bool) *ApiOpenApi2StarVasGetBoostItemGroupDetailGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarVasGetBoostItemGroupDetailGet Method for OpenApi2StarVasGetBoostItemGroupDetailGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarVasGetBoostItemGroupDetailGetRequest
*/
func (a *StarVasGetBoostItemGroupDetailV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarVasGetBoostItemGroupDetailGetRequest {
	return &ApiOpenApi2StarVasGetBoostItemGroupDetailGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarVasGetBoostItemGroupDetailV2Response
func (a *StarVasGetBoostItemGroupDetailV2ApiService) getExecute(r *ApiOpenApi2StarVasGetBoostItemGroupDetailGetRequest) (*StarVasGetBoostItemGroupDetailV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarVasGetBoostItemGroupDetailV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/vas/get_boost_item_group_detail/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.starId == nil {
		return localVarReturnValue, nil, ReportError("starId is required and must be specified")
	}
	if r.taskId == nil {
		return localVarReturnValue, nil, ReportError("taskId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "task_id", r.taskId)
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
