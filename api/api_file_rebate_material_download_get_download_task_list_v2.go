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

// FileRebateMaterialDownloadGetDownloadTaskListV2ApiService FileRebateMaterialDownloadGetDownloadTaskListV2Api service
type FileRebateMaterialDownloadGetDownloadTaskListV2ApiService service

type ApiOpenApi2FileRebateMaterialDownloadGetDownloadTaskListGetRequest struct {
	ctx        context.Context
	ApiService *FileRebateMaterialDownloadGetDownloadTaskListV2ApiService
	agentId    *int64
	queryId    *string
}

func (r *ApiOpenApi2FileRebateMaterialDownloadGetDownloadTaskListGetRequest) AgentId(agentId int64) *ApiOpenApi2FileRebateMaterialDownloadGetDownloadTaskListGetRequest {
	r.agentId = &agentId
	return r
}

func (r *ApiOpenApi2FileRebateMaterialDownloadGetDownloadTaskListGetRequest) QueryId(queryId string) *ApiOpenApi2FileRebateMaterialDownloadGetDownloadTaskListGetRequest {
	r.queryId = &queryId
	return r
}

func (r *ApiOpenApi2FileRebateMaterialDownloadGetDownloadTaskListGetRequest) Execute() (*FileRebateMaterialDownloadGetDownloadTaskListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2FileRebateMaterialDownloadGetDownloadTaskListGetRequest) AccessToken(accessToken string) *ApiOpenApi2FileRebateMaterialDownloadGetDownloadTaskListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2FileRebateMaterialDownloadGetDownloadTaskListGetRequest) WithLog(enable bool) *ApiOpenApi2FileRebateMaterialDownloadGetDownloadTaskListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2FileRebateMaterialDownloadGetDownloadTaskListGet Method for OpenApi2FileRebateMaterialDownloadGetDownloadTaskListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2FileRebateMaterialDownloadGetDownloadTaskListGetRequest
*/
func (a *FileRebateMaterialDownloadGetDownloadTaskListV2ApiService) Get(ctx context.Context) *ApiOpenApi2FileRebateMaterialDownloadGetDownloadTaskListGetRequest {
	return &ApiOpenApi2FileRebateMaterialDownloadGetDownloadTaskListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FileRebateMaterialDownloadGetDownloadTaskListV2Response
func (a *FileRebateMaterialDownloadGetDownloadTaskListV2ApiService) getExecute(r *ApiOpenApi2FileRebateMaterialDownloadGetDownloadTaskListGetRequest) (*FileRebateMaterialDownloadGetDownloadTaskListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *FileRebateMaterialDownloadGetDownloadTaskListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/file/rebate/material_download/get_download_task_list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.agentId == nil {
		return localVarReturnValue, nil, ReportError("agentId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "agent_id", r.agentId)
	if r.queryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "query_id", r.queryId)
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
