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

// Oauth2RefreshTokenApiService Oauth2RefreshTokenApi service
type Oauth2RefreshTokenApiService service

type ApiOpenApiOauth2RefreshTokenPostRequest struct {
	ctx                       context.Context
	ApiService                *Oauth2RefreshTokenApiService
	oauth2RefreshTokenRequest *Oauth2RefreshTokenRequest
}

func (r *ApiOpenApiOauth2RefreshTokenPostRequest) Oauth2RefreshTokenRequest(oauth2RefreshTokenRequest Oauth2RefreshTokenRequest) *ApiOpenApiOauth2RefreshTokenPostRequest {
	r.oauth2RefreshTokenRequest = &oauth2RefreshTokenRequest
	return r
}

func (r *ApiOpenApiOauth2RefreshTokenPostRequest) Execute() (*Oauth2RefreshTokenResponse, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiOauth2RefreshTokenPostRequest) AccessToken(accessToken string) *ApiOpenApiOauth2RefreshTokenPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiOauth2RefreshTokenPostRequest) WithLog(enable bool) *ApiOpenApiOauth2RefreshTokenPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiOauth2RefreshTokenPost Method for OpenApiOauth2RefreshTokenPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiOauth2RefreshTokenPostRequest
*/
func (a *Oauth2RefreshTokenApiService) Post(ctx context.Context) *ApiOpenApiOauth2RefreshTokenPostRequest {
	return &ApiOpenApiOauth2RefreshTokenPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Oauth2RefreshTokenResponse
func (a *Oauth2RefreshTokenApiService) postExecute(r *ApiOpenApiOauth2RefreshTokenPostRequest) (*Oauth2RefreshTokenResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *Oauth2RefreshTokenResponse
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/oauth2/refresh_token/"

	localVarHeaderParams := make(map[string]string)
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
	localVarPostBody = r.oauth2RefreshTokenRequest
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
