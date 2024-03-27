/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
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

// StarChallengeAuthorListV2ApiService StarChallengeAuthorListV2Api service
type StarChallengeAuthorListV2ApiService service

type ApiOpenApi2StarChallengeAuthorListGetRequest struct {
	ctx             context.Context
	ApiService      *StarChallengeAuthorListV2ApiService
	starId          *int64
	challengeTaskId *int64
	page            *int32
	limit           *int32
}

// 客户星图ID
func (r *ApiOpenApi2StarChallengeAuthorListGetRequest) StarId(starId int64) *ApiOpenApi2StarChallengeAuthorListGetRequest {
	r.starId = &starId
	return r
}

// 投稿任务ID
func (r *ApiOpenApi2StarChallengeAuthorListGetRequest) ChallengeTaskId(challengeTaskId int64) *ApiOpenApi2StarChallengeAuthorListGetRequest {
	r.challengeTaskId = &challengeTaskId
	return r
}

// 页码
func (r *ApiOpenApi2StarChallengeAuthorListGetRequest) Page(page int32) *ApiOpenApi2StarChallengeAuthorListGetRequest {
	r.page = &page
	return r
}

// 页大小，最大100
func (r *ApiOpenApi2StarChallengeAuthorListGetRequest) Limit(limit int32) *ApiOpenApi2StarChallengeAuthorListGetRequest {
	r.limit = &limit
	return r
}

func (r *ApiOpenApi2StarChallengeAuthorListGetRequest) Execute() (*StarChallengeAuthorListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarChallengeAuthorListGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarChallengeAuthorListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarChallengeAuthorListGetRequest) WithLog(enable bool) *ApiOpenApi2StarChallengeAuthorListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarChallengeAuthorListGet Method for OpenApi2StarChallengeAuthorListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarChallengeAuthorListGetRequest
*/
func (a *StarChallengeAuthorListV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarChallengeAuthorListGetRequest {
	return &ApiOpenApi2StarChallengeAuthorListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarChallengeAuthorListV2Response
func (a *StarChallengeAuthorListV2ApiService) getExecute(r *ApiOpenApi2StarChallengeAuthorListGetRequest) (*StarChallengeAuthorListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarChallengeAuthorListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/challenge/author_list/"

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
	if r.page == nil {
		return localVarReturnValue, nil, ReportError("page is required and must be specified")
	}
	if r.limit == nil {
		return localVarReturnValue, nil, ReportError("limit is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "challenge_task_id", r.challengeTaskId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit)
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
