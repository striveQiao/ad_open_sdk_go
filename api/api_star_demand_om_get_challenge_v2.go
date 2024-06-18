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

// StarDemandOmGetChallengeV2ApiService StarDemandOmGetChallengeV2Api service
type StarDemandOmGetChallengeV2ApiService service

type ApiOpenApi2StarDemandOmGetChallengeGetRequest struct {
	ctx             context.Context
	ApiService      *StarDemandOmGetChallengeV2ApiService
	starId          *int64
	challengeTaskId *int64
	developerId     *int64
}

// 客户星图ID
func (r *ApiOpenApi2StarDemandOmGetChallengeGetRequest) StarId(starId int64) *ApiOpenApi2StarDemandOmGetChallengeGetRequest {
	r.starId = &starId
	return r
}

// 任务ID
func (r *ApiOpenApi2StarDemandOmGetChallengeGetRequest) ChallengeTaskId(challengeTaskId int64) *ApiOpenApi2StarDemandOmGetChallengeGetRequest {
	r.challengeTaskId = &challengeTaskId
	return r
}

func (r *ApiOpenApi2StarDemandOmGetChallengeGetRequest) DeveloperId(developerId int64) *ApiOpenApi2StarDemandOmGetChallengeGetRequest {
	r.developerId = &developerId
	return r
}

func (r *ApiOpenApi2StarDemandOmGetChallengeGetRequest) Execute() (*StarDemandOmGetChallengeV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarDemandOmGetChallengeGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarDemandOmGetChallengeGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarDemandOmGetChallengeGetRequest) WithLog(enable bool) *ApiOpenApi2StarDemandOmGetChallengeGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarDemandOmGetChallengeGet Method for OpenApi2StarDemandOmGetChallengeGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarDemandOmGetChallengeGetRequest
*/
func (a *StarDemandOmGetChallengeV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarDemandOmGetChallengeGetRequest {
	return &ApiOpenApi2StarDemandOmGetChallengeGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarDemandOmGetChallengeV2Response
func (a *StarDemandOmGetChallengeV2ApiService) getExecute(r *ApiOpenApi2StarDemandOmGetChallengeGetRequest) (*StarDemandOmGetChallengeV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarDemandOmGetChallengeV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/demand/om_get_challenge/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.starId == nil {
		return localVarReturnValue, nil, ReportError("starId is required and must be specified")
	}
	if r.challengeTaskId == nil {
		return localVarReturnValue, nil, ReportError("challengeTaskId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "challenge_task_id", r.challengeTaskId)
	if r.developerId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "developer_id", r.developerId)
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
