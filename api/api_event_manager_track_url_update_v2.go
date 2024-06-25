/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
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

// EventManagerTrackUrlUpdateV2ApiService EventManagerTrackUrlUpdateV2Api service
type EventManagerTrackUrlUpdateV2ApiService service

type ApiOpenApi2EventManagerTrackUrlUpdatePostRequest struct {
	ctx                                 context.Context
	ApiService                          *EventManagerTrackUrlUpdateV2ApiService
	eventManagerTrackUrlUpdateV2Request *EventManagerTrackUrlUpdateV2Request
}

func (r *ApiOpenApi2EventManagerTrackUrlUpdatePostRequest) EventManagerTrackUrlUpdateV2Request(eventManagerTrackUrlUpdateV2Request EventManagerTrackUrlUpdateV2Request) *ApiOpenApi2EventManagerTrackUrlUpdatePostRequest {
	r.eventManagerTrackUrlUpdateV2Request = &eventManagerTrackUrlUpdateV2Request
	return r
}

func (r *ApiOpenApi2EventManagerTrackUrlUpdatePostRequest) Execute() (*EventManagerTrackUrlUpdateV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2EventManagerTrackUrlUpdatePostRequest) AccessToken(accessToken string) *ApiOpenApi2EventManagerTrackUrlUpdatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2EventManagerTrackUrlUpdatePostRequest) WithLog(enable bool) *ApiOpenApi2EventManagerTrackUrlUpdatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2EventManagerTrackUrlUpdatePost Method for OpenApi2EventManagerTrackUrlUpdatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2EventManagerTrackUrlUpdatePostRequest
*/
func (a *EventManagerTrackUrlUpdateV2ApiService) Post(ctx context.Context) *ApiOpenApi2EventManagerTrackUrlUpdatePostRequest {
	return &ApiOpenApi2EventManagerTrackUrlUpdatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return EventManagerTrackUrlUpdateV2Response
func (a *EventManagerTrackUrlUpdateV2ApiService) postExecute(r *ApiOpenApi2EventManagerTrackUrlUpdatePostRequest) (*EventManagerTrackUrlUpdateV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *EventManagerTrackUrlUpdateV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/event_manager/track_url/update/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.eventManagerTrackUrlUpdateV2Request
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
