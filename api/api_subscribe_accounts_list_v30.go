/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
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

// SubscribeAccountsListV30ApiService SubscribeAccountsListV30Api service
type SubscribeAccountsListV30ApiService service

type ApiOpenApiV30SubscribeAccountsListGetRequest struct {
	ctx             context.Context
	ApiService      *SubscribeAccountsListV30ApiService
	aPPAccessToken  *string
	appId           *int64
	subscribeTaskId *int64
	events          *[]string
	coreUserId      *int64
	advertiserIds   *[]int64
	statuses        *[]*SubscribeAccountsListV30Statuses
	cursor          *int64
	count           *int64
}

func (r *ApiOpenApiV30SubscribeAccountsListGetRequest) APPAccessToken(aPPAccessToken string) *ApiOpenApiV30SubscribeAccountsListGetRequest {
	r.aPPAccessToken = &aPPAccessToken
	return r
}

func (r *ApiOpenApiV30SubscribeAccountsListGetRequest) AppId(appId int64) *ApiOpenApiV30SubscribeAccountsListGetRequest {
	r.appId = &appId
	return r
}

func (r *ApiOpenApiV30SubscribeAccountsListGetRequest) SubscribeTaskId(subscribeTaskId int64) *ApiOpenApiV30SubscribeAccountsListGetRequest {
	r.subscribeTaskId = &subscribeTaskId
	return r
}

func (r *ApiOpenApiV30SubscribeAccountsListGetRequest) Events(events []string) *ApiOpenApiV30SubscribeAccountsListGetRequest {
	r.events = &events
	return r
}

func (r *ApiOpenApiV30SubscribeAccountsListGetRequest) CoreUserId(coreUserId int64) *ApiOpenApiV30SubscribeAccountsListGetRequest {
	r.coreUserId = &coreUserId
	return r
}

func (r *ApiOpenApiV30SubscribeAccountsListGetRequest) AdvertiserIds(advertiserIds []int64) *ApiOpenApiV30SubscribeAccountsListGetRequest {
	r.advertiserIds = &advertiserIds
	return r
}

func (r *ApiOpenApiV30SubscribeAccountsListGetRequest) Statuses(statuses []*SubscribeAccountsListV30Statuses) *ApiOpenApiV30SubscribeAccountsListGetRequest {
	r.statuses = &statuses
	return r
}

func (r *ApiOpenApiV30SubscribeAccountsListGetRequest) Cursor(cursor int64) *ApiOpenApiV30SubscribeAccountsListGetRequest {
	r.cursor = &cursor
	return r
}

func (r *ApiOpenApiV30SubscribeAccountsListGetRequest) Count(count int64) *ApiOpenApiV30SubscribeAccountsListGetRequest {
	r.count = &count
	return r
}

func (r *ApiOpenApiV30SubscribeAccountsListGetRequest) Execute() (*SubscribeAccountsListV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30SubscribeAccountsListGetRequest) AccessToken(accessToken string) *ApiOpenApiV30SubscribeAccountsListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30SubscribeAccountsListGetRequest) WithLog(enable bool) *ApiOpenApiV30SubscribeAccountsListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30SubscribeAccountsListGet Method for OpenApiV30SubscribeAccountsListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30SubscribeAccountsListGetRequest
*/
func (a *SubscribeAccountsListV30ApiService) Get(ctx context.Context) *ApiOpenApiV30SubscribeAccountsListGetRequest {
	return &ApiOpenApiV30SubscribeAccountsListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SubscribeAccountsListV30Response
func (a *SubscribeAccountsListV30ApiService) getExecute(r *ApiOpenApiV30SubscribeAccountsListGetRequest) (*SubscribeAccountsListV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *SubscribeAccountsListV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/subscribe/accounts/list/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aPPAccessToken == nil {
		return localVarReturnValue, nil, ReportError("aPPAccessToken is required and must be specified")
	}
	if r.appId == nil {
		return localVarReturnValue, nil, ReportError("appId is required and must be specified")
	}
	if r.subscribeTaskId == nil {
		return localVarReturnValue, nil, ReportError("subscribeTaskId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "app_id", r.appId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "subscribe_task_id", r.subscribeTaskId)
	if r.events != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "events", r.events)
	}
	if r.coreUserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "core_user_id", r.coreUserId)
	}
	if r.advertiserIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_ids", r.advertiserIds)
	}
	if r.statuses != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "statuses", r.statuses)
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor)
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count)
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "APP-Access-Token", r.aPPAccessToken)
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
