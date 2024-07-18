/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
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

// QueryStatementV2ApiService QueryStatementV2Api service
type QueryStatementV2ApiService service

type ApiOpenApi2QueryStatementGetRequest struct {
	ctx           context.Context
	ApiService    *QueryStatementV2ApiService
	agentId       *int64
	projectIdList *[]int64
}

// 代理商ID
func (r *ApiOpenApi2QueryStatementGetRequest) AgentId(agentId int64) *ApiOpenApi2QueryStatementGetRequest {
	r.agentId = &agentId
	return r
}

// 项目id列表，限制1000个
func (r *ApiOpenApi2QueryStatementGetRequest) ProjectIdList(projectIdList []int64) *ApiOpenApi2QueryStatementGetRequest {
	r.projectIdList = &projectIdList
	return r
}

func (r *ApiOpenApi2QueryStatementGetRequest) Execute() (*QueryStatementV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2QueryStatementGetRequest) AccessToken(accessToken string) *ApiOpenApi2QueryStatementGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2QueryStatementGetRequest) WithLog(enable bool) *ApiOpenApi2QueryStatementGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2QueryStatementGet Method for OpenApi2QueryStatementGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2QueryStatementGetRequest
*/
func (a *QueryStatementV2ApiService) Get(ctx context.Context) *ApiOpenApi2QueryStatementGetRequest {
	return &ApiOpenApi2QueryStatementGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryStatementV2Response
func (a *QueryStatementV2ApiService) getExecute(r *ApiOpenApi2QueryStatementGetRequest) (*QueryStatementV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QueryStatementV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/query/statement/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.agentId == nil {
		return localVarReturnValue, nil, ReportError("agentId is required and must be specified")
	}
	if r.projectIdList == nil {
		return localVarReturnValue, nil, ReportError("projectIdList is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "agent_id", r.agentId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "project_id_list", r.projectIdList)
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
